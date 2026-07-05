package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
		if info, err := os.Stat(filepath.Join(abs, "shell", ".zshrc")); err == nil && !info.IsDir() {
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

	var compLinks []symlink.Link
	prefix := comp.ConfigDir + "/"
	for _, link := range symlink.Links {
		if link.Src == comp.ConfigDir || strings.HasPrefix(link.Src, prefix) {
			compLinks = append(compLinks, link)
		}
	}

	if len(compLinks) > 0 {
		if err := sm.CreateLinks(compLinks); err != nil {
			return fmt.Errorf("failed to create symlinks for %s: %w", comp.Name, err)
		}
	} else if inst.Verbose {
		fmt.Printf("  - no symlinks found for %s\n", comp.Name)
	}

	return nil
}
