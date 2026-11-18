package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/hustavojhon/dotfiles/internal/installer"
)

type fakeDistro struct{}

func (fakeDistro) Name() string                { return "Omarchy" }
func (fakeDistro) PackageManager() string      { return "pacman" }
func (fakeDistro) IsInstalled(string) bool     { return true }
func (fakeDistro) InstallPackage(string) error { return nil }

func key(s string) tea.KeyMsg {
	if s == "enter" {
		return tea.KeyMsg{Type: tea.KeyEnter}
	}
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(s)}
}

func step(t *testing.T, m Model, msg tea.Msg) (Model, tea.Cmd) {
	t.Helper()
	next, cmd := m.Update(msg)
	return next.(Model), cmd
}

func TestFlowOnOmarchy(t *testing.T) {
	inst := &installer.Installer{
		HomeDir: t.TempDir(), DotfilesDir: "../..",
		Distro: fakeDistro{}, Omarchy: true, DryRun: true,
	}
	m := NewModel(inst, "test")
	m, _ = step(t, m, tea.WindowSizeMsg{Width: 100, Height: 50})

	if v := m.View(); !strings.Contains(v, "Omarchy detectado") || !strings.Contains(v, "dry-run") {
		t.Fatalf("welcome view missing platform info:\n%s", v)
	}

	m, _ = step(t, m, key("enter"))
	if m.Screen != ScreenMenu || !strings.Contains(m.View(), "Omarchy") {
		t.Fatal("expected menu with omarchy component")
	}

	m, _ = step(t, m, key("n")) // deselect all
	m, _ = step(t, m, key("enter"))
	if m.Screen != ScreenMenu {
		t.Fatal("enter with nothing selected must stay in menu")
	}

	m, _ = step(t, m, key(" ")) // select first (omarchy)
	m, _ = step(t, m, key("enter"))
	if m.Screen != ScreenBackup || len(m.Selected) != 1 || m.Selected[0].ID != "omarchy" {
		t.Fatalf("screen=%v selected=%v", m.Screen, m.Selected)
	}
	if !strings.Contains(m.View(), "ya están instalados") {
		t.Fatal("summary should say packages are installed")
	}

	// Skip backup; no packages missing → install directly, no sudo.
	m, cmd := step(t, m, key("n"))
	if m.Screen != ScreenInstall || cmd == nil {
		t.Fatal("expected install to start")
	}
	m, _ = step(t, m, cmd()) // run the dry-run install of omarchy
	if m.Screen != ScreenDone || m.Failures != 0 {
		t.Fatalf("screen=%v failures=%d log=%v", m.Screen, m.Failures, m.Log)
	}
	if v := m.View(); !strings.Contains(v, "omarchy theme set") {
		t.Fatalf("done view missing next steps:\n%s", v)
	}
}
