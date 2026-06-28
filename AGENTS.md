# PACMAN Dotfiles — AI Context

> **Single source of truth** para asistentes AI al trabajar en este repo.

## Estructura del Proyecto

```
shell/        # Shell configs (ZSH con Powerlevel10k)
term/         # Terminal emuladores (Kitty, WezTerm, Ghostty)
wm/           # Window manager + multiplexers (Hyprland, Tmux, Zellij)
editor/       # Neovim (LazyVim, LSP, completions)
prompt/       # Starship prompt
installer/    # Go TUI installer (Bubble Tea + Lip Gloss)
scripts/      # Shell + Python automation
assets/       # Logos, banners, wallpapers, fonts
docs/         # Documentación
shaders/      # GLSL effects
windows/      # Configs para Windows
```

## Stack

- **Go** — TUI installer con Bubble Tea + Lip Gloss
- **ZSH** — Modular shell con plugins (autocomplete, syntax highlighting, autosuggestions)
- **Neovim** — Config preservada intacta
- **Kitty/WezTerm/Ghostty** — Terminal emuladores
- **Hyprland** — Wayland compositor
- **Starship** — Cross-shell prompt
- **Tmux/Zellij** — Terminal multiplexers

## Convenciones

- Usar `make build` para compilar el instalador
- Usar `make run` para probar localmente
- Usar `make test` para tests Go
- Las configs van en sus carpetas categorizadas (`shell/`, `term/`, etc.)
- Los scripts van en `scripts/`
- Assets multimedia en `assets/`

## Skills Disponibles

| Skill | Descripción |
|-------|-------------|
| `tui-patterns` | Patrones Bubble Tea para el instalador |
| `installer-steps` | Pasos de instalación, OS detection |

Ver `skills/` para más detalles.
