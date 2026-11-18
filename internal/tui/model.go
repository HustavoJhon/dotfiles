package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/hustavojhon/dotfiles/internal/assets"
	"github.com/hustavojhon/dotfiles/internal/installer"
)

type Screen int

const (
	ScreenWelcome Screen = iota
	ScreenMenu
	ScreenBackup
	ScreenInstall
	ScreenDone
)

type resultMsg installer.Result

type backupMsg struct {
	path string
	err  error
}

type sudoMsg struct{ err error }

type Model struct {
	Screen    Screen
	Spinner   spinner.Model
	Progress  progress.Model
	Menu      MenuModel
	Installer *installer.Installer
	Version   string

	Selected []installer.Component
	Missing  []string
	Step     int
	Log      []string
	Failures int
	Warnings int

	Width, Height int
}

func NewModel(inst *installer.Installer, version string) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(Yellow)

	return Model{
		Screen:    ScreenWelcome,
		Spinner:   s,
		Progress:  progress.New(progress.WithGradient(string(Yellow), string(Orange)), progress.WithoutPercentage()),
		Installer: inst,
		Version:   version,
		Menu:      NewMenuModel(inst.AvailableComponents()),
	}
}

func (m Model) Init() tea.Cmd { return m.Spinner.Tick }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
		m.Progress.Width = clamp(msg.Width-24, 20, 60)
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		return m.handleKey(msg)

	case backupMsg:
		if msg.err != nil {
			m.Log = append(m.Log, "ERR backup: "+msg.err.Error())
		} else {
			m.Log = append(m.Log, "OK backup en "+msg.path)
		}
		return m.beginInstall()

	case sudoMsg:
		if msg.err != nil {
			m.Installer.SkipPackages = true
			m.Log = append(m.Log, "WARN sudo no disponible: solo se crearán los enlaces")
		}
		return m, m.installNext()

	case resultMsg:
		r := installer.Result(msg)
		switch {
		case r.Err != nil:
			m.Failures++
			m.Log = append(m.Log, fmt.Sprintf("ERR %s: %v", r.Component.Name, r.Err))
		case len(r.Warnings) > 0:
			m.Warnings += len(r.Warnings)
			m.Log = append(m.Log, "WARN "+r.Component.Name+" (enlaces OK)")
			for _, w := range r.Warnings {
				m.Log = append(m.Log, "     "+w)
			}
		default:
			m.Log = append(m.Log, "OK "+r.Component.Name)
		}
		for _, line := range r.Log {
			if strings.HasPrefix(line, "  ✗") || strings.Contains(line, "errores") {
				m.Log = append(m.Log, line)
			}
		}
		m.Step++
		if m.Step >= len(m.Selected) {
			m.Screen = ScreenDone
			return m, nil
		}
		return m, m.installNext()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.String()
	switch m.Screen {
	case ScreenWelcome:
		switch key {
		case "enter", " ":
			m.Screen = ScreenMenu
		case "q", "esc":
			return m, tea.Quit
		}

	case ScreenMenu:
		m.Menu = m.Menu.Update(msg)
		if m.Menu.Quit {
			return m, tea.Quit
		}
		if m.Menu.Done {
			m.Selected = m.Menu.SelectedComponents()
			m.Missing = m.Installer.MissingPackages(m.Selected)
			m.Screen = ScreenBackup
		}

	case ScreenBackup:
		switch key {
		case "y", "Y", "s", "S", "enter":
			m.Screen = ScreenInstall
			inst := m.Installer
			return m, func() tea.Msg {
				p, err := inst.RunBackup(io.Discard)
				return backupMsg{path: p, err: err}
			}
		case "n", "N":
			return m.beginInstall()
		case "q", "esc":
			return m, tea.Quit
		}

	case ScreenDone:
		if key == "enter" || key == "q" || key == "esc" {
			return m, tea.Quit
		}
	}
	return m, nil
}

// beginInstall asks for the sudo password (with the real terminal) only when
// packages are actually missing, then installs components one by one.
func (m Model) beginInstall() (tea.Model, tea.Cmd) {
	m.Screen = ScreenInstall
	if len(m.Missing) > 0 && !m.Installer.DryRun && !m.Installer.SkipPackages {
		return m, tea.ExecProcess(installer.SudoCmd(), func(err error) tea.Msg { return sudoMsg{err} })
	}
	return m, m.installNext()
}

func (m Model) installNext() tea.Cmd {
	if m.Step >= len(m.Selected) {
		return nil
	}
	comp := m.Selected[m.Step]
	inst := m.Installer
	return func() tea.Msg { return resultMsg(inst.InstallComponent(comp)) }
}

func clamp(v, lo, hi int) int {
	return max(lo, min(v, hi))
}

// ---------------------------------------------------------------- views

