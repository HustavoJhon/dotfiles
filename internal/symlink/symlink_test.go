package symlink

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func setup(t *testing.T) (*Manager, string, string) {
	t.Helper()
	repo, home := t.TempDir(), t.TempDir()
	if err := os.MkdirAll(filepath.Join(repo, "cfg"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repo, "cfg", "a.conf"), []byte("repo"), 0o644); err != nil {
		t.Fatal(err)
	}
	m := NewManager(repo, home)
	m.now = func() time.Time { return time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC) }
	return m, repo, home
}

func TestCreateLinkAndIdempotent(t *testing.T) {
	m, repo, home := setup(t)
	link := Link{Src: "cfg/a.conf", Dest: ".config/app/a.conf", Desc: "a"}

	for i := 0; i < 2; i++ {
		n, err := m.CreateLinks([]Link{link})
		if err != nil || n != 1 {
			t.Fatalf("run %d: n=%d err=%v", i, n, err)
		}
	}
	target, err := os.Readlink(filepath.Join(home, ".config/app/a.conf"))
	if err != nil || target != filepath.Join(repo, "cfg/a.conf") {
		t.Fatalf("bad link: %q %v", target, err)
	}
	if m.Check(link) != Linked {
		t.Fatal("expected Linked")
	}
	entries, _ := os.ReadDir(filepath.Join(home, ".config/app"))
	if len(entries) != 1 {
		t.Fatalf("second run must not create backups, got %d entries", len(entries))
	}
}

func TestBacksUpExistingFile(t *testing.T) {
	m, _, home := setup(t)
	dest := filepath.Join(home, "a.conf")
	if err := os.WriteFile(dest, []byte("mine"), 0o644); err != nil {
		t.Fatal(err)
	}
	link := Link{Src: "cfg/a.conf", Dest: "a.conf"}
	if m.Check(link) != Conflict {
		t.Fatal("expected Conflict")
	}
	if _, err := m.CreateLinks([]Link{link}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Lstat(dest + ".bak.20260102-030405"); !os.IsNotExist(err) {
		t.Fatal("backups must not be left next to the original")
	}
	data, err := os.ReadFile(filepath.Join(home, ".dotfiles-backup/links/20260102-030405/a.conf"))
	if err != nil || string(data) != "mine" {
		t.Fatalf("backup missing: %q %v", data, err)
	}
}

func TestDryRunAndMissingSource(t *testing.T) {
	m, _, home := setup(t)
	m.DryRun = true
	n, err := m.CreateLinks([]Link{{Src: "cfg/a.conf", Dest: "a.conf"}, {Src: "nope", Dest: "b"}})
	if err != nil || n != 1 {
		t.Fatalf("n=%d err=%v", n, err)
	}
	if _, err := os.Lstat(filepath.Join(home, "a.conf")); !os.IsNotExist(err) {
		t.Fatal("dry-run must not touch the filesystem")
	}
}

func TestRelativeLinkCountsAsLinked(t *testing.T) {
	m, repo, home := setup(t)
	rel, err := filepath.Rel(home, filepath.Join(repo, "cfg/a.conf"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(rel, filepath.Join(home, "a.conf")); err != nil {
		t.Fatal(err)
	}
	if got := m.Check(Link{Src: "cfg/a.conf", Dest: "a.conf"}); got != Linked {
		t.Fatalf("relative symlink: got %v, want Linked", got)
	}
}
