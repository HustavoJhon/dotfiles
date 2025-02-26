#!/usr/bin/env zsh
# =============================================
#              ALIASES CONFIGURATION
# =============================================

# ---------------------------------------------
#                 EXA ALIASES
# ---------------------------------------------
# Main ls replacement
alias ls='exa --icons --group-directories-first'
alias la='exa -a --group-directories-first --icons'
alias ll='exa -l --header --group-directories-first --icons --no-user --no-time'
alias lt='exa -aT --group-directories-first --git-ignore --icons'
alias l.='exa -a | grep "^\."'
alias li='exa --git-ignore --group-directories-first --icons'
alias lg='exa -al --header --git --group-directories-first --icons --no-user --no-time'
alias lp='exa -al --header --octal-permissions --group-directories-first --icons --no-user --no-time'
alias tree='exa --icons --tree -L 1 -I node_modules'

# Tree with different depths
alias tree2='exa --icons --tree -L 2 -I node_modules'
alias tree3='exa --icons --tree -L 3 -I node_modules'

# ---------------------------------------------
#                 FZF ALIASES
# ---------------------------------------------
alias fzfbat='fzf --preview="bat --theme=gruvbox-dark --color=always {}"'
alias fzfnvim='nvim $(fzf --preview="bat --theme=gruvbox-dark --color=always {}")'

# ---------------------------------------------
#              NAVIGATION ALIASES
# ---------------------------------------------
alias ..='cd ..'
alias ...='cd ../..'
alias ....='cd ../../..'
alias ~='cd ~'

# ---------------------------------------------
#                BASIC ALIASES
# ---------------------------------------------
alias cls='clear'
alias c='clear'
alias grep='grep --color'
alias egrep='egrep --color=auto'
alias fgrep='fgrep --color=auto'

# ---------------------------------------------
#              PYTHON ALIASES
# ---------------------------------------------
alias python='python3'
alias py='python3'
alias pip='pip3'

# ---------------------------------------------
#                GIT ALIASES
# ---------------------------------------------
alias g='git'
alias gs='git status'
alias ga='git add'
alias gaa='git add .'
alias gc='git commit -m'
alias gp='git push'
alias gpl='git pull'
alias gd='git diff'
alias gb='git branch'
alias gco='git checkout'
alias gcb='git checkout -b'
alias glog='git log --oneline --decorate --graph'
alias gitlog="git log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --all"

# ---------------------------------------------
#              UTILITY ALIASES
# ---------------------------------------------
alias diff='icdiff'
alias copy='xclip -selection clipboard <'
alias paste='xclip -selection clipboard -o'
alias reload='source ~/.zshrc'
alias edit-zsh='nvim ~/.zshrc'
alias edit-aliases='nvim ~/.dotfiles/zsh/aliases.zsh'

# ---------------------------------------------
#              SYSTEM ALIASES
# ---------------------------------------------
alias ports='netstat -tulanp'
alias myip='curl ifconfig.me'
alias localip='ipconfig getifaddr en0'

# ---------------------------------------------
#           FILE OPERATIONS ALIASES
# ---------------------------------------------
alias rm='rm -i'
alias cp='cp -i'
alias mv='mv -i'
alias mkdir='mkdir -pv'

# ---------------------------------------------
#              DEVELOPMENT ALIASES
# ---------------------------------------------
# Node/NPM
alias ni='npm install'
alias nid='npm install --save-dev'
alias nig='npm install -g'
alias nu='npm uninstall'
alias nup='npm update'
alias nri='rm -rf node_modules package-lock.json && npm install'
alias ns='npm start'
alias nt='npm test'
alias nr='npm run'

# Yarn
alias y='yarn'
alias ya='yarn add'
alias yad='yarn add --dev'
alias yr='yarn remove'
alias yi='yarn install'
alias ys='yarn start'
alias yt='yarn test'

# PNPM
alias pn='pnpm'
alias pni='pnpm install'
alias pna='pnpm add'
alias pnad='pnpm add -D'
alias pnr='pnpm remove'
alias pns='pnpm start'
alias pnt='pnpm test'

# Docker
alias d='docker'
alias dc='docker-compose'
alias dps='docker ps'
alias dpa='docker ps -a'
alias di='docker images'
alias dex='docker exec -it'
alias dlog='docker logs -f'

# Bun
alias b='bun'
alias bi='bun install'
alias ba='bun add'
alias bad='bun add -d'
alias br='bun remove'
alias bs='bun start'
alias bt='bun test'

# ---------------------------------------------
#              QUICK EDIT ALIASES
# ---------------------------------------------
alias zshconfig='nvim ~/.zshrc'
alias ohmyzsh='nvim ~/.oh-my-zsh'
alias hosts='sudo nvim /etc/hosts'

# ---------------------------------------------
#           FORCE EXA (IMPORTANTE)
# ---------------------------------------------
# Esto asegura que exa siempre sobrescriba ls
unalias ls 2>/dev/null
alias ls='exa --icons --group-directories-first'
