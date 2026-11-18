package symlink

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Link maps a path inside the dotfiles repo (Src) to a path under $HOME (Dest).
type Link struct {
	Src  string
	Dest string
	Desc string
}

// Status of a link on disk.
type Status int

const (
	Missing    Status = iota // nothing at Dest
	Linked                   // Dest is a symlink to Src
	Conflict                 // Dest exists but is something else
	SrcMissing               // Src does not exist in the repo
)

type Manager struct {
	DotfilesDir string
	HomeDir     string
	DryRun      bool
	// Out receives human readable progress. Never os.Stdout while a TUI runs.
	Out io.Writer
	// now is overridable for tests.
	now func() time.Time
}

func NewManager(dotfilesDir, homeDir string) *Manager {
	return &Manager{DotfilesDir: dotfilesDir, HomeDir: homeDir, Out: io.Discard, now: time.Now}
}

func (m *Manager) paths(l Link) (string, string) {
	return filepath.Join(m.DotfilesDir, l.Src), filepath.Join(m.HomeDir, l.Dest)
}

func (m *Manager) Check(l Link) Status {
	src, dest := m.paths(l)
	if _, err := os.Stat(src); err != nil {
		return SrcMissing
	}
	info, err := os.Lstat(dest)
	if err != nil {
		return Missing
	}
	if info.Mode()&os.ModeSymlink != 0 {
		if target, _ := os.Readlink(dest); target == src {
			return Linked
		}
	}
	return Conflict
}

// CreateLinks links every entry, returning how many links are in place and
// the first error encountered (it keeps going after errors).
func (m *Manager) CreateLinks(links []Link) (int, error) {
	var count int
	var firstErr error
	for _, l := range links {
		src, dest := m.paths(l)
		switch m.Check(l) {
		case SrcMissing:
			fmt.Fprintf(m.Out, "  - %s: %s no existe en el repo, se omite\n", l.Desc, l.Src)
			continue
		case Linked:
			fmt.Fprintf(m.Out, "  = %s ya enlazado\n", l.Desc)
			count++
			continue
		}

		if m.DryRun {
			fmt.Fprintf(m.Out, "  ~ enlazaría %s -> %s\n", dest, src)
			count++
			continue
		}

		if err := m.link(src, dest); err != nil {
			fmt.Fprintf(m.Out, "  ✗ %s: %v\n", l.Desc, err)
			if firstErr == nil {
				firstErr = err
			}
			continue
		}
		fmt.Fprintf(m.Out, "  ✓ %s\n", l.Desc)
		count++
	}
	return count, firstErr
}

func (m *Manager) link(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return fmt.Errorf("no se pudo crear %s: %w", filepath.Dir(dest), err)
	}

	if _, err := os.Lstat(dest); err == nil {
		// Keep backups out of config dirs: a sibling "theme.bak" would show up
		// as an extra Omarchy theme, an extra Hyprland file, etc.
		backup, err := m.backupPath(dest)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(backup), 0o755); err != nil {
			return err
		}
		if err := os.Rename(dest, backup); err != nil {
			return fmt.Errorf("no se pudo respaldar %s: %w", dest, err)
		}
		fmt.Fprintf(m.Out, "  ~ respaldo: %s\n", backup)
	}

	return os.Symlink(src, dest)
}

// BackupDir is where replaced files go: ~/.dotfiles-backup/links/<timestamp>.
func (m *Manager) BackupDir() string {
	return filepath.Join(m.HomeDir, ".dotfiles-backup", "links", m.now().Format("20060102-150405"))
}

func (m *Manager) backupPath(dest string) (string, error) {
	rel, err := filepath.Rel(m.HomeDir, dest)
	if err != nil || strings.HasPrefix(rel, "..") {
		return "", fmt.Errorf("%s está fuera de $HOME", dest)
	}
	path := filepath.Join(m.BackupDir(), rel)
	// Never overwrite an earlier backup taken in the same second.
	for i := 1; ; i++ {
		if _, err := os.Lstat(path); os.IsNotExist(err) {
			return path, nil
		}
		path = fmt.Sprintf("%s.%d", filepath.Join(m.BackupDir(), rel), i)
	}
}
