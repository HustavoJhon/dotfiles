package distro

import (
	"fmt"
	"os/exec"
)

type ArchAdapter struct {
	info DistroInfo
}

func (a *ArchAdapter) Name() string {
	return "Arch Linux"
}

func (a *ArchAdapter) PackageManager() string {
	return "pacman"
}

func (a *ArchAdapter) InstallPackage(pkg string) error {
	cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "--needed", pkg)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (a *ArchAdapter) InstallPackages(pkgs []string) error {
	args := append([]string{"pacman", "-S", "--noconfirm", "--needed"}, pkgs...)
	cmd := exec.Command("sudo", args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (a *ArchAdapter) RemovePackage(pkg string) error {
	cmd := exec.Command("sudo", "pacman", "-R", "--noconfirm", pkg)
	return cmd.Run()
}

func (a *ArchAdapter) UpdateSystem() error {
	cmd := exec.Command("sudo", "pacman", "-Syu", "--noconfirm")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (a *ArchAdapter) DetectAURHelper() string {
	helpers := []string{"yay", "paru", "trizen", "pikaur"}
	for _, h := range helpers {
		if _, err := exec.LookPath(h); err == nil {
			return h
		}
	}
	return ""
}

func (a *ArchAdapter) InstallAUR(pkg string) error {
	helper := a.DetectAURHelper()
	if helper == "" {
		return fmt.Errorf("no AUR helper found (install yay or paru)")
	}
	cmd := exec.Command(helper, "-S", "--noconfirm", pkg)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (a *ArchAdapter) IsCachyOS() bool {
	return a.info.ID == "cachyos"
}

func (a *ArchAdapter) IsEndeavourOS() bool {
	return a.info.ID == "endeavouros"
}
