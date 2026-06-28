#!/bin/bash
# PACMAN - System Health Check

echo "PACMAN Health Check"
echo "======================="

echo -n "OS: "
if [ -f /etc/os-release ]; then
    . /etc/os-release
    echo "$PRETTY_NAME"
else
    echo "Unknown"
fi

echo -n "Kernel: "
uname -r

echo -n "Uptime: "
uptime -p | sed 's/up //'

echo -n "Shell: "
basename "$SHELL"

echo -n "CPU: "
lscpu 2>/dev/null | grep "Model name" | head -1 | cut -d: -f2 | xargs || echo "N/A"

echo -n "Memory: "
free -h | awk '/^Mem:/ {print $3 "/" $2}'

echo -n "Disk: "
df -h / | awk 'NR==2 {print $3 "/" $2 " (" $5 ")"}'

echo ""
echo "--- Symlinks ---"
for f in "$HOME/.zshrc" "$HOME/.tmux.conf" "$HOME/.p10k.zsh"; do
    if [ -L "$f" ]; then
        echo "  ✓ $f -> $(readlink -f "$f")"
    elif [ -f "$f" ]; then
        echo "  ! $f (regular file, not symlink)"
    else
        echo "  ✗ $f (missing)"
    fi
done

echo ""
echo "--- Tools ---"
for tool in zsh kitty tmux nvim fzf bat eza zoxide starship zellij fastfetch; do
    if command -v "$tool" &> /dev/null; then
        echo "  ✓ $tool"
    else
        echo "  ✗ $tool"
    fi
done

echo ""
echo "✓ Health check complete"
