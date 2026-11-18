#!/usr/bin/env bash
# E2E del instalador PACMAN. Pensado para Dockerfile.test, pero se puede
# ejecutar en local con HOME apuntando a un directorio temporal:
#   HOME=$(mktemp -d) PACMAN_DOTFILES=$PWD e2e/e2e_test.sh
set -uo pipefail

DOTFILES="${PACMAN_DOTFILES:-$HOME/dotfiles}"
BIN="$DOTFILES/bin/installer"
export PACMAN_DOTFILES="$DOTFILES"

PASSED=0 FAILED=0
pass() { printf '\033[32m[PASS]\033[0m %s\n' "$1"; PASSED=$((PASSED + 1)); }
fail() { printf '\033[31m[FAIL]\033[0m %s\n' "$1"; FAILED=$((FAILED + 1)); }
check() { local desc="$1"; shift; if "$@" >/dev/null 2>&1; then pass "$desc"; else fail "$desc"; fi; }

check "--version muestra PACMAN"     bash -c "'$BIN' --version | grep -q PACMAN"
check "--help funciona"               bash -c "'$BIN' --help 2>&1 | grep -q -- --install"
check "--list muestra zsh"            bash -c "'$BIN' --list | grep -q zsh"
check "componente desconocido falla"  bash -c "! '$BIN' --install nope"

# dry-run no debe crear nada
"$BIN" --dry-run --install zsh >/dev/null 2>&1
check "dry-run no crea ~/.zshrc"      test ! -e "$HOME/.zshrc"

# un archivo previo debe respaldarse
echo "previo" > "$HOME/.zshrc"
check "instala zsh y tmux (sin paquetes)" "$BIN" --install zsh,tmux --no-packages
check ".zshrc es symlink al repo"     bash -c "[ \"\$(readlink '$HOME/.zshrc')\" = '$DOTFILES/shell/.zshrc' ]"
check "tmux.conf enlazado"            test -L "$HOME/.config/tmux/tmux.conf"
check "aliases.zsh enlazado"          test -L "$HOME/.dotfiles/zsh/aliases.zsh"
check ".zshrc previo respaldado"      bash -c "cat '$HOME'/.dotfiles-backup/links/*/.zshrc | grep -q previo"
check "sin .bak junto al original"    bash -c "! ls '$HOME'/.zshrc.bak.* 2>/dev/null"
check "backup completo creado"        test -d "$HOME/.dotfiles-backup"

# segunda ejecución idempotente: no crea más backups
before=$(find "$HOME/.dotfiles-backup/links" -name .zshrc | wc -l)
"$BIN" --install zsh --no-packages --no-backup >/dev/null 2>&1
check "reinstalar es idempotente"     test "$(find "$HOME/.dotfiles-backup/links" -name .zshrc | wc -l)" -eq "$before"
check "--status reporta enlazado"     bash -c "'$BIN' --status | grep -q 'enlazado'"

echo
echo "Resultados: $PASSED ok, $FAILED fallos"
exit $((FAILED > 0))
