#!/usr/bin/env zsh
# =============================================
#          CUSTOM FUNCTIONS
# =============================================

# ---------------------------------------------
#          WINDOW MANAGER AUTOSTART
# ---------------------------------------------
start_if_needed() {
    WM_VAR="/$ZELLIJ"
    WM_CMD="zellij"
    
    if [[ $- == *i* ]] && [[ -z "${WM_VAR#/}" ]] && [[ -t 1 ]]; then
        exec $WM_CMD
    fi
}

# ---------------------------------------------
#          DIRECTORY NAVIGATION
# ---------------------------------------------
# Create directory and cd into it
mkcd() {
    mkdir -p "$1" && cd "$1"
}

# Go up n directories
up() {
    local d=""
    local limit="$1"
    
    if [ -z "$limit" ] || [ "$limit" -le 0 ]; then
        limit=1
    fi
    
    for ((i=1;i<=limit;i++)); do
        d="../$d"
    done
    
    cd "$d" || return
}

# ---------------------------------------------
#          FILE OPERATIONS
# ---------------------------------------------
# Extract any compressed file
extract() {
    if [ -f "$1" ]; then
        case "$1" in
            *.tar.bz2)   tar xjf "$1"     ;;
            *.tar.gz)    tar xzf "$1"     ;;
            *.bz2)       bunzip2 "$1"     ;;
            *.rar)       unrar x "$1"     ;;
            *.gz)        gunzip "$1"      ;;
            *.tar)       tar xf "$1"      ;;
            *.tbz2)      tar xjf "$1"     ;;
            *.tgz)       tar xzf "$1"     ;;
            *.zip)       unzip "$1"       ;;
            *.Z)         uncompress "$1"  ;;
            *.7z)        7z x "$1"        ;;
            *.deb)       ar x "$1"        ;;
            *.tar.xz)    tar xf "$1"      ;;
            *.tar.zst)   unzstd "$1"      ;;
            *)           echo "'$1' cannot be extracted via extract()" ;;
        esac
    else
        echo "'$1' is not a valid file"
    fi
}

# Get file size
sizeof() {
    du -sh "$1"
}

# ---------------------------------------------
#          GIT FUNCTIONS
# ---------------------------------------------
# Git commit with message
gcm() {
    git add .
    git commit -m "$*"
}

# Git add, commit and push
gacp() {
    git add .
    git commit -m "$*"
    git push
}

# Create new branch and switch to it
gnb() {
    git checkout -b "$1"
}

# ---------------------------------------------
#          DEVELOPMENT FUNCTIONS
# ---------------------------------------------
# Start a simple HTTP server
serve() {
    local port="${1:-8000}"
    python3 -m http.server "$port"
}

# Kill process running on specific port
killport() {
    if [ -z "$1" ]; then
        echo "Usage: killport <port_number>"
        return 1
    fi
    
    local pid=$(lsof -ti:"$1")
    if [ -z "$pid" ]; then
        echo "No process found on port $1"
        return 1
    fi
    
    echo "Killing process $pid on port $1"
    kill -9 "$pid"
}

# Find large files
findlarge() {
    local size="${1:-100M}"
    find . -type f -size +"$size" -exec ls -lh {} \; | awk '{ print $9 ": " $5 }'
}

# ---------------------------------------------
#          SYSTEM FUNCTIONS
# ---------------------------------------------
# Update all package managers
update-all() {
    echo "🔄 Updating Homebrew..."
    brew update && brew upgrade && brew cleanup
    
    if command -v npm &> /dev/null; then
        echo "🔄 Updating NPM global packages..."
        npm update -g
    fi
    
    if command -v cargo &> /dev/null; then
        echo "🔄 Updating Cargo packages..."
        cargo install-update -a
    fi
    
    echo "✅ All updates complete!"
}

# Show system information
sysinfo() {
    echo "🖥️  System Information"
    echo "===================="
    echo "OS: $(uname -s)"
    echo "Kernel: $(uname -r)"
    echo "Hostname: $(hostname)"
    echo "Shell: $SHELL"
    echo "Terminal: $TERM"
}

# ---------------------------------------------
#          NETWORK FUNCTIONS
# ---------------------------------------------
# Get external IP
myip() {
    echo "External IP: $(curl -s ifconfig.me)"
    if [[ "$(uname)" == "Darwin" ]]; then
        echo "Local IP: $(ipconfig getifaddr en0 2>/dev/null)"
    else
        echo "Local IP: $(hostname -I | awk '{print $1}')"
    fi
}

# ---------------------------------------------
#          PROJECT FUNCTIONS
# ---------------------------------------------
# Quick jump to projects
proj() {
    local project_dir="${PROJECT_PATHS:-$HOME/work}"
    if [ -z "$1" ]; then
        cd "$project_dir"
    else
        cd "$project_dir/$1" || echo "Project '$1' not found"
    fi
}

# List all projects
projects() {
    local project_dir="${PROJECT_PATHS:-$HOME/work}"
    echo "📁 Projects in $project_dir:"
    exa -1 "$project_dir" 2>/dev/null || ls -1 "$project_dir"
}

# ---------------------------------------------
#          SEARCH FUNCTIONS
# ---------------------------------------------
# Find and edit file with fzf
fe() {
    local file=$(fzf --preview="bat --theme=gruvbox-dark --color=always {}")
    if [ -n "$file" ]; then
        $EDITOR "$file"
    fi
}

# ---------------------------------------------
#          PRODUCTIVITY FUNCTIONS
# ---------------------------------------------
# Quick note taking
note() {
    local note_dir="$HOME/notes"
    mkdir -p "$note_dir"
    local note_file="$note_dir/$(date +%Y-%m-%d).md"
    
    if [ -z "$1" ]; then
        $EDITOR "$note_file"
    else
        echo "$(date +%H:%M:%S) - $*" >> "$note_file"
        echo "✓ Note added to $(basename $note_file)"
    fi
}

# Backup file with timestamp
backup() {
    if [ -z "$1" ]; then
        echo "Usage: backup <file>"
        return 1
    fi
    
    local timestamp=$(date +%Y%m%d_%H%M%S)
    cp "$1" "${1}.backup_${timestamp}"
    echo "✓ Backup created: ${1}.backup_${timestamp}"
}
