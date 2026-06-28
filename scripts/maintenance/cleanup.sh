#!/bin/bash
# PACMAN - System Cleanup

echo "PACMAN Cleanup"

clean_pacman() {
    echo "→ Cleaning pacman cache..."
    sudo pacman -Sc --noconfirm 2>/dev/null || true
    sudo paccache -rk 2 2>/dev/null || true
    echo "→ Checking orphaned packages..."
    pacman -Qtdq 2>/dev/null | sudo pacman -Rns --noconfirm - 2>/dev/null || true
}

clean_apt() {
    echo "→ Cleaning apt cache..."
    sudo apt-get autoremove -y 2>/dev/null || true
    sudo apt-get autoclean -y 2>/dev/null || true
}

clean_dnf() {
    echo "→ Cleaning dnf cache..."
    sudo dnf clean all 2>/dev/null || true
    sudo dnf autoremove -y 2>/dev/null || true
}

clean_temp() {
    echo "→ Cleaning temporary files..."
    rm -rf "$HOME/.cache/thumbnails/"* 2>/dev/null || true
    rm -rf /tmp/* 2>/dev/null || true
}

clean_journal() {
    echo "→ Cleaning journal logs..."
    sudo journalctl --vacuum-size=100M 2>/dev/null || true
}

detect_pm() {
    if command -v pacman &> /dev/null; then clean_pacman; fi
    if command -v apt-get &> /dev/null; then clean_apt; fi
    if command -v dnf &> /dev/null; then clean_dnf; fi
}

detect_pm
clean_temp
clean_journal

echo "✓ Cleanup complete"
