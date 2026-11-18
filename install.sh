#!/usr/bin/env bash
# PACMAN dotfiles - instalador rápido
#   curl -fsSL https://raw.githubusercontent.com/hustavojhon/dotfiles/main/install.sh | bash
# Variables opcionales:
#   PACMAN_DOTFILES=/ruta   dónde clonar/usar el repo
#   PACMAN_BRANCH=main      rama a clonar
#   PACMAN_ARGS="--dry-run" flags para el instalador

set -euo pipefail

REPO_URL="https://github.com/hustavojhon/dotfiles.git"
BRANCH="${PACMAN_BRANCH:-main}"

Y='\033[1;33m'; C='\033[0;36m'; R='\033[0;31m'; N='\033[0m'
info() { printf "${C}→ %s${N}\n" "$*"; }
die()  { printf "${R}✗ %s${N}\n" "$*" >&2; exit 1; }

printf "${Y}"
cat <<'EOF'
██████╗  █████╗  ██████╗███╗   ███╗ █████╗ ███╗   ██╗
██╔══██╗██╔══██╗██╔════╝████╗ ████║██╔══██╗████╗  ██║
██████╔╝███████║██║     ██╔████╔██║███████║██╔██╗ ██║
██╔═══╝ ██╔══██║██║     ██║╚██╔╝██║██╔══██║██║╚██╗██║
██║     ██║  ██║╚██████╗██║ ╚═╝ ██║██║  ██║██║ ╚████║
╚═╝     ╚═╝  ╚═╝ ╚═════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝
EOF
printf "${N}\n"

# --- localizar o clonar el repo ------------------------------------------
script_dir=""
if [[ -n "${BASH_SOURCE[0]:-}" && -f "${BASH_SOURCE[0]}" ]]; then
    script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
fi

if [[ -n "$script_dir" && -f "$script_dir/cmd/installer/main.go" ]]; then
    DOTFILES="$script_dir"
else
    DOTFILES="${PACMAN_DOTFILES:-$HOME/.local/share/pacman-dotfiles}"
    command -v git >/dev/null || die "git es necesario"
    if [[ -d "$DOTFILES/.git" ]]; then
        info "Actualizando $DOTFILES"
        git -C "$DOTFILES" pull --ff-only
    else
        info "Clonando en $DOTFILES"
        git clone --depth=1 --branch "$BRANCH" "$REPO_URL" "$DOTFILES"
    fi
fi
export PACMAN_DOTFILES="$DOTFILES"

# --- Go -----------------------------------------------------------------
install_go() {
    if command -v pacman >/dev/null; then
        sudo pacman -S --needed --noconfirm go
    elif command -v dnf >/dev/null; then
        sudo dnf install -y golang
    else
        # apt suele traer un Go antiguo: usar el tarball oficial.
        local arch version
        case "$(uname -m)" in
            x86_64) arch=amd64 ;;
            aarch64|arm64) arch=arm64 ;;
            *) die "arquitectura no soportada: $(uname -m)" ;;
        esac
        version="$(curl -fsSL 'https://go.dev/VERSION?m=text' | head -1)"
        info "Descargando $version ($arch)"
        curl -fsSL "https://go.dev/dl/${version}.linux-${arch}.tar.gz" -o /tmp/go.tar.gz
        sudo rm -rf /usr/local/go
        sudo tar -C /usr/local -xzf /tmp/go.tar.gz
        rm -f /tmp/go.tar.gz
        export PATH="/usr/local/go/bin:$PATH"
    fi
}

[[ -x /usr/local/go/bin/go ]] && export PATH="/usr/local/go/bin:$PATH"
if ! command -v go >/dev/null; then
    info "Instalando Go"
    install_go
fi

# go.mod puede pedir una versión más nueva: deja que Go la descargue.
export GOTOOLCHAIN=auto

# --- compilar y ejecutar -------------------------------------------------
info "Compilando instalador"
cd "$DOTFILES"
version="$(git describe --tags --always --dirty 2>/dev/null || echo dev)"
go build -ldflags "-s -w -X main.version=$version" -o bin/installer ./cmd/installer

info "Abriendo instalador"
# Con `curl | bash` stdin es el script: conecta la TUI a la terminal real.
# shellcheck disable=SC2086
if [[ -t 0 ]]; then
    exec ./bin/installer ${PACMAN_ARGS:-}
else
    exec ./bin/installer ${PACMAN_ARGS:-} </dev/tty
fi