func (m Model) View() string {
	if m.Width == 0 {
		return ""
	}

	var content string
	switch m.Screen {
	case ScreenWelcome:
		content = m.welcomeView()
	case ScreenMenu:
		content = m.menuView()
	case ScreenBackup:
		content = m.backupView()
	case ScreenInstall:
		content = m.installView()
	case ScreenDone:
		content = m.doneView()
	}

	style := BorderStyle.Width(min(m.Width-4, 84))
	if m.Screen == ScreenWelcome {
		style = style.Align(lipgloss.Center)
	}
	box := style.Render(content)
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, box)
}

func center(s string) string { return lipgloss.PlaceHorizontal(80, lipgloss.Center, s) }

func (m Model) platform() string {
	if m.Installer.Omarchy {
		return OmarchyStyle.Render("● Omarchy " + m.Installer.Info.Version + " detectado")
	}
	return m.Installer.Distro.Name()
}

func (m Model) welcomeView() string {
	lines := []string{
		GhostStyle.Render(assets.Ghost),
		"",
		BannerStyle.Render(assets.Logo),
		"",
		TitleStyle.Render("dotfiles installer " + m.Version),
		SubtitleStyle.Render(m.platform()),
		DimStyle.Render("repo: " + m.Installer.DotfilesDir),
	}
	if m.Installer.DryRun {
		lines = append(lines, WarnStyle.Render("modo --dry-run: no se modificará nada"))
	}
	lines = append(lines, "", HelpStyle.Render("enter empezar · q salir"))
	return lipgloss.JoinVertical(lipgloss.Center, lines...)
}

func (m Model) menuView() string {
	head := lipgloss.JoinVertical(lipgloss.Center,
		TitleStyle.Render("Componentes"),
		SubtitleStyle.Render(m.platform()),
	)
	return center(head) + "\n\n" + m.Menu.View()
}

func (m Model) backupView() string {
	var b strings.Builder
	b.WriteString(TitleStyle.Render("Resumen") + "\n\n")
	for _, c := range m.Selected {
		b.WriteString("  " + SuccessStyle.Render("●") + " " + c.Name + DimStyle.Render(fmt.Sprintf("  (%d enlaces)", len(c.Links))) + "\n")
	}
	b.WriteString("\n")
	if len(m.Missing) > 0 {
		b.WriteString(WarnStyle.Render("Paquetes a instalar (se pedirá sudo): ") + strings.Join(m.Missing, " ") + "\n")
	} else {
		b.WriteString(DimStyle.Render("Todos los paquetes ya están instalados.") + "\n")
	}
	b.WriteString(DimStyle.Render("Los archivos existentes se mueven a ~/.dotfiles-backup/links/ antes de enlazar.") + "\n\n")
	b.WriteString(SelectedStyle.Render("¿Crear backup completo en ~/.dotfiles-backup?") + "  [S/n]  " + HelpStyle.Render("q salir"))
	return b.String()
}

func (m Model) logView(limit int) string {
	lines := m.Log
	if len(lines) > limit {
		lines = lines[len(lines)-limit:]
	}
	var b strings.Builder
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "OK"):
			b.WriteString(SuccessStyle.Render("✓ " + strings.TrimPrefix(line, "OK ")))
		case strings.HasPrefix(line, "WARN"):
			b.WriteString(WarnStyle.Render("! " + strings.TrimPrefix(line, "WARN ")))
		case strings.HasPrefix(line, "ERR"):
			b.WriteString(ErrorStyle.Render("✗ " + strings.TrimPrefix(line, "ERR ")))
		default:
			b.WriteString(DimStyle.Render(line))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (m Model) installView() string {
	total := max(len(m.Selected), 1)
	s := TitleStyle.Render("Instalando") + "\n\n"
	s += m.Progress.ViewAs(float64(m.Step)/float64(total)) + DimStyle.Render(fmt.Sprintf("  %d/%d", m.Step, total)) + "\n\n"
	s += m.logView(max(m.Height-14, 5))
	if m.Step < len(m.Selected) {
		s += "\n" + m.Spinner.View() + " " + m.Selected[m.Step].Name + "..."
	}
	return s
}

func (m Model) doneView() string {
	title := SuccessStyle.Render("¡Listo!")
	if m.Failures > 0 {
		title = ErrorStyle.Render(fmt.Sprintf("Terminado con %d error(es)", m.Failures))
	} else if m.Warnings > 0 {
		title = WarnStyle.Render(fmt.Sprintf("Terminado con %d aviso(s)", m.Warnings))
	}
	s := title + "\n\n" + m.logView(max(m.Height-16, 5)) + "\n"

	var next []string
	for _, c := range m.Selected {
		switch c.ID {
		case "omarchy":
			next = append(next,
				"omarchy theme set wallhaven-gjr6ke   # aplicar un tema PACMAN",
				"fastfetch                            # ver el nuevo logo",
				"omarchy restart terminal             # recargar terminales")
		case "zsh":
			next = append(next, "exec zsh                             # recargar la shell")
		}
	}
	if len(next) > 0 {
		s += TitleStyle.Render("Siguientes pasos") + "\n"
		for _, n := range next {
			s += "  " + CodeStyle.Render(n) + "\n"
		}
	}
	return s + "\n" + HelpStyle.Render("enter salir")
}
