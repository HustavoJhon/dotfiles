#!/usr/bin/env zsh
# =============================================
#          ENVIRONMENT EXPORTS
# =============================================

# ---------------------------------------------
#              EDITOR CONFIG
# ---------------------------------------------
export EDITOR='nvim'
export VISUAL='nvim'
export PAGER='less'

# ---------------------------------------------
#              LOCALE
# ---------------------------------------------
export LANG='en_US.UTF-8'
export LC_ALL='en_US.UTF-8'

# ---------------------------------------------
#              LS COLORS
# ---------------------------------------------
export LS_COLORS="di=38;5;67:ow=48;5;60:ex=38;5;132:ln=38;5;144:*.tar=38;5;180:*.zip=38;5;180:*.jpg=38;5;175:*.png=38;5;175:*.mp3=38;5;175:*.wav=38;5;175:*.txt=38;5;223:*.sh=38;5;132"

# ---------------------------------------------
#              EXA COLORS
# ---------------------------------------------
export EXA_COLORS="da=38;5;252:sb=38;5;204:sn=38;5;43:uu=38;5;245:un=38;5;241:ur=38;5;40:uw=38;5;203:ux=38;5;76:ue=38;5;76:gr=38;5;40:gw=38;5;203:gx=38;5;76:tr=38;5;40:tw=38;5;203:tx=38;5;76"

# ---------------------------------------------
#              FZF CONFIG
# ---------------------------------------------
export FZF_DEFAULT_COMMAND="fd --hidden --strip-cwd-prefix --exclude .git"
export FZF_DEFAULT_T_COMMAND="$FZF_DEFAULT_COMMAND"
export FZF_ALT_COMMAND="fd --type=d --hidden --strip-cwd-prefix --exclude .git"

# FZF Theme (Gruvbox)
export FZF_DEFAULT_OPTS="
--height 40%
--layout=reverse
--border
--inline-info
--color=fg:#ebdbb2,bg:#282828,hl:#fabd2f
--color=fg+:#ebdbb2,bg+:#3c3836,hl+:#fabd2f
--color=info:#83a598,prompt:#bdae93,pointer:#fe8019
--color=marker:#fe8019,spinner:#fabd2f,header:#665c54
"

# ---------------------------------------------
#              BAT CONFIG
# ---------------------------------------------
export BAT_THEME="gruvbox-dark"
export BAT_STYLE="numbers,changes,header"

# ---------------------------------------------
#              LESS CONFIG
# ---------------------------------------------
export LESS='-R -M -i -j10 --use-color'
export LESS_TERMCAP_mb=$'\e[1;32m'
export LESS_TERMCAP_md=$'\e[1;32m'
export LESS_TERMCAP_me=$'\e[0m'
export LESS_TERMCAP_se=$'\e[0m'
export LESS_TERMCAP_so=$'\e[01;33m'
export LESS_TERMCAP_ue=$'\e[0m'
export LESS_TERMCAP_us=$'\e[1;4;31m'

# ---------------------------------------------
#              HISTORY CONFIG
# ---------------------------------------------
export HISTFILE="$HOME/.zsh_history"
export HISTSIZE=50000
export SAVEHIST=50000
setopt EXTENDED_HISTORY
setopt INC_APPEND_HISTORY
setopt SHARE_HISTORY
setopt HIST_EXPIRE_DUPS_FIRST
setopt HIST_IGNORE_DUPS
setopt HIST_IGNORE_ALL_DUPS
setopt HIST_FIND_NO_DUPS
setopt HIST_IGNORE_SPACE
setopt HIST_SAVE_NO_DUPS
setopt HIST_VERIFY

# ---------------------------------------------
#              COMPLETION CONFIG
# ---------------------------------------------
setopt ALWAYS_TO_END
setopt AUTO_MENU
setopt COMPLETE_IN_WORD
setopt MENU_COMPLETE

# ---------------------------------------------
#              DIRECTORY CONFIG
# ---------------------------------------------
setopt AUTO_CD
setopt AUTO_PUSHD
setopt PUSHD_IGNORE_DUPS
setopt PUSHD_SILENT

# ---------------------------------------------
#              OTHER OPTIONS
# ---------------------------------------------
setopt INTERACTIVE_COMMENTS
setopt MULTIOS
setopt NO_BEEP
setopt PROMPT_SUBST

# ---------------------------------------------
#              CARAPACE CONFIG
# ---------------------------------------------
export CARAPACE_BRIDGES='zsh,fish,bash,inshellisense'
zstyle ':completion:*' format $'\e[2;37mCompleting %d\e[m'

# ---------------------------------------------
#              MAN PAGES COLORS
# ---------------------------------------------
export MANPAGER="sh -c 'col -bx | bat -l man -p'"
export MANROFFOPT="-c"
