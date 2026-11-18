package distro

type DebianAdapter struct {
	info DistroInfo
}

func (d *DebianAdapter) Name() string {
	if d.info.Name != "" {
		return d.info.Name
	}
	return "Debian/Ubuntu"
}

func (d *DebianAdapter) PackageManager() string { return "apt" }

func (d *DebianAdapter) IsInstalled(pkg string) bool {
	return quiet("dpkg", "-s", pkg)
}

func (d *DebianAdapter) InstallPackage(pkg string) error {
	return run("sudo", "-n", "apt-get", "install", "-y", pkg)
}
