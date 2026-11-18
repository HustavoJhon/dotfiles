# Módulos de PACMAN

Definidos en `internal/installer/modules.go`. `make list` muestra los
disponibles en el sistema actual.

| ID | Paquetes | Origen → destino | Plataforma | Por defecto |
|---|---|---|---|---|
| omarchy | — | `omarchy/*` → `~/.config/{hypr,omarchy,kitty,ghostty,fastfetch}` | solo Omarchy | ✓ |
| zsh | zsh fzf zoxide fd bat eza | `shell/*` → `~/.zshrc`, `~/.dotfiles/zsh/` | todas | ✓ |
| zellij | zellij | `wm/zellij` → `~/.config/zellij` | todas | ✓ |
| tmux | tmux | `wm/tmux/.tmux.conf` → `~/.config/tmux/tmux.conf` | todas | ✓ |
| starship | starship | `prompt/starship.toml` → `~/.config/starship.toml` | todas | |
| kitty | kitty | `term/kitty` → `~/.config/kitty` | no Omarchy | ✓ |
| ghostty | ghostty | `term/ghostty/config` → `~/.config/ghostty/config` | no Omarchy | |
| wezterm | wezterm | `term/wezterm/.wezterm.lua` → `~/.wezterm.lua` | todas | |
| fastfetch | fastfetch | `term/fastfetch` → `~/.config/fastfetch` | no Omarchy | ✓ |
| hyprland | hyprland hyprlock hypridle hyprpaper waybar wofi dunst | `wm/hypr` → `~/.config/hypr` | no Omarchy | |
| nvim | neovim | `editor/nvim` → `~/.config/nvim` | todas | |

## Paquetes

- Se instalan uno a uno; si uno no existe en la distro es un **aviso**, no un error.
- Un paquete cuenta como instalado si el gestor lo conoce **o** su binario ya
  está en el `PATH` (Homebrew, mise, …).
- En Arch/Omarchy se intenta el repo oficial y luego el AUR (yay/paru).
- `sudo` se pide una sola vez, con la terminal real, y solo si falta algo.

## Backups

- Cada archivo existente se mueve a `~/.dotfiles-backup/links/<AAAAMMDD-HHMMSS>/` (misma ruta relativa) antes de enlazar.
- Backup completo opcional en `~/.dotfiles-backup/<AAAAMMDD_HHMMSS>/`
  (lista en `internal/backup/backup.go`).

## Añadir un módulo

1. Pon la config en su carpeta (`term/`, `wm/`, …).
2. Añade un `Component` en `modules.go` con sus `Links`.
3. `make test`: comprueba que las rutas existen y que ningún destino se repite.
