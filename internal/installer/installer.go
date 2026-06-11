package installer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hustavojhon/dotfiles/internal/backup"
	"github.com/hustavojhon/dotfiles/internal/distro"
	"github.com/hustavojhon/dotfiles/internal/symlink"
)

type Installer struct {
	HomeDir     string
	DotfilesDir string
	Distro      distro.DistroAdapter
	Components  []Component
	Backup      bool
	DryRun      bool
	Verbose     bool
}

func NewInstaller() (*Installer, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot get home dir: %w", err)
	}

	dotfiles := findDotfilesDir()
	info, err := distro.Detect()
	if err != nil {
		return nil, fmt.Errorf("cannot detect distro: %w", err)
	}

	adapter := distro.GetAdapter(info)

	return &Installer{
		HomeDir:     home,
		DotfilesDir: dotfiles,
		Distro:      adapter,
		Backup:      true,
	}, nil
}

func findDotfilesDir() string {
	candidates := []string{
		".",
		filepath.Join(os.Getenv("HOME"), "Github", "hustavojhon", "dotfiles"),
		filepath.Join(os.Getenv("HOME"), "dotfiles"),
	}

	for _, c := range candidates {
		abs, err := filepath.Abs(c)
		if err != nil {
			continue
		}
		if info, err := os.Stat(filepath.Join(abs, "cmd", "installer", "main.go")); err == nil && !info.IsDir() {
			return abs
		}
		if info, err := os.Stat(filepath.Join(abs, "configs")); err == nil && info.IsDir() {
			return abs
		}
	}

	wd, _ := os.Getwd()
	return wd
}

func (inst *Installer) RunBackup() (string, error) {
	b := backup.NewBackup(inst.HomeDir)
	b.DryRun = inst.DryRun
	return b.Run()
}

func (inst *Installer) InstallComponent(comp Component) error {
	fmt.Printf("  → Installing %s...\n", comp.Name)

	if len(comp.Packages) > 0 {
		if err := inst.Distro.InstallPackages(comp.Packages); err != nil {
			return fmt.Errorf("failed to install packages for %s: %w", comp.Name, err)
		}
	}

	if comp.NeedsAUR && len(comp.AURPackages) > 0 {
		if arch, ok := inst.Distro.(*distro.ArchAdapter); ok {
			for _, pkg := range comp.AURPackages {
				if err := arch.InstallAUR(pkg); err != nil {
					fmt.Fprintf(os.Stderr, "  ⚠  AUR install warning (%s): %v\n", pkg, err)
				}
			}
		} else if _, ok := inst.Distro.(*distro.DebianAdapter); ok {
			fmt.Printf("  ⚠  %s not in apt repos, install manually\n", comp.Name)
		}
	}

	sm := symlink.NewManager(inst.DotfilesDir, inst.HomeDir)
	sm.DryRun = inst.DryRun
	sm.Verbose = inst.Verbose

	for _, link := range symlink.Links {
		if filepath.Base(link.Src) == comp.ConfigDir || containsPath(link.Src, comp.ConfigDir) {
			src := filepath.Join(inst.DotfilesDir, link.Src)
			dest := filepath.Join(inst.HomeDir, link.Dest)
			_ = sm.CreateAll()
			_ = src
			_ = dest
		}
	}

	return nil
}

func containsPath(path, dir string) bool {
	absPath, _ := filepath.Abs(path)
	absDir, _ := filepath.Abs(dir)
	return len(absPath) >= len(absDir) && absPath[:len(absDir)] == absDir
}
