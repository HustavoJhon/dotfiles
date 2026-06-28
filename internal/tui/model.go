package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/hustavojhon/dotfiles/internal/installer"
)

type Screen int

const (
	ScreenWelcome Screen = iota
	ScreenDetect
	ScreenMenu
	ScreenBackup
	ScreenInstall
	ScreenDone
)

type installMsg struct {
	Component installer.Component
}

type installErrMsg struct {
	Component installer.Component
	Err       error
}

type Model struct {
	Screen       Screen
	Spinner      spinner.Model
	Progress     progress.Model
	Menu         MenuModel
	Installer    *installer.Installer
	Selected     []installer.Component
	DistroName   string
	InstallStep  int
	InstallTotal int
	InstallLog   []string
	Width        int
	Height       int
	Ready        bool
	Err          error
}

func NewModel(inst *installer.Installer) Model {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(Aqua)

	return Model{
		Screen:    ScreenWelcome,
		Spinner:   s,
		Progress:  progress.New(progress.WithDefaultGradient()),
		Installer: inst,
		Menu:      NewMenuModel(installer.Components),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.Spinner.Tick, tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Ready = true
		m.Progress.Width = clamp(msg.Width-20, 20, 100)
		return m, nil

	case tea.KeyMsg:
		switch m.Screen {
		case ScreenWelcome:
			if msg.String() == "enter" || msg.String() == " " {
				m.Screen = ScreenDetect
				info, _ := m.Installer.Distro.(interface{ Name() string })
				m.DistroName = info.Name()
				m.Screen = ScreenMenu
				return m, nil
			}
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}

		case ScreenMenu:
			updatedMenu, cmd := m.Menu.Update(msg)
			m.Menu = updatedMenu.(MenuModel)
			if m.Menu.Done {
				m.Selected = m.Menu.SelectedComponents()
				if len(m.Selected) == 0 {
					m.Err = fmt.Errorf("no components selected")
					return m, tea.Quit
				}
				m.Screen = ScreenBackup
				return m, nil
			}
			return m, cmd

		case ScreenBackup:
			switch msg.String() {
			case "y", "Y":
				m.Installer.RunBackup()
				m.Screen = ScreenInstall
				m.InstallTotal = len(m.Selected)
				return m, m.startInstall()
			case "n", "N":
				m.Screen = ScreenInstall
				m.InstallTotal = len(m.Selected)
				return m, m.startInstall()
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case ScreenInstall:
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}

		case ScreenDone:
			if msg.String() == "enter" || msg.String() == "q" || msg.String() == " " {
				return m, tea.Quit
			}
		}

	case installMsg:
		m.InstallStep++
		m.InstallLog = append(m.InstallLog, fmt.Sprintf("OK %s installed", msg.Component.Name))
		if m.InstallStep >= m.InstallTotal {
			m.Screen = ScreenDone
			return m, nil
		}
		return m, m.startInstall()

	case installErrMsg:
		m.InstallLog = append(m.InstallLog, fmt.Sprintf("ERR %s: %v", msg.Component.Name, msg.Err))
		m.InstallStep++
		if m.InstallStep >= m.InstallTotal {
			m.Screen = ScreenDone
			return m, nil
		}
		return m, m.startInstall()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd

	case progress.FrameMsg:
		pm, cmd := m.Progress.Update(msg)
		m.Progress = pm.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m Model) startInstall() tea.Cmd {
	return func() tea.Msg {
		if m.InstallStep >= len(m.Selected) {
			return nil
		}
		comp := m.Selected[m.InstallStep]
		if err := m.Installer.InstallComponent(comp); err != nil {
			return installErrMsg{Component: comp, Err: err}
		}
		return installMsg{Component: comp}
	}
}

func (m Model) View() string {
	if !m.Ready {
		return "\n  Initializing..."
	}

	var content string
	switch m.Screen {
	case ScreenWelcome:
		content = m.welcomeView()
	case ScreenDetect:
		content = m.detectView()
	case ScreenMenu:
		content = m.menuView()
	case ScreenBackup:
		content = m.backupView()
	case ScreenInstall:
		content = m.installView()
	case ScreenDone:
		content = m.doneView()
	}

	box := BorderStyle.
		Width(m.Width - 4).
		Height(m.Height - 2).
		Render(content)

	return lipgloss.Place(
		m.Width, m.Height,
		lipgloss.Center, lipgloss.Center,
		box,
		lipgloss.WithWhitespaceBackground(DarkBg),
	)
}

