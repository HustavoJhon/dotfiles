# PACMAN Architecture

## Overview

PACMAN is a modern dotfiles management platform built with:

- **Go** — TUI installer with Bubble Tea
- **ZSH** — Modular shell configuration
- **Python** — Auxiliary automation scripts
- **GLSL** — Terminal/ compositor shaders

## Project Structure

```
cmd/installer/     ← Go entrypoint
internal/          ← Go packages
├── distro/        ← Multi-distro adapter pattern
├── installer/     ← Orchestrator + module definitions
├── tui/           ← Bubble Tea screens
├── backup/        ← Timestamped backup system
├── symlink/       ← Symlink management layer
└── assets/        ← Pacman Ghost ASCII art

shell/ term/ wm/   ← Dotfiles organized by tool
editor/ prompt/
omarchy/           ← Omarchy overrides (Hyprland Lua, shell, themes, branding)
scripts/           ← Shell + Python automation
shaders/           ← Optional GLSL effects
assets/            ← Logos, banners, wallpapers, fonts
docs/              ← Documentation
```

## Distro Adapter Pattern

```go
type DistroAdapter interface {
    Name() string
    PackageManager() string
    IsInstalled(pkg string) bool
    InstallPackage(pkg string) error // sudo -n: credentials cached beforehand
}
```

Each distro family implements this interface:

| Adapter | Package Manager | Distros |
|---|---|---|
| ArchAdapter | pacman (+AUR) | Arch, Omarchy, CachyOS, EndeavourOS, Manjaro |
| DebianAdapter | apt | Debian, Ubuntu, Pop, Mint, Kali |
| FedoraAdapter | dnf | Fedora, RHEL |

Detection reads `/etc/os-release` at runtime. `DistroInfo.IsOmarchy()` switches
the component set (see [OMARCHY.md](OMARCHY.md)).

## TUI Flow

```
Welcome → Component Select → Summary/Backup → sudo -v (only if packages missing) → Install → Done
```

## Symlink Management

Files are NEVER copied. The symlink manager creates `ln -sf` relationships:

```
dotfiles/shell/.zshrc         → ~/.zshrc
dotfiles/omarchy/hypr/*.lua   → ~/.config/hypr/*.lua
```

Pre-existing files are moved to `~/.dotfiles-backup/links/<timestamp>/`. Re-running is
idempotent. Nothing is printed to stdout while the TUI runs: output goes to an
`io.Writer` and is shown inside the TUI.

## CLI

```
installer               TUI
installer --list        components for this system
installer --status      link state
installer --install omarchy,zsh [--dry-run] [--no-packages] [--no-backup]
installer --backup-only
```
