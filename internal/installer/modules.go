package installer

type Component struct {
	ID          string
	Name        string
	Description string
	Packages    []string
	ConfigDir   string
	ConfigFiles []string
	NeedsAUR    bool
	AURPackages []string
	Optional    bool
}

var Components = []Component{
	{
		ID:          "zsh",
		Name:        "ZSH",
		Description: "Z shell with plugins, completions, and theme",
		Packages:    []string{"zsh", "zsh-completions", "fzf", "zoxide"},
		ConfigDir:   "shell",
		ConfigFiles: []string{".zshrc", "aliases.zsh", "exports.zsh", "functions.zsh", "paths.zsh", "p10k.zsh"},
	},
	{
		ID:          "zellij",
		Name:        "Zellij",
		Description: "Terminal multiplexer with layout management",
		Packages:    []string{"zellij"},
		ConfigDir:   "wm/zellij",
		AURPackages: []string{"zellij"},
		NeedsAUR:    true,
	},
	{
		ID:          "kitty",
		Name:        "Kitty",
		Description: "GPU-accelerated terminal emulator",
		Packages:    []string{"kitty"},
		ConfigDir:   "term/kitty",
		ConfigFiles: []string{"kitty.conf", "color.ini"},
	},
	{
		ID:          "wezterm",
		Name:        "WezTerm",
		Description: "GPU-accelerated terminal with Lua config",
		Packages:    []string{"wezterm"},
		ConfigDir:   "term/wezterm",
		AURPackages: []string{"wezterm"},
		NeedsAUR:    true,
	},
	{
		ID:          "ghostty",
		Name:        "Ghostty",
		Description: "Native GPU terminal emulator",
		ConfigDir:   "term/ghostty",
		AURPackages: []string{"ghostty"},
		NeedsAUR:    true,
	},
	{
		ID:          "fastfetch",
		Name:        "Fastfetch",
		Description: "System information tool (neofetch replacement)",
		Packages:    []string{"fastfetch"},
		ConfigDir:   "term/fastfetch",
	},
	{
		ID:          "starship",
		Name:        "Starship",
		Description: "Cross-shell prompt",
		Packages:    []string{"starship"},
		ConfigDir:   "prompt",
	},
	{
		ID:          "hyprland",
		Name:        "Hyprland",
		Description: "Dynamic tiling Wayland compositor",
		Packages:    []string{"hyprland", "hyprlock", "hypridle", "waybar", "wofi", "dunst"},
		ConfigDir:   "wm/hypr",
		AURPackages: []string{"hyprland", "hyprlock", "hypridle"},
		NeedsAUR:    true,
	},
	{
		ID:          "tmux",
		Name:        "Tmux",
		Description: "Terminal multiplexer",
		Packages:    []string{"tmux"},
		ConfigDir:   "wm/tmux",
		ConfigFiles: []string{".tmux.conf"},
	},
	{
		ID:          "nvim",
		Name:        "Neovim",
		Description: "Modern Vim-based editor",
		Packages:    []string{"neovim"},
		ConfigDir:   "editor/nvim",
	},
}
