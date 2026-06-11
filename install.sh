#!/bin/bash
# PACMAN.DOTS - Quick Install Script
# One-liner: curl -fsSL https://dotfiles.example.com/install.sh | bash

set -e

BOLD='\033[1m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${PURPLE}${BOLD}"
echo "██████╗  █████╗  ██████╗███╗   ███╗ █████╗ ███╗   ██╗"
echo "██╔══██╗██╔══██╗██╔════╝████╗ ████║██╔══██╗████╗  ██║"
echo "██████╔╝███████║██║     ██╔████╔██║███████║██╔██╗ ██║"
echo "██╔═══╝ ██╔══██║██║     ██║╚██╔╝██║██╔══██║██║╚██╗██║"
echo "██║     ██║  ██║╚██████╗██║ ╚═╝ ██║██║  ██║██║ ╚████║"
echo "╚═╝     ╚═╝  ╚═╝ ╚═════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝"
echo -e "${NC}"
echo -e "${BOLD}PACMAN.DOTS - Dotfiles Installer${NC}"
echo ""

# Get dotfiles directory
DOTFILES_DIR="$(cd "$(dirname "$0")" && pwd)"

# Check for Go
if ! command -v go &> /dev/null; then
    echo -e "${CYAN}→ Installing Go...${NC}"
    GOVERSION="1.22.2"
    if command -v wget &> /dev/null; then
        wget -q "https://go.dev/dl/go${GOVERSION}.linux-amd64.tar.gz" -O /tmp/go.tar.gz
    elif command -v curl &> /dev/null; then
        curl -sL "https://go.dev/dl/go${GOVERSION}.linux-amd64.tar.gz" -o /tmp/go.tar.gz
    else
        echo "Error: need wget or curl to install Go"
        exit 1
    fi
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf /tmp/go.tar.gz
    rm /tmp/go.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    echo -e "${CYAN}✓ Go installed${NC}"
fi

echo -e "${CYAN}→ Building installer...${NC}"
cd "$DOTFILES_DIR"
go build -o bin/installer ./cmd/installer

echo -e "${PURPLE}→ Starting TUI installer...${NC}"
./bin/installer
