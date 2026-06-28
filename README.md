<div align="center">
  <img src="assets/logo/ghost.png" alt="PACMAN" width="100" height="100">
  <h1>PACMAN</h1>
  <p>dotfiles con instalador TUI</p>
  <p>
    <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go" alt="Go">
    <img src="https://img.shields.io/badge/Linux-FCC624?style=flat&logo=linux&logoColor=black" alt="Linux">
    <img src="https://img.shields.io/badge/License-GPLv3-blue" alt="License">
  </p>
</div>

```bash
curl -fsSL https://raw.githubusercontent.com/hustavojhon/dotfiles/main/install.sh | bash
```

## Estructura

```
shell/        ZSH modular con plugins, fzf, zoxide
term/         Terminales (Kitty, WezTerm, Ghostty)
wm/           Hyprland, Tmux, Zellij
editor/       Neovim (preservado intacto)
prompt/       Starship
cmd/          Go TUI installer
internal/     Lógica del instalador
assets/       Logos, banners, wallpapers, fonts
scripts/      Automación shell/Python
docs/         Documentación
shaders/      Efectos GLSL
windows/      Configs para Windows
```

## Módulos

| Módulo | Estado | Descripción |
|---|---|---|
| ZSH | ✓ | Modular shell con plugins, fzf, zoxide |
| Zellij | ✓ | Terminal multiplexer |
| Kitty | ✓ | GPU terminal |
| WezTerm | ✓ | Terminal Lua-configurable |
| Ghostty | ✓ | Terminal nativa GPU |
| Fastfetch | ✓ | System info |
| Starship | ✓ | Prompt cross-shell |
| Hyprland | ✓ | Wayland compositor |
| Tmux | ✓ | Terminal multiplexer |
| Neovim | 🔒 | Preservado intacto |
| Windows | 📁 | Configs para Windows |

## Pruebas con Docker

```bash
docker build -t pacman-test -f Dockerfile.test .
docker run --rm pacman-test
```

## Comandos

```bash
make build     # Compilar instalador
make install   # Ejecutar instalador
make backup    # Backup timestamp
make update    # Pull + rebuild
```

## Licencia

GNU GPL v3
