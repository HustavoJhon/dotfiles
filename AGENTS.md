# PACMAN Dotfiles — AI Context

> **Single source of truth** para asistentes AI al trabajar en este repo.

## Estructura del Proyecto

```
omarchy/      # Overrides de Omarchy: Hyprland Lua, shell.json, temas, branding
shell/        # Shell configs (ZSH con Powerlevel10k)
term/         # Terminal emuladores (Kitty, WezTerm, Ghostty)
wm/           # Window manager + multiplexers (Hyprland, Tmux, Zellij)
editor/       # Neovim (LazyVim, LSP, completions)
prompt/       # Starship prompt
cmd/ internal/ # Go TUI installer (Bubble Tea + Lip Gloss)
e2e/          # Tests end-to-end (make e2e, Docker)
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
- Usar `make test` para tests Go y `make e2e` para tests en Docker
- Los módulos se definen en `internal/installer/modules.go` (cada uno con sus `Links`)
- En Omarchy: nunca enlazar `~/.config/hypr` entero; solo los `.lua` de usuario.
  Nunca tocar `/usr/share/omarchy/`. `monitors.lua` no se versiona.
- Commits cortos, uno por paso, sin líneas de coautoría
- Las configs van en sus carpetas categorizadas (`shell/`, `term/`, etc.)
- Los scripts van en `scripts/`
- Assets multimedia en `assets/`

## Skills Disponibles

| Skill | Descripción |
|-------|-------------|
| `tui-patterns` | Patrones Bubble Tea para el instalador |

Ver `skills/` para más detalles.
