package installer

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hustavojhon/dotfiles/internal/backup"
	"github.com/hustavojhon/dotfiles/internal/distro"
	"github.com/hustavojhon/dotfiles/internal/symlink"
)

type Installer struct {
	HomeDir     string
	DotfilesDir string
	Info        distro.DistroInfo
	Distro      distro.DistroAdapter
	Omarchy     bool
	WSL         bool
	DryRun      bool
	// SkipPackages links configs only (set when sudo is unavailable).
	SkipPackages bool
}

func NewInstaller() (*Installer, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot get home dir: %w", err)
	}

	dotfiles, err := FindDotfilesDir()
	if err != nil {
		return nil, err
	}

	info, err := distro.Detect()
	if err != nil {
		return nil, err
	}

	return &Installer{
		HomeDir:     home,
		DotfilesDir: dotfiles,
		Info:        info,
		Distro:      distro.GetAdapter(info),
		Omarchy:     info.IsOmarchy(),
		WSL:         distro.IsWSL(),
	}, nil
}

// FindDotfilesDir locates the repo: $PACMAN_DOTFILES, then the directory of
// the running binary (bin/installer), then the cwd and common clone paths.
func FindDotfilesDir() (string, error) {
	var candidates []string
	if env := os.Getenv("PACMAN_DOTFILES"); env != "" {
		candidates = append(candidates, env)
	}
	if exe, err := os.Executable(); err == nil {
		if exe, err = filepath.EvalSymlinks(exe); err == nil {
			candidates = append(candidates, filepath.Dir(filepath.Dir(exe)))
		}
	}
	home, _ := os.UserHomeDir()
	candidates = append(candidates,
		".",
		filepath.Join(home, "Github", "hustavojhon", "dotfiles"),
		filepath.Join(home, ".local", "share", "pacman-dotfiles"),
		filepath.Join(home, "dotfiles"),
	)

	for _, c := range candidates {
		abs, err := filepath.Abs(c)
		if err != nil {
			continue
		}
		if isRepo(abs) {
			return abs, nil
		}
	}
	return "", fmt.Errorf("no se encontró el repo de dotfiles (define PACMAN_DOTFILES=/ruta/al/repo)")
}

func isRepo(dir string) bool {
	for _, marker := range []string{"shell/.zshrc", "cmd/installer/main.go"} {
		if info, err := os.Stat(filepath.Join(dir, marker)); err != nil || info.IsDir() {
			return false
		}
	}
	return true
}

// AvailableComponents filters Components for the current platform.
func (inst *Installer) AvailableComponents() []Component {
	var out []Component
	for _, c := range Components {
		if c.Platform == OmarchyOnly && !inst.Omarchy {
			continue
		}
		if c.Platform == NotOmarchy && inst.Omarchy {
			continue
		}
		if c.GUI && inst.WSL {
			continue
		}
		out = append(out, c)
	}
	return out
}

