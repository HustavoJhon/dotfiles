#!/usr/bin/env python3
"""Theme switcher: switch terminal/editor/system themes."""

import json
import os
import shutil
from pathlib import Path

THEMES = {
    "tokyo-night": {
        "name": "Tokyo Night",
        "kitty": "color.ini",
        "bat": "tokyonight",
        "fzf": "tokyonight",
    },
    "catppuccin-mocha": {
        "name": "Catppuccin Mocha",
        "kitty": "catppuccin-mocha.ini",
        "bat": "Catppuccin Mocha",
        "fzf": "catppuccin-mocha",
    },
    "nord": {
        "name": "Nord",
        "kitty": "themes/nord.ini",
        "bat": "Nord",
        "fzf": "nord",
    },
}

DOTFILES_DIR = Path(os.getenv("DOTFILES_DIR", "."))

KITTY_CONFIG_DIR = DOTFILES_DIR / "configs" / "kitty"
KITTY_TARGET = Path.home() / ".config" / "kitty"

BAT_CONFIG = Path.home() / ".config" / "bat" / "config"


def list_themes():
    print("Available themes:")
    for key, theme in THEMES.items():
        indicator = " *" if is_active(theme) else ""
        print(f"  {key}: {theme['name']}{indicator}")


def is_active(theme: dict) -> bool:
    kitty_conf = KITTY_TARGET / "kitty.conf"
    if kitty_conf.exists():
        content = kitty_conf.read_text()
        return theme["kitty"] in content
    return False


def switch_kitty(theme_key: str):
    theme = THEMES[theme_key]
    kitty_conf = KITTY_CONFIG_DIR / "kitty.conf"
    if not kitty_conf.exists():
        print("Kitty config not found, skipping.")
        return

    lines = kitty_conf.read_text().splitlines()
    new_lines = []
    for line in lines:
        if line.strip().startswith("include"):
            new_lines.append(f"include {theme['kitty']}")
        else:
            new_lines.append(line)
    kitty_conf.write_text("\n".join(new_lines))
    print(f"  ✓ Kitty theme set to {theme['name']}")


def switch_bat(theme_key: str):
    theme = THEMES[theme_key]
    BAT_CONFIG.parent.mkdir(parents=True, exist_ok=True)
    BAT_CONFIG.write_text(f'--theme="{theme["bat"]}"\n')
    print(f"  ✓ Bat theme set to {theme['bat']}")


def apply_all(theme_key: str):
    if theme_key not in THEMES:
        print(f"Unknown theme: {theme_key}")
        print(f"Available: {', '.join(THEMES.keys())}")
        return

    theme = THEMES[theme_key]
    print(f"Applying {theme['name']}...")
    switch_kitty(theme_key)
    switch_bat(theme_key)
    print(f"\n✓ Theme switched to {theme['name']}")
    print("  Reload Kitty to see changes.")


def main():
    import argparse
    parser = argparse.ArgumentParser(description="Theme Switcher")
    parser.add_argument("action", nargs="?", default="list",
                        choices=["list", "apply"])
    parser.add_argument("theme", nargs="?")
    args = parser.parse_args()

    match args.action:
        case "list":
            list_themes()
        case "apply":
            if not args.theme:
                print("Usage: theme_switcher.py apply <theme_name>")
                print(f"Themes: {', '.join(THEMES.keys())}")
                return
            apply_all(args.theme)


if __name__ == "__main__":
    main()
