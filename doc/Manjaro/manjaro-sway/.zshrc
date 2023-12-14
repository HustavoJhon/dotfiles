# Use powerline
USE_POWERLINE="true"
# Has weird character width
# Example:
#  is not a diamond
HAS_WIDECHARS="false"
# Source manjaro-zsh-configuration
if [[ -e /usr/share/zsh/manjaro-zsh-config ]]; then
  source /usr/share/zsh/manjaro-zsh-config
fi
# Use manjaro zsh prompt
if [[ -e /usr/share/zsh/manjaro-zsh-prompt ]]; then
  source /usr/share/zsh/manjaro-zsh-prompt
fi

# ALIASES

# LSD
# alias ls='lsd'
# alias l='ls -l'
# alias la='ls -a'
alias lla='ls -la'
# alias lt='ls --tree'

# EXA
alias ls='exa -D --group-directories-first --icons' # Solo Folder
alias la='exa -a --group-directories-first --icons' # Nuevo ls con opciones preferidas
alias ll='exa -l --header --group-directories-first --icons --no-user --no-time' # Formato Largo
alias lt='exa -aT --group-directories-first --icons' #Mostra dentro dolder
alias l.='exa -a | egrep "^\."' # Solo ocultos 
alias li='exa --git-ignore --group-directories-first --icons' # Ignora archivos de .gitignore
alias lg='exa -al --header --git --group-directories-first --icons --no-user --no-time' # Formato largp mas git data
alias lp='exa -al --header --octal-permissions --group-directories-first --icons --no-user --no-time' # formato largo mas permisos