// FindComponents resolves IDs ("all" selects every available component).
func (inst *Installer) FindComponents(ids []string) ([]Component, error) {
	available := inst.AvailableComponents()
	if len(ids) == 1 && ids[0] == "all" {
		return available, nil
	}
	var out []Component
	for _, id := range ids {
		id = strings.TrimSpace(id)
		found := false
		for _, c := range available {
			if c.ID == id {
				out = append(out, c)
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("componente desconocido o no disponible en este sistema: %q", id)
		}
	}
	return out, nil
}

// MissingPackages lists packages of comps that are not installed yet.
func (inst *Installer) MissingPackages(comps []Component) []string {
	seen := map[string]bool{}
	var missing []string
	for _, c := range comps {
		for _, p := range c.Packages {
			if seen[p] {
				continue
			}
			seen[p] = true
			if !inst.isInstalled(p) {
				missing = append(missing, p)
			}
		}
	}
	return missing
}

// binaryOf maps package names whose binary differs from the package name.
var binaryOf = map[string]string{"neovim": "nvim", "ripgrep": "rg"}

// isInstalled accepts a package as present when the package manager knows it
// or its binary is already on PATH (e.g. installed through Homebrew/mise).
func (inst *Installer) isInstalled(pkg string) bool {
	bin := pkg
	if b, ok := binaryOf[pkg]; ok {
		bin = b
	}
	if _, err := exec.LookPath(bin); err == nil {
		return true
	}
	return inst.Distro.IsInstalled(pkg)
}

// SudoCmd returns the command that caches sudo credentials. It must run with
// the real terminal attached (tea.ExecProcess or plain os.Stdin).
func SudoCmd() *exec.Cmd { return exec.Command("sudo", "-v") }

func (inst *Installer) RunBackup(out io.Writer) (string, error) {
	b := backup.NewBackup(inst.HomeDir)
	b.Out = out
	return b.Run()
}

// Result of installing one component.
type Result struct {
	Component Component
	Log       []string
	Warnings  []string
	Err       error
}

func (inst *Installer) InstallComponent(comp Component) Result {
	res := Result{Component: comp}

	if !inst.SkipPackages {
		for _, pkg := range comp.Packages {
			if inst.isInstalled(pkg) {
				continue
			}
			if inst.DryRun {
				res.Log = append(res.Log, "  ~ instalaría paquete "+pkg)
				continue
			}
			if err := inst.Distro.InstallPackage(pkg); err != nil {
				res.Warnings = append(res.Warnings, fmt.Sprintf("paquete %s: %v", pkg, err))
			} else {
				res.Log = append(res.Log, "  ✓ paquete "+pkg)
			}
		}
	}

	var buf strings.Builder
	sm := symlink.NewManager(inst.DotfilesDir, inst.HomeDir)
	sm.DryRun = inst.DryRun
	sm.Out = &buf
	n, err := sm.CreateLinks(comp.Links)
	res.Log = append(res.Log, splitLines(buf.String())...)
	res.Log = append(res.Log, fmt.Sprintf("  %d/%d enlaces listos", n, len(comp.Links)))
	if err != nil {
		res.Err = err
	}

	if comp.PostInstall != nil && !inst.DryRun && res.Err == nil {
		res.Log = append(res.Log, comp.PostInstall(inst)...)
	}
	return res
}

// Status reports the link state of every component, for --status.
func (inst *Installer) Status(w io.Writer) {
	sm := symlink.NewManager(inst.DotfilesDir, inst.HomeDir)
	labels := map[symlink.Status]string{
		symlink.Linked:     "✓ enlazado ",
		symlink.Missing:    "✗ ausente  ",
		symlink.Conflict:   "! conflicto",
		symlink.SrcMissing: "? sin src  ",
	}
	for _, c := range inst.AvailableComponents() {
		fmt.Fprintf(w, "%s\n", c.Name)
		for _, link := range c.Links {
			fmt.Fprintf(w, "  %s  ~/%s\n", labels[sm.Check(link)], link.Dest)
		}
	}
}

func reloadHyprland(inst *Installer) []string {
	if os.Getenv("HYPRLAND_INSTANCE_SIGNATURE") == "" {
		return []string{"  - Hyprland no está corriendo; se aplicará al iniciar sesión"}
	}
	if err := exec.Command("hyprctl", "reload").Run(); err != nil {
		return []string{"  ✗ hyprctl reload: " + err.Error()}
	}
	out, _ := exec.Command("hyprctl", "configerrors").Output()
	errs := strings.TrimSpace(string(out))
	if errs == "" || strings.Contains(errs, "no errors") {
		return []string{"  ✓ Hyprland recargado sin errores"}
	}
	return append([]string{"  ✗ Hyprland reporta errores de config:"}, splitLines(errs)...)
}

func splitLines(s string) []string {
	s = strings.TrimRight(s, "\n")
	if s == "" {
		return nil
	}
	return strings.Split(s, "\n")
}
