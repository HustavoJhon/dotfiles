# PACMAN.DOTS Architecture

## Overview

PACMAN.DOTS is a modern dotfiles management platform built with:

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
├── config/        ← JSON config loader
└── assets/        ← Pacman Ghost ASCII art

configs/           ← All dotfiles organized by tool
scripts/           ← Shell + Python automation
shaders/           ← Optional GLSL effects
assets/            ← Logos, banners, wallpapers, fonts
docs/              ← Documentation
```

## Distro Adapter Pattern

```go
type DistroAdapter interface {
    InstallPackage(pkg string) error
    InstallPackages(pkgs []string) error
    RemovePackage(pkg string) error
    UpdateSystem() error
    PackageManager() string
}
```

Each distro family implements this interface:

| Adapter | Package Manager | Distros |
|---|---|---|
| ArchAdapter | pacman | Arch, CachyOS, EndeavourOS |
| DebianAdapter | apt | Debian, Ubuntu, Pop, Mint, Kali |
| FedoraAdapter | dnf | Fedora, RHEL |

Detection reads `/etc/os-release` at runtime.

## TUI Flow

```
Welcome → Distro Detect → Component Select → Backup Prompt → Install → Done
```

## Symlink Management

Files are NEVER copied. The symlink manager creates `ln -sf` relationships:

```
~/dotfiles/configs/zsh/.zshrc → ~/.zshrc
~/dotfiles/configs/kitty/     → ~/.config/kitty
```

Pre-existing files get a `.bak` timestamp suffix before symlink creation.
