package tui

import "github.com/charmbracelet/lipgloss"

// Gruvbox palette
var (
	DarkBg  = lipgloss.Color("#282828")
	Darker  = lipgloss.Color("#1d2021")
	Light   = lipgloss.Color("#ebdbb2")
	White   = lipgloss.Color("#fbf1c7")
	Red     = lipgloss.Color("#fb4934")
	Green   = lipgloss.Color("#b8bb26")
	Yellow  = lipgloss.Color("#fabd2f")
	Blue    = lipgloss.Color("#83a598")
	Purple  = lipgloss.Color("#d3869b")
	Aqua    = lipgloss.Color("#8ec07c")
	Orange  = lipgloss.Color("#fe8019")
	Gray    = lipgloss.Color("#928374")

	GhostRed = lipgloss.Color("#cc241d")
	DimFg    = lipgloss.Color("#a89984")

	TitleStyle = lipgloss.NewStyle().
			Foreground(Orange).
			Bold(true)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(Blue)

	GhostStyle = lipgloss.NewStyle().
			Foreground(GhostRed).
			Bold(true)

	BannerStyle = lipgloss.NewStyle().
			Foreground(Red).
			Bold(true)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(Yellow).
			Bold(true)

	DimStyle = lipgloss.NewStyle().
			Foreground(DimFg)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(Red).
			Bold(true)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(Green).
			Bold(true)

	ItemStyle = lipgloss.NewStyle().
			Foreground(Light)

	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Gray).
			Padding(0, 1)

	HeaderStyle = lipgloss.NewStyle().
			Foreground(DimFg).
			Align(lipgloss.Center)
)
