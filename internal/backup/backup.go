package backup

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Backup struct {
	HomeDir    string
	BackupPath string
	Out        io.Writer
}

func NewBackup(homeDir string) *Backup {
	timestamp := time.Now().Format("20060102_150405")
	return &Backup{
		HomeDir:    homeDir,
		BackupPath: filepath.Join(homeDir, ".dotfiles-backup", timestamp),
		Out:        io.Discard,
	}
}

// Paths (relative to $HOME) copied before installing. Symlinks are resolved,
// so the backup always holds real file contents.
var Paths = []string{
	".zshrc",
	".zshenv",
	".p10k.zsh",
	".gitconfig",
	".wezterm.lua",
	".config/tmux/tmux.conf",
	".config/kitty",
	".config/ghostty",
	".config/starship.toml",
	".config/fastfetch",
	".config/zellij",
	".config/nvim",
	".config/hypr",
	".config/omarchy/shell.json",
	".config/omarchy/shell.toml",
	".config/omarchy/branding",
}

func (b *Backup) Run() (string, error) {
	if err := os.MkdirAll(b.BackupPath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create backup dir: %w", err)
	}

	var count int
	for _, rel := range Paths {
		src := filepath.Join(b.HomeDir, rel)
		if _, err := os.Stat(src); err != nil {
			continue
		}
		if err := copyPath(src, filepath.Join(b.BackupPath, rel)); err != nil {
			fmt.Fprintf(b.Out, "  ⚠  backup %s: %v\n", rel, err)
			continue
		}
		count++
	}

	fmt.Fprintf(b.Out, "  ✓ Backup: %d elementos en %s\n", count, b.BackupPath)
	return b.BackupPath, nil
}

func copyPath(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return copyFile(src, dst, info.Mode().Perm())
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}
	for _, e := range entries {
		// Skip nested VCS dirs and sockets/devices; they bloat the backup.
		if e.Name() == ".git" {
			continue
		}
		if e.Type()&(os.ModeSocket|os.ModeNamedPipe|os.ModeDevice) != 0 {
			continue
		}
		if err := copyPath(filepath.Join(src, e.Name()), filepath.Join(dst, e.Name())); err != nil {
			return err
		}
	}
	return nil
}

func copyFile(src, dst string, perm os.FileMode) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	d, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}
