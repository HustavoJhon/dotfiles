package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hustavojhon/dotfiles/internal/installer"
	"github.com/hustavojhon/dotfiles/internal/platform"
)

type MenuModel struct {
	Components []installer.Component
	Selected   map[string]bool
	Cursor     int
	Done       bool
	Platform   platform.PlatformInfo
}

func NewMenuModel(components []installer.Component) MenuModel {
	selected := make(map[string]bool)
	for _, c := range components {
		selected[c.ID] = true
	}
	return MenuModel{
		Components: components,
		Selected:   selected,
		Platform:   platform.Detect(),
	}
}

func (m MenuModel) Init() tea.Cmd { return nil }

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Components)-1 {
				m.Cursor++
			}
		case " ":
			comp := m.Components[m.Cursor]
			if platform.IsComponentCompatible(comp.ID, m.Platform) {
				id := comp.ID
				m.Selected[id] = !m.Selected[id]
			}
		case "enter":
			m.Done = true
			return m, nil
		case "q", "ctrl+c":
			m.Done = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MenuModel) SelectedComponents() []installer.Component {
	var components []installer.Component
	for _, c := range m.Components {
		if m.Selected[c.ID] {
			components = append(components, c)
		}
	}
	return components
}

func (m MenuModel) View() string {
	help := HeaderStyle.Render("↑/↓ navigate  ·  Space toggle  ·  Enter confirm  ·  q quit")
	lines := []string{help, ""}

	for i, comp := range m.Components {
		cursor := "  "
		if i == m.Cursor {
			cursor = "❯ "
		}

		compatible := platform.IsComponentCompatible(comp.ID, m.Platform)

		if !compatible {
			reason := platform.IncompatibleReason(comp.ID, m.Platform)
			line := cursor + "[✗] " + comp.Name + "  " + DimStyle.Render("(" + reason + ")")
			if i == m.Cursor {
				lines = append(lines, DimStyle.Render(line))
				lines = append(lines, DimStyle.Render("   "+comp.Description))
			} else {
				lines = append(lines, DimStyle.Render("[✗] "+comp.Name+"  ("+reason+")"))
			}
			continue
		}

		checked := "[ ]"
		if m.Selected[comp.ID] {
			checked = "[✓]"
		}

		line := cursor + checked + " " + comp.Name

		if i == m.Cursor {
			lines = append(lines, SelectedStyle.Render(line))
			lines = append(lines, DimStyle.Render("   "+comp.Description))
		} else {
			lines = append(lines, ItemStyle.Render(checked+" "+comp.Name))
		}
	}

	content := lipgloss.JoinVertical(lipgloss.Center, lines...)
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Render(content)
}
