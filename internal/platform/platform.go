package platform

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type PlatformType string

const (
	PlatformOmarchy PlatformType = "omarchy"
	PlatformLinux   PlatformType = "linux"
	PlatformMacOS   PlatformType = "macos"
	PlatformWSL     PlatformType = "wsl"
)

type PlatformInfo struct {
	Type        PlatformType
	DisplayName string
	HasGUI      bool
	IsWayland   bool
}

func Detect() PlatformInfo {
	if runtime.GOOS == "darwin" {
		return PlatformInfo{
			Type:        PlatformMacOS,
			DisplayName: "macOS",
			HasGUI:      true,
			IsWayland:   false,
		}
	}

	if isWSL() {
		return PlatformInfo{
			Type:        PlatformWSL,
			DisplayName: "WSL (Windows)",
			HasGUI:      false,
			IsWayland:   false,
		}
	}

	if isOmarchy() {
		return PlatformInfo{
			Type:        PlatformOmarchy,
			DisplayName: "Omarchy (Hyprland)",
			HasGUI:      true,
			IsWayland:   true,
		}
	}

	hasGUI := os.Getenv("DISPLAY") != "" || os.Getenv("WAYLAND_DISPLAY") != ""
	isWayland := os.Getenv("WAYLAND_DISPLAY") != ""

	return PlatformInfo{
		Type:        PlatformLinux,
		DisplayName: "Linux",
		HasGUI:      hasGUI,
		IsWayland:   isWayland,
	}
}

func isWSL() bool {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	lower := strings.ToLower(string(data))
	return strings.Contains(lower, "microsoft") || strings.Contains(lower, "wsl")
}

func isOmarchy() bool {
	desktop := os.Getenv("XDG_CURRENT_DESKTOP")
	if strings.Contains(strings.ToLower(desktop), "hyprland") {
		return true
	}
	if _, err := exec.LookPath("hyprctl"); err == nil {
		if os.Getenv("HYPRLAND_INSTANCE_SIGNATURE") != "" {
			return true
		}
	}
	return false
}

var (
	omarchyComponents = []string{"zsh", "kitty", "ghostty", "fastfetch", "starship", "hyprland", "zellij", "tmux", "nvim"}
	linuxComponents   = []string{"zsh", "kitty", "fastfetch", "starship", "tmux", "nvim"}
	macosComponents   = []string{"zsh", "kitty", "fastfetch", "starship", "tmux", "nvim"}
	wslComponents     = []string{"zsh", "fastfetch", "starship", "tmux", "nvim"}
)

func DefaultComponents(p PlatformInfo) []string {
	switch p.Type {
	case PlatformOmarchy:
		return omarchyComponents
	case PlatformMacOS:
		return macosComponents
	case PlatformWSL:
		return wslComponents
	default:
		return linuxComponents
	}
}

func IsComponentCompatible(componentID string, p PlatformInfo) bool {
	switch componentID {
	case "hyprland":
		return p.Type == PlatformOmarchy || (p.Type == PlatformLinux && p.IsWayland)
	case "kitty", "wezterm", "ghostty":
		return p.HasGUI || p.Type == PlatformMacOS
	default:
		return true
	}
}

func IncompatibleReason(componentID string, p PlatformInfo) string {
	switch componentID {
	case "hyprland":
		if p.Type == PlatformWSL {
			return "Not available on WSL"
		}
		if p.Type == PlatformMacOS {
			return "Not available on macOS"
		}
		if !p.IsWayland {
			return "Requires Wayland"
		}
	case "kitty", "wezterm", "ghostty":
		if p.Type == PlatformWSL {
			return "GUI not available on WSL"
		}
	}
	return ""
}
