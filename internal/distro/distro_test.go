package distro

import (
	"strings"
	"testing"
)

func TestParseOSReleaseOmarchy(t *testing.T) {
	info := ParseOSRelease(strings.NewReader("NAME=\"Omarchy\"\nPRETTY_NAME=\"Omarchy\"\nID=omarchy\nID_LIKE=arch\nVERSION_ID=\"4.5.0\"\n"))
	if info.ID != "omarchy" || info.IDLike != "arch" || info.Version != "4.5.0" || info.Name != "Omarchy" {
		t.Fatalf("unexpected: %+v", info)
	}
	if !info.IsOmarchy() {
		t.Fatal("ID=omarchy must be detected")
	}
}

func TestGetAdapter(t *testing.T) {
	cases := map[string]struct{ id, like, pm string }{
		"omarchy":    {"omarchy", "arch", "pacman"},
		"arch":       {"arch", "", "pacman"},
		"manjaro":    {"manjaro", "arch", "pacman"},
		"ubuntu":     {"ubuntu", "debian", "apt"},
		"elementary": {"elementary", "ubuntu debian", "apt"},
		"fedora":     {"fedora", "", "dnf"},
		"nobara":     {"nobara", "fedora", "dnf"},
		"garuda":     {"garuda", "arch", "pacman"},
	}
	for name, c := range cases {
		got := GetAdapter(DistroInfo{ID: c.id, IDLike: c.like}).PackageManager()
		if got != c.pm {
			t.Errorf("%s: got %s want %s", name, got, c.pm)
		}
	}
}
