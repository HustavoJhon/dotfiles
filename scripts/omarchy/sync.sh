#!/usr/bin/env bash
# Copia la config viva de Omarchy (~/.config) al repo.
# Útil cuando una herramienta (omarchy bar, cloofy, ...) reescribe un archivo
# y reemplaza el symlink por un archivo normal. Tras copiar, vuelve a enlazar.
#   scripts/omarchy/sync.sh          # copiar y re-enlazar
#   scripts/omarchy/sync.sh --check  # solo mostrar diferencias
set -euo pipefail

REPO="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
CHECK=0
[[ "${1:-}" == "--check" ]] && CHECK=1

# repo:destino (relativo a $HOME)
MAP=(
    omarchy/hypr/hyprland.lua:.config/hypr/hyprland.lua
    omarchy/hypr/bindings.lua:.config/hypr/bindings.lua
    omarchy/hypr/input.lua:.config/hypr/input.lua
    omarchy/hypr/looknfeel.lua:.config/hypr/looknfeel.lua
    omarchy/hypr/autostart.lua:.config/hypr/autostart.lua
    omarchy/hypr/hypridle.conf:.config/hypr/hypridle.conf
    omarchy/hypr/hyprlock.conf:.config/hypr/hyprlock.conf
    omarchy/config/shell.json:.config/omarchy/shell.json
    omarchy/config/shell.toml:.config/omarchy/shell.toml
    omarchy/kitty/kitty.conf:.config/kitty/kitty.conf
    omarchy/ghostty/config:.config/ghostty/config
    omarchy/fastfetch/config.jsonc:.config/fastfetch/config.jsonc
)

changed=0
for entry in "${MAP[@]}"; do
    src="$REPO/${entry%%:*}" dest="$HOME/${entry#*:}"
    if [[ -L "$dest" && "$(readlink "$dest")" == "$src" ]]; then
        continue                                   # enlazado, nada que hacer
    fi
    [[ -f "$dest" ]] || { echo "  - falta ~/${entry#*:}"; continue; }
    if cmp -s "$dest" "$src"; then
        echo "  = ~/${entry#*:} (igual, sin enlazar)"
    else
        echo "  ≠ ~/${entry#*:}"
        diff -u "$src" "$dest" | sed 's/^/      /' | head -40 || true
        changed=1
    fi
    if (( ! CHECK )); then
        cp "$dest" "$src"
        ln -sfn "$src" "$dest"
        echo "    → copiado al repo y re-enlazado"
    fi
done

if (( CHECK )); then
    (( changed )) && echo "Hay cambios: ejecuta sin --check para traerlos." || echo "✓ Todo sincronizado"
else
    echo "✓ Revisa con: git -C \"$REPO\" diff omarchy/"
fi