func (m Model) welcomeView() string {
	s := BannerStyle.Render(logoAscii()) + "\n\n"
	s += GhostStyle.Render(pacmanGhostArt()) + "\n\n"
	s += TitleStyle.Align(lipgloss.Center).Render("PACMAN INSTALLER v1.0") + "\n\n"
	s += DimStyle.Align(lipgloss.Center).Render("Press ENTER to start  ·  Press q to quit")
	return s
}

func (m Model) detectView() string {
	s := TitleStyle.Align(lipgloss.Center).Render("Detecting System") + "\n\n"
	s += lipgloss.NewStyle().Align(lipgloss.Center).Render(
		m.Spinner.View() + " Detecting Linux distribution...",
	)
	return s
}

func (m Model) menuView() string {
	s := TitleStyle.Align(lipgloss.Center).Render("Select Components") + "\n"
	s += SubtitleStyle.Align(lipgloss.Center).Render("Distro: "+m.DistroName) + "\n\n"
	s += m.Menu.View()
	return s
}

func (m Model) backupView() string {
	s := TitleStyle.Align(lipgloss.Center).Render("Backup") + "\n\n"
	s += lipgloss.NewStyle().Align(lipgloss.Center).Render(
		"Create backup of existing configs before installing?\n\n",
	)
	s += lipgloss.NewStyle().Align(lipgloss.Center).Render(
		SelectedStyle.Render("[Y] Yes") + "  " +
			DimStyle.Render("[N] No  [Q] Quit"),
	)
	return s
}

func (m Model) installView() string {
	s := TitleStyle.Align(lipgloss.Center).Render("Installing") + "\n\n"

	bar := m.Progress.ViewAs(float64(m.InstallStep) / float64(max(m.InstallTotal, 1)))
	s += lipgloss.NewStyle().Align(lipgloss.Center).Render(bar) + "\n\n"

	for _, log := range m.InstallLog {
		if strings.HasPrefix(log, "OK") {
			s += SuccessStyle.Render(log) + "\n"
		} else {
			s += ErrorStyle.Render(log) + "\n"
		}
	}

	if m.InstallStep < m.InstallTotal {
		comp := m.Selected[m.InstallStep]
		s += "\n" + lipgloss.NewStyle().Align(lipgloss.Center).Render(
			m.Spinner.View()+" Installing "+comp.Name+"...",
		)
	}

	return s
}

func (m Model) doneView() string {
	s := TitleStyle.Align(lipgloss.Center).Render("Installation Complete") + "\n\n"

	for _, log := range m.InstallLog {
		if strings.HasPrefix(log, "OK") {
			s += SuccessStyle.Render(log) + "\n"
		} else {
			s += ErrorStyle.Render(log) + "\n"
		}
	}

	if m.Err != nil {
		s += "\n" + ErrorStyle.Render("Error: "+m.Err.Error()) + "\n"
	}

	s += "\n" + DimStyle.Align(lipgloss.Center).Render("Press ENTER to exit")
	return s
}

func logoAscii() string {
	return `██████╗  █████╗  ██████╗███╗   ███╗ █████╗ ███╗   ██╗
██╔══██╗██╔══██╗██╔════╝████╗ ████║██╔══██╗████╗  ██║
██████╔╝███████║██║     ██╔████╔██║███████║██╔██╗ ██║
██╔═══╝ ██╔══██║██║     ██║╚██╔╝██║██╔══██║██║╚██╗██║
██║     ██║  ██║╚██████╗██║ ╚═╝ ██║██║  ██║██║ ╚████║
╚═╝     ╚═╝  ╚═╝ ╚═════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝`
}

func pacmanGhostArt() string {
	return `         .--.
        /    \
       / /-_-\ \
      / | (") | \
     /  \  \_/  / \
    /   |       |  \
   /    |       |   \
  /     |       |    \
 /______|_______|_____\
  |     |       |     |
  |     |       |     |
  |_____|_______|_____|`
}
