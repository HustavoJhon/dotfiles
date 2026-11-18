package distro

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// DistroAdapter abstracts the package manager of the running distribution.
type DistroAdapter interface {
	Name() string
	PackageManager() string
	// IsInstalled reports whether pkg is already installed (no root needed).
	IsInstalled(pkg string) bool
	// InstallPackage installs a single package. It runs `sudo -n`, so the
	// caller must refresh sudo credentials first (see installer.EnsureSudo).
	InstallPackage(pkg string) error
}

type DistroInfo struct {
	ID      string
	IDLike  string
	Version string
	Name    string
}

// IsOmarchy reports whether the system is Omarchy (Arch + Hyprland by DHH).
func (d DistroInfo) IsOmarchy() bool {
	if d.ID == "omarchy" {
		return true
	}
	// Older Omarchy releases kept ID=arch; detect them by their install path.
	if _, err := os.Stat("/usr/share/omarchy"); err == nil {
		return true
	}
	if _, err := os.Stat(os.ExpandEnv("$HOME/.local/share/omarchy")); err == nil {
		return true
	}
	return false
}

func Detect() (DistroInfo, error) {
	f, err := os.Open("/etc/os-release")
	if err != nil {
		return DistroInfo{}, fmt.Errorf("cannot detect distro: %w", err)
	}
	defer f.Close()
	return ParseOSRelease(f), nil
}

// ParseOSRelease parses the os-release(5) format.
func ParseOSRelease(r io.Reader) DistroInfo {
	var info DistroInfo
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		key, value, ok := strings.Cut(scanner.Text(), "=")
		if !ok {
			continue
		}
		value = strings.Trim(value, `"'`)
		switch key {
		case "ID":
			info.ID = value
		case "ID_LIKE":
			info.IDLike = value
		case "VERSION_ID":
			info.Version = value
		case "PRETTY_NAME":
			info.Name = value
		}
	}
	return info
}

func GetAdapter(info DistroInfo) DistroAdapter {
	switch info.ID {
	case "arch", "archarm", "cachyos", "endeavouros", "manjaro", "omarchy":
		return &ArchAdapter{info: info}
	case "debian", "ubuntu", "pop", "linuxmint", "kali":
		return &DebianAdapter{info: info}
	case "fedora":
		return &FedoraAdapter{info: info}
	}

	for _, like := range strings.Fields(info.IDLike) {
		switch like {
		case "arch":
			return &ArchAdapter{info: info}
		case "debian", "ubuntu":
			return &DebianAdapter{info: info}
		case "fedora", "rhel":
			return &FedoraAdapter{info: info}
		}
	}

	return &DebianAdapter{info: info}
}

// run executes a command and folds the tail of its output into the error.
func run(name string, args ...string) error {
	out, err := exec.Command(name, args...).CombinedOutput()
	if err == nil {
		return nil
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if last := strings.TrimSpace(lines[len(lines)-1]); last != "" {
		return fmt.Errorf("%w: %s", err, last)
	}
	return err
}

func quiet(name string, args ...string) bool {
	return exec.Command(name, args...).Run() == nil
}
