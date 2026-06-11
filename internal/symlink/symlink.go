package symlink

import (
	"fmt"
	"os"
	"path/filepath"
)

type Manager struct {
	DotfilesDir string
	HomeDir     string
	DryRun      bool
	Verbose     bool
}

func NewManager(dotfilesDir, homeDir string) *Manager {
	return &Manager{
		DotfilesDir: dotfilesDir,
		HomeDir:     homeDir,
	}
}

type Link struct {
	Src  string
	Dest string
	Desc string
}

var Links = []Link{
	{Src: "configs/zsh/.zshrc", Dest: ".zshrc", Desc: "ZSH configuration"},
	{Src: "configs/zsh/aliases.zsh", Dest: ".dotfiles/zsh/aliases.zsh", Desc: "ZSH aliases"},
	{Src: "configs/zsh/exports.zsh", Dest: ".dotfiles/zsh/exports.zsh", Desc: "ZSH exports"},
	{Src: "configs/zsh/functions.zsh", Dest: ".dotfiles/zsh/functions.zsh", Desc: "ZSH functions"},
	{Src: "configs/zsh/paths.zsh", Dest: ".dotfiles/zsh/paths.zsh", Desc: "ZSH paths"},
	{Src: "configs/zsh/p10k.zsh", Dest: ".p10k.zsh", Desc: "Powerlevel10k theme"},
	{Src: "configs/kitty", Dest: ".config/kitty", Desc: "Kitty terminal"},
	{Src: "configs/tmux/.tmux.conf", Dest: ".tmux.conf", Desc: "Tmux configuration"},
	{Src: "configs/zellij", Dest: ".config/zellij", Desc: "Zellij multiplexer"},
	{Src: "configs/starship/starship.toml", Dest: ".config/starship.toml", Desc: "Starship prompt"},
	{Src: "configs/fastfetch", Dest: ".config/fastfetch", Desc: "Fastfetch"},
	{Src: "configs/wezterm/.wezterm.lua", Dest: ".wezterm.lua", Desc: "WezTerm config"},
	{Src: "configs/ghostty/config", Dest: ".config/ghostty/config", Desc: "Ghostty config"},
	{Src: "configs/hypr", Dest: ".config/hypr", Desc: "Hyprland config"},
	{Src: "assets/banners/banner.txt", Dest: ".dotfiles/banner.txt", Desc: "Startup banner"},
}

func (m *Manager) CreateAll() error {
	var count int
	for _, link := range Links {
		src := filepath.Join(m.DotfilesDir, link.Src)
		dest := filepath.Join(m.HomeDir, link.Dest)

		if _, err := os.Stat(src); os.IsNotExist(err) {
			if m.Verbose {
				fmt.Printf("  - %s: source not found, skipping\n", link.Desc)
			}
			continue
		}

		if err := m.createLink(src, dest); err != nil {
			fmt.Fprintf(os.Stderr, "  ✗ %s: %v\n", link.Desc, err)
			continue
		}

		if m.Verbose {
			fmt.Printf("  ✓ %s\n", link.Desc)
		}
		count++
	}

	fmt.Printf("\n  ✓ %d symlinks created\n", count)
	return nil
}

func (m *Manager) createLink(src, dest string) error {
	if m.DryRun {
		fmt.Printf("  ~ would link: %s -> %s\n", dest, src)
		return nil
	}

	destDir := filepath.Dir(dest)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", destDir, err)
	}

	if info, err := os.Lstat(dest); err == nil {
		if info.Mode()&os.ModeSymlink != 0 {
			existing, _ := os.Readlink(dest)
			if existing == src {
				return nil
			}
		}
		backup := dest + ".bak"
		if err := os.Rename(dest, backup); err != nil {
			return fmt.Errorf("failed to backup %s: %w", dest, err)
		}
		fmt.Printf("  ~ backed up existing: %s -> %s\n", dest, backup)
	}

	if err := os.Symlink(src, dest); err != nil {
		return fmt.Errorf("failed to create symlink %s: %w", dest, err)
	}

	return nil
}
