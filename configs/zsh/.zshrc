# ================================================
#███████╗░██████╗██╗░░██╗
#╚════██║██╔════╝██║░░██║
#░░███╔═╝╚█████╗░███████║
#██╔══╝░░░╚═══██╗██╔══██║
#███████╗██████╔╝██║░░██║
#╚══════╝╚═════╝░╚═╝░░╚═╝
#
# █▀▀ █▀█ █▄░█ █▀▀ █ █▀▀ █░█ █▀█ ▄▀█ ▀█▀ █ █▀█ █▄░█
# █▄▄ █▄█ █░▀█ █▀░ █ █▄█ █▄█ █▀▄ █▀█ ░█░ █ █▄█ █░▀█
# =================================================
# PACMAN  MODULAR ZSHRC
# =============================================

# ---------------------------------------------
#         POWERLEVEL10K INSTANT PROMPT
# ---------------------------------------------
if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi

# ---------------------------------------------
#              OH-MY-ZSH SETUP
# ---------------------------------------------
export ZSH="$HOME/.oh-my-zsh"

plugins=(
    command-not-found
    git
)

source $ZSH/oh-my-zsh.sh

# ---------------------------------------------
#                 GO PATH
# ---------------------------------------------
export PATH="$PATH:/usr/local/go/bin"

# ---------------------------------------------
#              HOMEBREW SETUP
# ---------------------------------------------
if [[ "$(uname)" == "Darwin" ]]; then
  BREW_BIN="/opt/homebrew/bin"
else
  BREW_BIN="/home/linuxbrew/.linuxbrew/bin"
fi

if [[ -f "$BREW_BIN/brew" ]]; then
  eval "$("$BREW_BIN"/brew shellenv)"
fi

# ---------------------------------------------
#           ZSH PLUGINS (HOMEBREW)
# ---------------------------------------------
[ -f "$(dirname $BREW_BIN)/share/zsh-autocomplete/zsh-autocomplete.plugin.zsh" ] && \
    source "$(dirname $BREW_BIN)/share/zsh-autocomplete/zsh-autocomplete.plugin.zsh"

[ -f "$(dirname $BREW_BIN)/share/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh" ] && \
    source "$(dirname $BREW_BIN)/share/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh"

[ -f "$(dirname $BREW_BIN)/share/zsh-autosuggestions/zsh-autosuggestions.zsh" ] && \
    source "$(dirname $BREW_BIN)/share/zsh-autosuggestions/zsh-autosuggestions.zsh"

[ -f "$(dirname $BREW_BIN)/share/powerlevel10k/powerlevel10k.zsh-theme" ] && \
    source "$(dirname $BREW_BIN)/share/powerlevel10k/powerlevel10k.zsh-theme"

# ---------------------------------------------
#      LOAD MODULAR CONFIGS FROM .DOTFILES
# ---------------------------------------------
DOTFILES_ZSH="$HOME/.dotfiles/zsh"

[ -f "$DOTFILES_ZSH/exports.zsh" ] && source "$DOTFILES_ZSH/exports.zsh"
[ -f "$DOTFILES_ZSH/paths.zsh" ] && source "$DOTFILES_ZSH/paths.zsh"
[ -f "$DOTFILES_ZSH/functions.zsh" ] && source "$DOTFILES_ZSH/functions.zsh"
[ -f "$DOTFILES_ZSH/aliases.zsh" ] && source "$DOTFILES_ZSH/aliases.zsh"

# ---------------------------------------------
#           EXTERNAL TOOLS INIT
# ---------------------------------------------
command -v carapace &> /dev/null && source <(carapace _carapace)
command -v fzf &> /dev/null && eval "$(fzf --zsh)"
command -v zoxide &> /dev/null && eval "$(zoxide init zsh)"
command -v atuin &> /dev/null && eval "$(atuin init zsh)"

# ---------------------------------------------
#         POWERLEVEL10K THEME CONFIG
# ---------------------------------------------
[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh

# ---------------------------------------------
#          WINDOW MANAGER AUTOSTART
# ---------------------------------------------
start_if_needed

# ---------------------------------------------
#              STARTUP BANNER
# ---------------------------------------------
if [ -f "$HOME/.dotfiles/banner.txt" ]; then
    clear
    command -v lolcat &> /dev/null && cat "$HOME/.dotfiles/banner.txt" | lolcat || cat "$HOME/.dotfiles/banner.txt"
fi
