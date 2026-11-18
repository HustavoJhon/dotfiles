package distro

type DebianAdapter struct {
	info    DistroInfo
	updated bool
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

// InstallPackage refreshes the apt index once per run (fresh containers and
// installs often have an empty or stale index).
func (d *DebianAdapter) InstallPackage(pkg string) error {
	if !d.updated {
		if err := run("sudo", "-n", "apt-get", "update", "-qq"); err != nil {
			return err
		}
		d.updated = true
	}
	return run("sudo", "-n", "apt-get", "install", "-y", pkg)
}
