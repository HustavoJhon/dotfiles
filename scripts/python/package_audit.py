#!/usr/bin/env python3
"""Audit installed packages against desired package list."""

import subprocess
import sys
from pathlib import Path

PACKAGE_LISTS = {
    "core": [
        "zsh", "git", "curl", "wget", "unzip", "zip",
        "neovim", "tmux", "fzf", "ripgrep", "fd",
        "bat", "zoxide", "eza",
    ],
    "terminal": [
        "kitty", "zellij", "starship",
    ],
    "dev": [
        "nodejs", "npm", "python3", "go", "rustup",
        "docker", "docker-compose",
    ],
    "media": [
        "ffmpeg", "imagemagick", "mpv",
    ],
    "hypr": [
        "hyprland", "hyprlock", "hypridle", "waybar",
        "wofi", "dunst", "swaync", "hyprpaper",
    ],
}


def detect_distro() -> str:
    try:
        with open("/etc/os-release") as f:
            for line in f:
                if line.startswith("ID="):
                    return line.strip().split("=")[1].strip('"')
    except FileNotFoundError:
        pass
    return "unknown"


def get_installed_packages() -> dict[str, str]:
    distro = detect_distro()
    pkgs = {}

    match distro:
        case "arch" | "cachyos" | "endeavouros":
            result = subprocess.run(
                ["pacman", "-Q"], capture_output=True, text=True, check=False
            )
            for line in result.stdout.splitlines():
                parts = line.split()
                if len(parts) >= 1:
                    pkgs[parts[0]] = parts[1] if len(parts) > 1 else "?"

        case "debian" | "ubuntu" | "pop":
            result = subprocess.run(
                ["dpkg", "-l"], capture_output=True, text=True, check=False
            )
            for line in result.stdout.splitlines():
                parts = line.split()
                if len(parts) >= 2 and parts[0] == "ii":
                    pkgs[parts[1]] = parts[2]

        case "fedora":
            result = subprocess.run(
                ["rpm", "-qa"], capture_output=True, text=True, check=False
            )
            for line in result.stdout.splitlines():
                name = line.rsplit("-", 2)[0]
                pkgs[name] = line

    return pkgs


def audit():
    print(f"Distribution: {detect_distro()}")
    installed = get_installed_packages()
    print(f"Total installed packages: {len(installed)}\n")

    total_needed = 0
    total_missing = 0

    for category, pkgs in PACKAGE_LISTS.items():
        missing = [p for p in pkgs if p not in installed]
        total_needed += len(pkgs)
        total_missing += len(missing)

        status = "✓" if not missing else f"✗ ({len(missing)}/{len(pkgs)} missing)"
        print(f"  [{status}] {category}:")

        for pkg in pkgs:
            if pkg in installed:
                print(f"      ✓ {pkg} ({installed[pkg]})")
            else:
                print(f"      ✗ {pkg} (not installed)")

        print()

    pct = ((total_needed - total_missing) / total_needed * 100) if total_needed else 100
    print(f"System readiness: {total_needed - total_missing}/{total_needed} ({pct:.0f}%)")

    return total_missing


def main():
    missing = audit()
    sys.exit(1 if missing > 0 else 0)


if __name__ == "__main__":
    main()
