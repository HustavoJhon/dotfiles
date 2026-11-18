package tui

import "github.com/charmbracelet/lipgloss"

// Pac-Man arcade palette on a Gruvbox base.
var (
	Light  = lipgloss.Color("#ebdbb2")
	Red    = lipgloss.Color("#fb4934")
	Green  = lipgloss.Color("#b8bb26")
	Yellow = lipgloss.Color("#fabd2f")
	Blue   = lipgloss.Color("#83a598")
	Aqua   = lipgloss.Color("#8ec07c")
	Orange = lipgloss.Color("#fe8019")
	Gray   = lipgloss.Color("#665c54")
	DimFg  = lipgloss.Color("#a89984")

	TitleStyle    = lipgloss.NewStyle().Foreground(Yellow).Bold(true)
	SubtitleStyle = lipgloss.NewStyle().Foreground(Blue)
	GhostStyle    = lipgloss.NewStyle().Foreground(Red)
	BannerStyle   = lipgloss.NewStyle().Foreground(Yellow).Bold(true)
	SelectedStyle = lipgloss.NewStyle().Foreground(Yellow).Bold(true)
	ItemStyle     = lipgloss.NewStyle().Foreground(Light)
	DimStyle      = lipgloss.NewStyle().Foreground(DimFg)
	HelpStyle     = lipgloss.NewStyle().Foreground(Gray)
	ErrorStyle    = lipgloss.NewStyle().Foreground(Red).Bold(true)
	WarnStyle     = lipgloss.NewStyle().Foreground(Orange)
	SuccessStyle  = lipgloss.NewStyle().Foreground(Green).Bold(true)
	OmarchyStyle  = lipgloss.NewStyle().Foreground(Aqua).Bold(true)
	CodeStyle     = lipgloss.NewStyle().Foreground(Aqua)

	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Yellow).
			Padding(1, 2)
)
