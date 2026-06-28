# PACMAN Modules

## Available Components

| Component | Packages | Config | AUR |
|---|---|---|---|
| ZSH | zsh, fzf, zoxide | configs/zsh/ | — |
| Zellij | zellij | configs/zellij/ | zellij |
| Kitty | kitty | configs/kitty/ | — |
| WezTerm | wezterm | configs/wezterm/ | wezterm |
| Ghostty | — | configs/ghostty/ | ghostty |
| Fastfetch | fastfetch | configs/fastfetch/ | — |
| Starship | starship | configs/starship/ | — |
| Hyprland | hyprland, hyprlock, hypridle, waybar, wofi, dunst | configs/hypr/ | hyprland |
| Tmux | tmux | configs/tmux/ | — |
| Neovim | neovim | configs/nvim/ | — |

## Python Scripts

| Script | Purpose |
|---|---|
| update_fonts.py | Install Nerd Fonts from assets/ |
| wallpaper_manager.py | Set/rotate wallpapers |
| theme_switcher.py | Switch between Tokyo Night, Catppuccin, Nord |
| package_audit.py | Check installed vs desired packages |

## Backup System

Backups are stored at `~/.dotfiles-backup/YYYYMMDD_HHMMSS/`

Backed up paths:
- `~/.zshrc`, `~/.zshenv`, `~/.p10k.zsh`
- `~/.config/kitty`, `~/.config/nvim`
- `~/.tmux.conf`, `~/.gitconfig`
- `~/.local/share/zsh`
