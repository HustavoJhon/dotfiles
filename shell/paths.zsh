#!/usr/bin/env zsh
# =============================================
#           EXTRA SOFTWARE PATHS
# =============================================

# ---------------------------------------------
#              CUSTOM BINS
# ---------------------------------------------
export PATH="$HOME/.local/bin:$PATH"
export PATH="$HOME/.opencode/bin:$PATH"

# ---------------------------------------------
#              RUST / CARGO
# ---------------------------------------------
export PATH="$HOME/.cargo/bin:$PATH"

# ---------------------------------------------
#              NODE / VOLTA
# ---------------------------------------------
export PATH="$HOME/.volta/bin:$PATH"

# ---------------------------------------------
#              BUN
# ---------------------------------------------
export PATH="$HOME/.bun/bin:$PATH"
export BUN_INSTALL="$HOME/.bun"

# ---------------------------------------------
#              NIX
# ---------------------------------------------
export PATH="$HOME/.nix-profile/bin:$PATH"
export PATH="/nix/var/nix/profiles/default/bin:$PATH"

# ---------------------------------------------
#              FLUTTER
# ---------------------------------------------
export PATH="$PATH:/opt/flutter/bin"

# ---------------------------------------------
#              PNPM
# ---------------------------------------------
export PATH="$PATH:$HOME/.local/share/pnpm"
export PNPM_HOME="$HOME/.local/share/pnpm"

# ---------------------------------------------
#              GO
# ---------------------------------------------
export GOPATH=/usr/share/go
export PATH="$PATH:$GOPATH/bin"

# ---------------------------------------------
#              CUSTOM CONFIG PATH
# ---------------------------------------------
export PATH="$HOME/.config:$PATH"

# ---------------------------------------------
#              LOCAL LIB PATH
# ---------------------------------------------
export PATH="/usr/local/lib/*:$PATH"
export PATH="/usr/local/bin:$PATH"

# ---------------------------------------------
#              PROJECT PATHS
# ---------------------------------------------
export PROJECT_PATHS="$HOME/work"
export PROJECTS="$HOME/work"

# Shortcuts to common directories
export DOWNLOADS="$HOME/Downloads"
export DOCUMENTS="$HOME/Documents"
export DESKTOP="$HOME/Desktop"
