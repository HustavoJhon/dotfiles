#!/bin/bash
# E2E tests for PACMAN installer

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'
PASSED=0
FAILED=0

log_test() { printf "${YELLOW}[TEST]${NC} %s\n" "$1"; }
log_pass() { printf "${GREEN}[PASS]${NC} %s\n" "$1"; PASSED=$((PASSED + 1)); }
log_fail() { printf "${RED}[FAIL]${NC} %s\n" "$1"; FAILED=$((FAILED + 1)); }

# --- Basic Tests ---

test_help() {
    log_test "installer --help exits cleanly"
    if /usr/local/bin/pacman-installer --help | head -5; then
        log_pass "installer --help works"
    else
        log_fail "installer --help failed"
    fi
}

test_version() {
    log_test "installer --version shows version"
    output=$(/usr/local/bin/pacman-installer --version 2>&1)
    if echo "$output" | grep -q "PACMAN"; then
        log_pass "installer --version shows PACMAN"
    else
        log_fail "installer --version output: $output"
    fi
}

test_backup() {
    log_test "installer --backup-only runs without error"
    output=$(/usr/local/bin/pacman-installer --backup-only 2>&1)
    if echo "$output" | grep -q "Backup"; then
        log_pass "backup completed"
    else
        log_fail "backup output: $output"
    fi
}

# --- Config Files Check ---

test_shell_configs() {
    log_test "shell configs exist"
    if [ -f "/app/shell/.zshrc" ]; then
        log_pass "shell/.zshrc exists"
    else
        log_fail "shell/.zshrc missing"
    fi
}

test_term_configs() {
    log_test "terminal configs exist"
    if [ -f "/app/term/kitty/kitty.conf" ]; then
        log_pass "kitty config exists"
    else
        log_fail "kitty config missing"
    fi
    if [ -f "/app/term/wezterm/.wezterm.lua" ]; then
        log_pass "wezterm config exists"
    else
        log_fail "wezterm config missing"
    fi
}

test_wm_configs() {
    log_test "wm configs exist"
    if [ -f "/app/wm/hypr/hyprland.conf" ]; then
        log_pass "hyprland config exists"
    fi
    if [ -f "/app/wm/zellij/config.kdl" ]; then
        log_pass "zellij config exists"
    fi
    if [ -f "/app/wm/tmux/.tmux.conf" ]; then
        log_pass "tmux config exists"
    fi
}

# --- Run Tests ---

echo "=================================="
echo "  PACMAN E2E Tests"
echo "=================================="
echo ""

test_help
test_version
test_backup
test_shell_configs
test_term_configs
test_wm_configs

echo ""
echo "=================================="
echo "  Results: $PASSED passed, $FAILED failed"
echo "=================================="
exit $FAILED
