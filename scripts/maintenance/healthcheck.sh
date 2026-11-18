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
echo "--- Enlaces ---"
REPO="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
if [ -x "$REPO/bin/installer" ]; then
    "$REPO/bin/installer" --status
else
    echo "  (compila primero: make build)"
fi
echo ""
echo "--- Tools ---"
TOOLS="zsh tmux nvim fzf bat eza zoxide starship zellij fastfetch"
if [ -d /usr/share/omarchy ] || grep -q '^ID=omarchy' /etc/os-release 2>/dev/null; then
    TOOLS="$TOOLS omarchy hyprctl"
    echo -n "Omarchy: "; omarchy version 2>/dev/null || echo "?"
else
    TOOLS="$TOOLS kitty"
fi
for tool in $TOOLS; do
    if command -v "$tool" &> /dev/null; then
        echo "  ✓ $tool"
    else
        echo "  ✗ $tool"
    fi
done

echo ""
echo "✓ Health check complete"
