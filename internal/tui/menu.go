package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/hustavojhon/dotfiles/internal/installer"
)

type MenuModel struct {
	Components []installer.Component
	Selected   map[string]bool
	Cursor     int
	Done       bool
	Quit       bool
}

func NewMenuModel(components []installer.Component) MenuModel {
	selected := make(map[string]bool)
	for _, c := range components {
		selected[c.ID] = c.Default
	}
	return MenuModel{Components: components, Selected: selected}
}

func (m MenuModel) Update(msg tea.KeyMsg) MenuModel {
	switch msg.String() {
	case "up", "k":
		if m.Cursor > 0 {
			m.Cursor--
		}
	case "down", "j":
		if m.Cursor < len(m.Components)-1 {
			m.Cursor++
		}
	case " ", "x":
		id := m.Components[m.Cursor].ID
		m.Selected[id] = !m.Selected[id]
	case "a":
		for _, c := range m.Components {
			m.Selected[c.ID] = true
		}
	case "n":
		for _, c := range m.Components {
			m.Selected[c.ID] = false
		}
	case "enter":
		if len(m.SelectedComponents()) > 0 {
			m.Done = true
		}
	case "q", "esc", "ctrl+c":
		m.Quit = true
	}
	return m
}

func (m MenuModel) SelectedComponents() []installer.Component {
	var out []installer.Component
	for _, c := range m.Components {
		if m.Selected[c.ID] {
			out = append(out, c)
		}
	}
	return out
}

func (m MenuModel) View() string {
	var b strings.Builder
	for i, c := range m.Components {
		check := DimStyle.Render("○")
		if m.Selected[c.ID] {
			check = SuccessStyle.Render("●")
		}
		name := ItemStyle.Render(c.Name)
		cursor := "  "
		if i == m.Cursor {
			cursor = SelectedStyle.Render("❯ ")
			name = SelectedStyle.Render(c.Name)
		}
		b.WriteString(cursor + check + " " + name + "\n")
		if i == m.Cursor {
			b.WriteString("      " + DimStyle.Render(c.Description) + "\n")
		}
	}
	b.WriteString("\n" + HelpStyle.Render("↑/↓ mover · espacio marcar · a todo · n nada · enter continuar · q salir"))
	return b.String()
}
