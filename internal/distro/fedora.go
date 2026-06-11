package distro

import (
	"os/exec"
)

type FedoraAdapter struct {
	info DistroInfo
}

func (f *FedoraAdapter) Name() string {
	return "Fedora"
}

func (f *FedoraAdapter) PackageManager() string {
	return "dnf"
}

func (f *FedoraAdapter) InstallPackage(pkg string) error {
	cmd := exec.Command("sudo", "dnf", "install", "-y", pkg)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (f *FedoraAdapter) InstallPackages(pkgs []string) error {
	args := append([]string{"dnf", "install", "-y"}, pkgs...)
	cmd := exec.Command("sudo", args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (f *FedoraAdapter) RemovePackage(pkg string) error {
	cmd := exec.Command("sudo", "dnf", "remove", "-y", pkg)
	return cmd.Run()
}

func (f *FedoraAdapter) UpdateSystem() error {
	cmd := exec.Command("sudo", "dnf", "update", "-y")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
