# Set name of the theme to load --- if set to "random", it will
# load a random theme each time oh-my-zsh is loaded, in which case,
# to know which specific one was loaded, run: echo $RANDOM_THEME
# See https://github.com/robbyrussell/oh-my-zsh/wiki/Themes
ZSH_THEME=""

DEFAULT_USER=`whoami`

# configure tmux plugin
if [ ! -z "$SSH_TTY" ]; then
  ZSH_TMUX_AUTOSTART=true
  ZSH_TMUX_AUTOCONNECT=true
  ZSH_TMUX_AUTOQUIT=false
fi

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
. ~/.bash_aliases
