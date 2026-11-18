<div align="center">
  <img src="assets/logo/ghost.png" alt="PACMAN" width="110" height="110">
  <h1>PACMAN</h1>
  <p><b>dotfiles con instalador TUI · edición Omarchy</b></p>
  <p>
    <img src="https://img.shields.io/badge/Omarchy-4.x-8ec07c?style=flat" alt="Omarchy">
    <img src="https://img.shields.io/badge/Hyprland-Lua-58E1FF?style=flat&logo=hyprland&logoColor=white" alt="Hyprland">
    <img src="https://img.shields.io/badge/Go-Bubble%20Tea-00ADD8?style=flat&logo=go" alt="Go">
    <img src="https://img.shields.io/badge/Linux-Arch%20·%20Debian%20·%20Fedora-FCC624?style=flat&logo=linux&logoColor=black" alt="Linux">
    <img src="https://img.shields.io/badge/License-GPLv3-blue" alt="License">
  </p>
</div>

```bash
curl -fsSL https://raw.githubusercontent.com/hustavojhon/dotfiles/main/install.sh | bash
```

Clona el repo en `~/.local/share/pacman-dotfiles`, compila el instalador y abre
la TUI. Nada se sobrescribe sin backup.

---

## Omarchy

En [Omarchy](https://omarchy.org) el instalador lo detecta solo y activa el
módulo **omarchy**: overrides de Hyprland en Lua, barra, temas propios y la
presentación PACMAN — sin pisar los defaults de Omarchy, así que
`omarchy update` sigue funcionando.

```
          ████████
      ████████████████            fastfetch / About
    ████████████████████          ─────────────────
  ██        ████        ██        logo:   omarchy/branding/about.txt
  ██        ████        ██        screen: omarchy/branding/screensaver.txt
████  ████  ████  ████  ████      temas:  wallhaven-3z7dd3 · gjr6ke · gwdvkd
████  ████  ████  ████  ████
██████    ████████    ██████
████████████████████████████
████████████████████████████
████  ██████    ██████  ████
██      ████    ████      ██

  ●   ●   ●   ●   ●   ●   ●
```

| Qué | Dónde |
|---|---|
| Hyprland (Lua): atajos, input, apariencia, idle, lock | `omarchy/hypr/` |
| Barra / widgets / idle del shell | `omarchy/config/shell.json` |
| Logo y salvapantallas PACMAN | `omarchy/branding/` |
| Temas propios con fondos | `omarchy/themes/` |
| Kitty · Ghostty · Fastfetch que siguen el tema | `omarchy/{kitty,ghostty,fastfetch}/` |

```bash
make omarchy                         # instalar solo el módulo
omarchy theme set wallhaven-gjr6ke   # aplicar un tema PACMAN
make omarchy-sync                    # traer al repo cambios hechos en vivo
```

Guía completa: [docs/OMARCHY.md](docs/OMARCHY.md)

---

## Módulos

| Módulo | Qué incluye | Plataforma |
|---|---|---|
| **Omarchy** | Hyprland Lua, shell, temas, branding, terminales | Omarchy |
| **ZSH** | Shell modular, aliases, fzf, zoxide, Powerlevel10k | todas |
| **Zellij** · **Tmux** | Multiplexores con layouts | todas |
| **Starship** | Prompt multi-shell | todas |
| **WezTerm** | Terminal en Lua | todas |
| **Kitty** · **Ghostty** · **Fastfetch** | Con tema propio | fuera de Omarchy |
| **Hyprland** | Config clásica `.conf` | fuera de Omarchy |
| **Neovim** | LazyVim | todas (opcional) |

Detalle de rutas y paquetes: [docs/MODULES.md](docs/MODULES.md)

## Comandos

```bash
make run        # TUI interactiva
make dry-run    # TUI sin tocar nada
make list       # módulos disponibles en este sistema
make status     # qué está enlazado, ausente o en conflicto
make backup     # backup completo en ~/.dotfiles-backup/
make test       # tests Go
make e2e        # tests end-to-end en Docker (Ubuntu limpio)
make help       # todo lo demás
```

Sin TUI (scripts, CI, otra máquina):

```bash
bin/installer --install omarchy,zsh,zellij   # o --install all
bin/installer --install zsh --no-packages    # solo enlaces
bin/installer --install all --dry-run        # ver qué haría
```

## Cómo funciona

- **Symlinks**, nunca copias: editar en el repo = editar la config viva.
- Si ya existe un archivo, se renombra a `*.bak.<fecha>` antes de enlazar. Reinstalar es idempotente.
- Paquetes uno por uno: si uno no existe en tu distro es un aviso, no un error.
  Si el binario ya está en el `PATH` (Homebrew, mise…) no se reinstala.
- `sudo` se pide una sola vez y solo si falta algún paquete.
- Distros: Arch / Omarchy / CachyOS / EndeavourOS / Manjaro (pacman + AUR), Debian / Ubuntu / Mint / Pop (apt), Fedora (dnf).

## Estructura

```
omarchy/      Overrides de Omarchy (Hyprland Lua, shell, temas, branding)
shell/        ZSH modular
term/         Kitty, WezTerm, Ghostty, Fastfetch
wm/           Hyprland (.conf), Tmux, Zellij
editor/       Neovim
prompt/       Starship
cmd/ internal/  Instalador Go (Bubble Tea + Lip Gloss)
scripts/      Instalación, mantenimiento, sync de Omarchy
e2e/          Tests end-to-end
assets/ docs/ shaders/ windows/
```

## Licencia

GNU GPL v3
