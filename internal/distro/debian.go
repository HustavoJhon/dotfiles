package distro

import (
	"fmt"
	"os/exec"
)

type DebianAdapter struct {
	info DistroInfo
}

func (d *DebianAdapter) Name() string {
	return "Debian/Ubuntu"
}

func (d *DebianAdapter) PackageManager() string {
	return "apt"
}

func (d *DebianAdapter) InstallPackage(pkg string) error {
	cmd := exec.Command("sudo", "apt-get", "install", "-y", pkg)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (d *DebianAdapter) InstallPackages(pkgs []string) error {
	updateCmd := exec.Command("sudo", "apt-get", "update", "-qq")
	updateCmd.Stdout = nil
	updateCmd.Stderr = nil
	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("apt-get update failed: %w", err)
	}
	args := append([]string{"apt-get", "install", "-y"}, pkgs...)
	cmd := exec.Command("sudo", args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (d *DebianAdapter) RemovePackage(pkg string) error {
	cmd := exec.Command("sudo", "apt-get", "remove", "-y", pkg)
	return cmd.Run()
}

func (d *DebianAdapter) UpdateSystem() error {
	cmd := exec.Command("sudo", "apt-get", "update")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (d *DebianAdapter) IsUbuntu() bool {
	return d.info.ID == "ubuntu"
}
