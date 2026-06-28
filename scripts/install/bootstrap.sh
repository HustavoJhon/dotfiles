#!/bin/bash
# PACMAN - Bootstrap Installer
# Installs Go (if missing) and runs the TUI installer

set -e

BOLD='\033[1m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${PURPLE}${BOLD}"
echo "  PACMAN"
echo "  Bootstrap Installer"
echo -e "${NC}"

# Check Go
if ! command -v go &> /dev/null; then
    echo -e "${CYAN}→ Installing Go...${NC}"
    GOVERSION="1.22.2"
    wget -q "https://go.dev/dl/go${GOVERSION}.linux-amd64.tar.gz" -O /tmp/go.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf /tmp/go.tar.gz
    rm /tmp/go.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    echo -e "${CYAN}✓ Go installed${NC}"
fi

echo -e "${CYAN}→ Building installer...${NC}"
cd "$(dirname "$0")/.."
go build -o bin/installer ./cmd/installer

echo -e "${PURPLE}→ Running TUI installer...${NC}"
./bin/installer
