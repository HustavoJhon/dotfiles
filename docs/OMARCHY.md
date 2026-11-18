# PACMAN en Omarchy

[Omarchy](https://omarchy.org) es Arch + Hyprland con su propio sistema de temas,
barra (Quickshell) y config de Hyprland en **Lua**. PACMAN lo detecta
(`ID=omarchy` en `/etc/os-release`) y cambia de modo:

| En Omarchy | Fuera de Omarchy |
|---|---|
| módulo **omarchy** (overrides en Lua) | módulo **hyprland** (`wm/hypr`, formato `.conf`) |
| kitty / ghostty / fastfetch que siguen el tema activo | kitty / ghostty / fastfetch con tema propio |

Así nunca se pisa `~/.config/hypr` entero: solo se enlazan los archivos de
usuario que Omarchy carga **después** de sus defaults
(`/usr/share/omarchy/default/hypr`), y `omarchy update` sigue funcionando.

## Qué instala el módulo

```
omarchy/
├── hypr/
│   ├── hyprland.lua      → ~/.config/hypr/hyprland.lua   (carga defaults + overrides)
│   ├── bindings.lua      → atajos propios (zellij, tmux, apps, webapps)
│   ├── input.lua         → teclado us, compose en Caps, repeat 40/250, touchpad
│   ├── looknfeel.lua     → bordes redondeados
│   ├── autostart.lua
│   ├── hypridle.conf     → sin apagado/suspensión automática
│   └── hyprlock.conf     → pantalla de bloqueo con el tema actual
├── config/
│   ├── shell.json        → ~/.config/omarchy/shell.json (barra, widgets, idle)
│   └── shell.toml        → tamaño de fuente del shell
├── branding/
│   ├── about.txt         → logo PACMAN en fastfetch / About
│   └── screensaver.txt   → salvapantallas PACMAN
├── themes/wallhaven-*    → ~/.config/omarchy/themes/ (temas propios)
├── kitty/kitty.conf      → incluye ~/.local/state/omarchy/current/theme/kitty.conf
├── ghostty/config        → incluye el ghostty.conf del tema
└── fastfetch/config.jsonc
```

`monitors.lua` **no** se versiona: depende de cada máquina.

## Uso

```bash
make omarchy            # solo el módulo omarchy, sin TUI
make status             # ver qué está enlazado
make run                # TUI (omarchy viene preseleccionado)
```

Después:

```bash
omarchy theme set wallhaven-gjr6ke   # aplicar un tema PACMAN
fastfetch                            # ver el logo nuevo
omarchy restart terminal
```

El instalador ejecuta `hyprctl reload` y muestra `hyprctl configerrors` si
Hyprland está corriendo.

## Mantener el repo al día

Algunas herramientas reescriben sus archivos (p. ej. `omarchy bar ...` en
`shell.json`, o `cloofy` en `bindings.lua`) y pueden reemplazar el symlink por
un archivo normal. Para traer esos cambios al repo y volver a enlazar:

```bash
scripts/omarchy/sync.sh --check   # ver diferencias
make omarchy-sync                 # copiar al repo + re-enlazar
git diff omarchy/
```

## Dependencias externas

`shell.json` usa widgets de plugins de terceros (`cloofy.nowplaying`,
`canberkbekir.omarchy-remote`, `io.voxtype.settings`). No se versionan; en una
máquina nueva instálalos con sus propios instaladores o quítalos con
`omarchy bar remove <id>`.

## Volver a los valores de Omarchy

```bash
omarchy refresh hyprland   # restaura ~/.config/hypr (hace backup)
omarchy refresh shell
```

Los archivos que había antes de enlazar quedan en `~/.dotfiles-backup/links/<fecha>/` y el
backup completo en `~/.dotfiles-backup/<fecha>/`.
