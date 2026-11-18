package installer

import (
	"os"
	"path/filepath"
	"testing"
)

// Every link source must exist in the repo; this catches stale paths after
// moving directories around.
func TestComponentSourcesExist(t *testing.T) {
	repo, _ := filepath.Abs("../..")
	if !isRepo(repo) {
		t.Fatalf("%s is not the repo root", repo)
	}
	ids := map[string]bool{}
	for _, c := range Components {
		if ids[c.ID] {
			t.Errorf("duplicate component id %q", c.ID)
		}
		ids[c.ID] = true
		if len(c.Links) == 0 {
			t.Errorf("%s has no links", c.ID)
		}
		for _, l := range c.Links {
			if _, err := os.Stat(filepath.Join(repo, l.Src)); err != nil {
				t.Errorf("%s: missing source %s", c.ID, l.Src)
			}
		}
	}
}

// Omarchy and generic configs must never fight over the same destination.
func TestPlatformFiltering(t *testing.T) {
	for _, omarchy := range []bool{true, false} {
		inst := &Installer{Omarchy: omarchy}
		dests := map[string]string{}
		hasOmarchy := false
		for _, c := range inst.AvailableComponents() {
			if c.ID == "omarchy" {
				hasOmarchy = true
			}
			if omarchy && c.ID == "hyprland" {
				t.Error("generic hyprland must be hidden on Omarchy")
			}
			for _, l := range c.Links {
				if prev, ok := dests[l.Dest]; ok {
					t.Errorf("omarchy=%v: %s and %s both link ~/%s", omarchy, prev, c.ID, l.Dest)
				}
				dests[l.Dest] = c.ID
			}
		}
		if hasOmarchy != omarchy {
			t.Errorf("omarchy component visible=%v on omarchy=%v", hasOmarchy, omarchy)
		}
	}
}

func TestFindComponents(t *testing.T) {
	inst := &Installer{Omarchy: true}
	if _, err := inst.FindComponents([]string{"hyprland"}); err == nil {
		t.Error("hyprland must not be installable on Omarchy")
	}
	got, err := inst.FindComponents([]string{"omarchy", " zsh"})
	if err != nil || len(got) != 2 {
		t.Fatalf("got %d, err %v", len(got), err)
	}
}
