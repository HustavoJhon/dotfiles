#!/usr/bin/env python3
"""Wallpaper manager: set, rotate, and manage wallpapers."""

import os
import random
import subprocess
import sys
from pathlib import Path


WALLPAPER_DIR = Path(os.getenv("DOTFILES_DIR", ".")) / "assets" / "wallpapers"
CACHE_FILE = Path.home() / ".cache" / "wallpaper-current"


def get_wallpapers() -> list[Path]:
    if not WALLPAPER_DIR.exists():
        print(f"Wallpaper directory not found: {WALLPAPER_DIR}")
        return []
    return sorted(WALLPAPER_DIR.glob("*.*"))


def set_hyprpaper(wallpaper: Path):
    cmd = ["hyprctl", "hyprpaper", "wallpaper", f",{wallpaper}"]
    subprocess.run(cmd, check=False)


def set_feh(wallpaper: Path):
    cmd = ["feh", "--bg-fill", str(wallpaper)]
    subprocess.run(cmd, check=False)


def set_gsettings(wallpaper: Path):
    uri = Path(wallpaper).as_uri()
    cmd = ["gsettings", "set", "org.gnome.desktop.background", "picture-uri", uri]
    subprocess.run(cmd, check=False)


def detect_wm() -> str:
    if os.environ.get("HYPRLAND_INSTANCE_SIGNATURE"):
        return "hyprland"
    if os.environ.get("XDG_CURRENT_DESKTOP", "").lower() in ("gnome", "unity"):
        return "gnome"
    if os.environ.get("DESKTOP_SESSION", "").lower() in ("i3", "sway", "bspwm"):
        return "other"
    return "unknown"


def set_wallpaper(wallpaper: Path):
    wm = detect_wm()
    print(f"Detected WM: {wm}")

    match wm:
        case "hyprland":
            set_hyprpaper(wallpaper)
        case "gnome":
            set_gsettings(wallpaper)
        case _:
            if subprocess.run(["which", "feh"], capture_output=True).returncode == 0:
                set_feh(wallpaper)
            else:
                print("No wallpaper setter found (try: hyprpaper, feh)")
                return

    CACHE_FILE.parent.mkdir(parents=True, exist_ok=True)
    CACHE_FILE.write_text(str(wallpaper))
    print(f"Wallpaper set: {wallpaper.name}")


def main():
    import argparse
    parser = argparse.ArgumentParser(description="Wallpaper Manager")
    parser.add_argument("action", nargs="?", default="random",
                        choices=["random", "next", "list", "current"])
    args = parser.parse_args()

    wallpapers = get_wallpapers()
    if not wallpapers:
        print("No wallpapers found.")
        sys.exit(1)

    match args.action:
        case "list":
            for w in wallpapers:
                print(w.name)
        case "random":
            set_wallpaper(random.choice(wallpapers))
        case "next":
            if CACHE_FILE.exists():
                current = Path(CACHE_FILE.read_text().strip())
                try:
                    idx = wallpapers.index(current)
                    next_w = wallpapers[(idx + 1) % len(wallpapers)]
                except ValueError:
                    next_w = wallpapers[0]
            else:
                next_w = wallpapers[0]
            set_wallpaper(next_w)
        case "current":
            if CACHE_FILE.exists():
                print(CACHE_FILE.read_text().strip())
            else:
                print("No current wallpaper cached.")


if __name__ == "__main__":
    main()
