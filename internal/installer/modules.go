package installer

import (
	"github.com/hustavojhon/dotfiles/internal/symlink"
)

// Platform restricts on which systems a component is offered.
type Platform int

const (
	Any         Platform = iota
	OmarchyOnly          // only on Omarchy
	NotOmarchy           // replaced by the Omarchy component on Omarchy
)

type Component struct {
	ID          string
	Name        string
	Description string
	// Packages are installed one by one; a missing package is a warning,
	// never a fatal error (names differ between distros).
	Packages []string
	Links    []symlink.Link
	Platform Platform
	// Default marks the component as pre-selected in the menu.
	Default bool
	// PostInstall runs after links are in place (e.g. reload Hyprland).
	PostInstall func(inst *Installer) []string
}

func l(src, dest, desc string) symlink.Link { return symlink.Link{Src: src, Dest: dest, Desc: desc} }

var Components = []Component{
	{
		ID:          "omarchy",
		Name:        "Omarchy",
		Description: "Hyprland Lua, barra, temas, branding PACMAN y terminales",
		Platform:    OmarchyOnly,
		Default:     true,
		Links: []symlink.Link{
			l("omarchy/hypr/hyprland.lua", ".config/hypr/hyprland.lua", "Hyprland principal"),
			l("omarchy/hypr/bindings.lua", ".config/hypr/bindings.lua", "Atajos de teclado"),
			l("omarchy/hypr/input.lua", ".config/hypr/input.lua", "Teclado y touchpad"),
			l("omarchy/hypr/looknfeel.lua", ".config/hypr/looknfeel.lua", "Apariencia"),
			l("omarchy/hypr/autostart.lua", ".config/hypr/autostart.lua", "Autostart"),
			l("omarchy/hypr/hypridle.conf", ".config/hypr/hypridle.conf", "Hypridle"),
			l("omarchy/hypr/hyprlock.conf", ".config/hypr/hyprlock.conf", "Hyprlock"),
			l("omarchy/config/shell.json", ".config/omarchy/shell.json", "Barra / shell de Omarchy"),
			l("omarchy/config/shell.toml", ".config/omarchy/shell.toml", "Fuente del shell"),
			l("omarchy/branding/about.txt", ".config/omarchy/branding/about.txt", "Logo PACMAN (fastfetch/about)"),
			l("omarchy/branding/screensaver.txt", ".config/omarchy/branding/screensaver.txt", "Salvapantallas PACMAN"),
			l("omarchy/themes/wallhaven-3z7dd3", ".config/omarchy/themes/wallhaven-3z7dd3", "Tema wallhaven-3z7dd3"),
			l("omarchy/themes/wallhaven-gjr6ke", ".config/omarchy/themes/wallhaven-gjr6ke", "Tema wallhaven-gjr6ke"),
			l("omarchy/themes/wallhaven-gwdvkd", ".config/omarchy/themes/wallhaven-gwdvkd", "Tema wallhaven-gwdvkd"),
			l("omarchy/kitty/kitty.conf", ".config/kitty/kitty.conf", "Kitty (sigue el tema)"),
			l("omarchy/ghostty/config", ".config/ghostty/config", "Ghostty (sigue el tema)"),
			l("omarchy/fastfetch/config.jsonc", ".config/fastfetch/config.jsonc", "Fastfetch"),
		},
		PostInstall: reloadHyprland,
	},
	{
		ID:          "zsh",
		Name:        "ZSH",
		Description: "Shell modular: aliases, funciones, fzf, zoxide, Powerlevel10k",
		Packages:    []string{"zsh", "fzf", "zoxide", "fd", "bat", "eza"},
		Default:     true,
		Links: []symlink.Link{
			l("shell/.zshrc", ".zshrc", "ZSH"),
			l("shell/aliases.zsh", ".dotfiles/zsh/aliases.zsh", "Aliases"),
			l("shell/exports.zsh", ".dotfiles/zsh/exports.zsh", "Exports"),
			l("shell/functions.zsh", ".dotfiles/zsh/functions.zsh", "Funciones"),
			l("shell/paths.zsh", ".dotfiles/zsh/paths.zsh", "PATH"),
			l("shell/p10k.zsh", ".p10k.zsh", "Powerlevel10k"),
			l("assets/banners/banner.txt", ".dotfiles/banner.txt", "Banner de inicio"),
		},
	},
	{
		ID:          "zellij",
		Name:        "Zellij",
		Description: "Multiplexor de terminal con layouts",
		Packages:    []string{"zellij"},
		Default:     true,
		Links:       []symlink.Link{l("wm/zellij", ".config/zellij", "Zellij")},
	},
	{
		ID:          "tmux",
		Name:        "Tmux",
		Description: "Multiplexor de terminal",
		Packages:    []string{"tmux"},
		Default:     true,
		Links:       []symlink.Link{l("wm/tmux/.tmux.conf", ".config/tmux/tmux.conf", "Tmux")},
	},
	{
		ID:          "starship",
		Name:        "Starship",
		Description: "Prompt multi-shell",
		Packages:    []string{"starship"},
		Links:       []symlink.Link{l("prompt/starship.toml", ".config/starship.toml", "Starship")},
	},
	{
		ID:          "kitty",
		Name:        "Kitty",
		Description: "Terminal acelerada por GPU (tema propio)",
		Packages:    []string{"kitty"},
		Platform:    NotOmarchy,
		Default:     true,
		Links:       []symlink.Link{l("term/kitty", ".config/kitty", "Kitty")},
	},
	{
		ID:          "ghostty",
		Name:        "Ghostty",
		Description: "Terminal nativa por GPU",
		Packages:    []string{"ghostty"},
		Platform:    NotOmarchy,
		Links:       []symlink.Link{l("term/ghostty/config", ".config/ghostty/config", "Ghostty")},
	},
	{
		ID:          "wezterm",
		Name:        "WezTerm",
		Description: "Terminal configurable en Lua",
		Packages:    []string{"wezterm"},
		Links:       []symlink.Link{l("term/wezterm/.wezterm.lua", ".wezterm.lua", "WezTerm")},
	},
	{
		ID:          "fastfetch",
		Name:        "Fastfetch",
		Description: "Información del sistema",
		Packages:    []string{"fastfetch"},
		Platform:    NotOmarchy,
		Default:     true,
		Links:       []symlink.Link{l("term/fastfetch", ".config/fastfetch", "Fastfetch")},
	},
	{
		ID:          "hyprland",
		Name:        "Hyprland",
		Description: "Compositor Wayland (config clásica .conf, fuera de Omarchy)",
		Packages:    []string{"hyprland", "hyprlock", "hypridle", "waybar", "wofi", "dunst"},
		Platform:    NotOmarchy,
		Links:       []symlink.Link{l("wm/hypr", ".config/hypr", "Hyprland")},
	},
	{
		ID:          "nvim",
		Name:        "Neovim",
		Description: "Editor (config LazyVim) — reemplaza ~/.config/nvim",
		Packages:    []string{"neovim"},
		Links:       []symlink.Link{l("editor/nvim", ".config/nvim", "Neovim")},
	},
}
