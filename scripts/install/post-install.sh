#!/bin/bash
# PACMAN.DOTS - Post-Installation Script
# Run after the TUI installer to finalize setup

set -e

echo "PACMAN.DOTS Post-Install"

# Change shell to ZSH if not already
if [ "$SHELL" != "$(which zsh)" ]; then
    echo "→ Changing shell to zsh..."
    chsh -s "$(which zsh)"
fi

# Install Oh-My-Zsh if not present
if [ ! -d "$HOME/.oh-my-zsh" ]; then
    echo "→ Installing Oh-My-Zsh..."
    sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended
fi

# Create .dotfiles structure
mkdir -p "$HOME/.dotfiles/zsh"

# Ensure zsh plugin directories exist
ZSH_CUSTOM="${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}"
mkdir -p "$ZSH_CUSTOM/plugins"

# Install zsh plugins via Homebrew or git
if command -v brew &> /dev/null; then
    brew install zsh-autocomplete zsh-autosuggestions zsh-syntax-highlighting powerlevel10k 2>/dev/null || true
else
    # Fallback: install via git
    [ ! -d "$ZSH_CUSTOM/plugins/zsh-autosuggestions" ] && \
        git clone --depth=1 https://github.com/zsh-users/zsh-autosuggestions "$ZSH_CUSTOM/plugins/zsh-autosuggestions" 2>/dev/null || true
    [ ! -d "$ZSH_CUSTOM/plugins/zsh-syntax-highlighting" ] && \
        git clone --depth=1 https://github.com/zsh-users/zsh-syntax-highlighting.git "$ZSH_CUSTOM/plugins/zsh-syntax-highlighting" 2>/dev/null || true
fi

echo ""
echo "✓ Post-install complete!"
echo "→ Restart your shell or run: source ~/.zshrc"
