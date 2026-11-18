#!/usr/bin/env bash
# PACMAN - Bootstrap: delega en install.sh de la raíz del repo.
set -euo pipefail
exec "$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)/install.sh" "$@"
