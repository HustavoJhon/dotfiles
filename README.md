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
| Windows | 📁 | Configs para Windows (PowerShell, nvim, VS Code) |

## Comandos

```bash
make install   # Instalar todo
make backup    # Backup timestamp
make update    # Pull + rebuild
make build     # Build TUI
```

## Licencia

GNU GPL v3
