package distro

import (
	"fmt"
	"os/exec"
)

type ArchAdapter struct {
	info DistroInfo
}

func (a *ArchAdapter) Name() string {
	if a.info.Name != "" {
		return a.info.Name
	}
	return "Arch Linux"
}

func (a *ArchAdapter) PackageManager() string { return "pacman" }

func (a *ArchAdapter) IsInstalled(pkg string) bool {
	return quiet("pacman", "-Q", pkg)
}

// InstallPackage tries the official repos first and falls back to an AUR
// helper when the package is not found there.
func (a *ArchAdapter) InstallPackage(pkg string) error {
	if quiet("pacman", "-Si", pkg) {
		return run("sudo", "-n", "pacman", "-S", "--noconfirm", "--needed", pkg)
	}
	return a.InstallAUR(pkg)
}

func (a *ArchAdapter) DetectAURHelper() string {
	for _, h := range []string{"yay", "paru", "trizen", "pikaur"} {
		if _, err := exec.LookPath(h); err == nil {
			return h
		}
	}
	return ""
}

func (a *ArchAdapter) InstallAUR(pkg string) error {
	helper := a.DetectAURHelper()
	if helper == "" {
		return fmt.Errorf("%s is not in the official repos and no AUR helper (yay/paru) was found", pkg)
	}
	// The helper calls sudo itself; credentials were cached by EnsureSudo.
	return run(helper, "-S", "--noconfirm", "--needed", pkg)
}
