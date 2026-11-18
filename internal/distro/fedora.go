package distro

type FedoraAdapter struct {
	info DistroInfo
}

func (f *FedoraAdapter) Name() string {
	if f.info.Name != "" {
		return f.info.Name
	}
	return "Fedora"
}

func (f *FedoraAdapter) PackageManager() string { return "dnf" }

func (f *FedoraAdapter) IsInstalled(pkg string) bool {
	return quiet("rpm", "-q", pkg)
}

func (f *FedoraAdapter) InstallPackage(pkg string) error {
	return run("sudo", "-n", "dnf", "install", "-y", pkg)
}
