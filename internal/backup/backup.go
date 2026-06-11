package backup

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Backup struct {
	HomeDir     string
	BackupRoot  string
	Timestamp   string
	BackupPath  string
	DryRun      bool
}

func NewBackup(homeDir string) *Backup {
	timestamp := time.Now().Format("20060102_150405")
	return &Backup{
		HomeDir:    homeDir,
		BackupRoot: filepath.Join(homeDir, ".dotfiles-backup"),
		Timestamp:  timestamp,
		BackupPath: filepath.Join(homeDir, ".dotfiles-backup", timestamp),
	}
}

var BackupPaths = []string{
	".zshrc",
	".zshenv",
	".p10k.zsh",
	".tmux.conf",
	".gitconfig",
	".config/kitty",
	".config/starship.toml",
	".config/fastfetch",
	".config/hypr",
	".config/zellij",
	".config/nvim",
	".local/share/zsh",
}

func (b *Backup) Run() (string, error) {
	if err := os.MkdirAll(b.BackupPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create backup dir: %w", err)
	}

	var count int
	for _, relPath := range BackupPaths {
		src := filepath.Join(b.HomeDir, relPath)
		if _, err := os.Stat(src); os.IsNotExist(err) {
			continue
		}

		dst := filepath.Join(b.BackupPath, relPath)
		if err := copyPath(src, dst); err != nil {
			fmt.Fprintf(os.Stderr, "  ⚠  backup warning: %s: %v\n", relPath, err)
			continue
		}
		count++
	}

	fmt.Printf("  ✓ %d files/directories backed up to %s\n", count, b.BackupPath)
	return b.BackupPath, nil
}

func copyPath(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return copyDir(src, dst)
	}

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	return copyFile(src, dst)
}

func copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if err := copyPath(srcPath, dstPath); err != nil {
			return err
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	return err
}
