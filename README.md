<div align="center">
  <img src="assets/logo/ghost.png" alt="PACMAN.DOTS" width="120" height="120">
  <h1>PACMAN.DOTS</h1>
  <p>Modern dotfiles management platform with Go TUI installer</p>
  <p>
    <a href="#installation">Installation</a> •
    <a href="#modules">Modules</a> •
    <a href="#architecture">Architecture</a> •
    <a href="docs/ROADMAP.md">Roadmap</a>
  </p>
  <p>
    <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go" alt="Go">
    <img src="https://img.shields.io/badge/Linux-FCC624?style=flat&logo=linux&logoColor=black" alt="Linux">
    <img src="https://img.shields.io/badge/License-GPLv3-blue" alt="License">
  </p>
</div>

---

## Installation

**One-liner:**
```bash
curl -fsSL https://raw.githubusercontent.com/hustavojhon/dotfiles/main/install.sh | bash
```

**Manual:**
```bash
git clone https://github.com/hustavojhon/dotfiles.git ~/dotfiles
cd ~/dotfiles
make install
```

**Backup only:**
```bash
make backup
```

## Modules

| Module | Status | Description |
|---|---|---|
| ZSH | ✓ | Modular shell with plugins, fzf, zoxide |
| Zellij | ✓ | Terminal multiplexer with custom layouts |
| Kitty | ✓ | GPU terminal with Tokyo Night theme |
| WezTerm | ✓ | Lua-configured GPU terminal |
| Ghostty | ✓ | Native GPU terminal |
| Fastfetch | ✓ | System info (neofetch replacement) |
| Starship | ✓ | Cross-shell prompt |
| Hyprland | ✓ | Dynamic tiling Wayland compositor |
| Tmux | ✓ | Terminal multiplexer |
| Neovim | 🔒 | UNTOUCHED — preserved as-is |

**Legend:** ✓ Active · 🔒 Preserved · 🚧 In Progress

## Architecture

```
cmd/installer/     Go TUI (Bubble Tea + Lip Gloss)
internal/          Core logic (distro, backup, symlink)
configs/           Dotfiles organized by tool
scripts/           Shell + Python automation
shaders/           Optional GLSL effects
assets/            Logos, banners, wallpapers
```

See [ARCHITECTURE.md](docs/ARCHITECTURE.md) for full details.

## Distro Support

| Family | Distros | Package Manager |
|---|---|---|
| Arch | Arch, CachyOS, EndeavourOS | pacman + yay/paru |
| Debian | Debian, Ubuntu, Pop, Mint, Kali | apt |
| Fedora | Fedora, RHEL | dnf |

Automatic detection via `/etc/os-release`.

## Visual Identity

- **Mascot:** Pacman Ghost
- **Theme:** Purple + Cyan + Black
- **Style:** Hacker minimalism
- **Terminal Font:** Hack Nerd Font

## Quick Start

```bash
# Clone and install
git clone https://github.com/hustavojhon/dotfiles.git
cd dotfiles
make install

# Or use make commands
make build       # Build Go TUI installer
make backup      # Create timestamped backup
make update      # Pull latest + rebuild
make clean       # Remove artifacts
```

## Tools Used

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — Styling
- [Bubbles](https://github.com/charmbracelet/bubbles) — UI components
- [Oh My Zsh](https://ohmyz.sh) — ZSH framework
- [Powerlevel10k](https://github.com/romkatv/powerlevel10k) — ZSH theme

## License

GNU GPL v3 — See [LICENSE](LICENSE)

## Acknowledgments

Inspired by [Gentleman.Dots](https://github.com/gentleman/dots) architectural patterns, but built with independent design and implementation.
