package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DotfilesDir string   `json:"dotfiles_dir"`
	Components  []string `json:"components"`
	Backup      bool     `json:"backup"`
	DryRun      bool     `json:"dry_run"`
	Verbose     bool     `json:"verbose"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Default(), nil
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}

func Default() *Config {
	return &Config{
		Components: []string{
			"zsh",
			"zellij",
			"kitty",
			"starship",
			"fastfetch",
			"tmux",
		},
		Backup:  true,
		Verbose: true,
	}
}

func FindConfig() string {
	candidates := []string{
		"pacman-dots.json",
		"pacman-dots.jsonc",
		".pacman-dots.json",
		filepath.Join(os.Getenv("HOME"), ".config", "pacman-dots.json"),
	}

	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}

	return ""
}
