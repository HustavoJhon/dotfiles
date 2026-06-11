package distro

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DistroAdapter interface {
	Name() string
	InstallPackage(pkg string) error
	InstallPackages(pkgs []string) error
	RemovePackage(pkg string) error
	UpdateSystem() error
	PackageManager() string
}

type DistroInfo struct {
	ID      string
	IDLike  string
	Version string
	Name    string
}

func Detect() (DistroInfo, error) {
	f, err := os.Open("/etc/os-release")
	if err != nil {
		return DistroInfo{}, fmt.Errorf("cannot detect distro: %w", err)
	}
	defer f.Close()

	var info DistroInfo
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			info.ID = strings.Trim(strings.TrimPrefix(line, "ID="), `"'`)
		}
		if strings.HasPrefix(line, "ID_LIKE=") {
			info.IDLike = strings.Trim(strings.TrimPrefix(line, "ID_LIKE="), `"'`)
		}
		if strings.HasPrefix(line, "VERSION_ID=") {
			info.Version = strings.Trim(strings.TrimPrefix(line, "VERSION_ID="), `"'`)
		}
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			info.Name = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), `"'`)
		}
	}

	return info, nil
}

func GetAdapter(info DistroInfo) DistroAdapter {
	switch info.ID {
	case "arch", "archarm":
		return &ArchAdapter{info: info}
	case "cachyos":
		return &ArchAdapter{info: info}
	case "endeavouros":
		return &ArchAdapter{info: info}
	case "debian":
		return &DebianAdapter{info: info}
	case "ubuntu":
		return &DebianAdapter{info: info}
	case "fedora":
		return &FedoraAdapter{info: info}
	case "pop":
		return &DebianAdapter{info: info}
	case "linuxmint":
		return &DebianAdapter{info: info}
	case "kali":
		return &DebianAdapter{info: info}
	}

	if strings.Contains(info.IDLike, "debian") {
		return &DebianAdapter{info: info}
	}
	if strings.Contains(info.IDLike, "fedora") {
		return &FedoraAdapter{info: info}
	}
	if strings.Contains(info.IDLike, "arch") {
		return &ArchAdapter{info: info}
	}

	return &DebianAdapter{info: info}
}
