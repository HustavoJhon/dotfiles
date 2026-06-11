#!/usr/bin/env python3
"""Update font cache and verify Nerd Font installations."""

import os
import shutil
import subprocess
import sys
from pathlib import Path


FONT_DIRS = [
    Path.home() / ".local/share/fonts",
    Path("/usr/share/fonts"),
    Path("/usr/local/share/fonts"),
]

NERD_FONT_SIG = "Nerd Font"


def get_dotfiles_fonts() -> list[Path]:
    dotfiles_root = Path(os.getenv("DOTFILES_DIR", "."))
    fonts_dir = dotfiles_root / "assets" / "fonts"
    if not fonts_dir.exists():
        print(f"Fonts directory not found: {fonts_dir}")
        return []
    return list(fonts_dir.glob("**/*.ttf")) + list(fonts_dir.glob("**/*.otf"))


def install_fonts(fonts: list[Path], target: Path) -> int:
    target.mkdir(parents=True, exist_ok=True)
    count = 0
    for font in fonts:
        dest = target / font.name
        if not dest.exists():
            shutil.copy2(font, dest)
            print(f"  + {font.name}")
            count += 1
    return count


def update_cache():
    subprocess.run(["fc-cache", "-f"], check=False)


def list_installed_nerd_fonts() -> list[str]:
    result = subprocess.run(
        ["fc-list", f":family={NERD_FONT_SIG}"],
        capture_output=True, text=True, check=False,
    )
    return [line.strip() for line in result.stdout.split("\n") if line.strip()]


def main():
    fonts = get_dotfiles_fonts()
    if not fonts:
        print("No fonts found in assets/fonts/")
        sys.exit(0)

    target = FONT_DIRS[0]
    installed = install_fonts(fonts, target)

    if installed > 0:
        print(f"Installed {installed} new fonts. Updating cache...")
        update_cache()
    else:
        print("All fonts already installed.")

    nerds = list_installed_nerd_fonts()
    print(f"\nNerd Fonts available: {len(nerds)}")
    for nf in nerds[:10]:
        print(f"  - {nf}")


if __name__ == "__main__":
    main()
