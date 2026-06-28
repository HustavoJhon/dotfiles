package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/hustavojhon/dotfiles/internal/installer"
	"github.com/hustavojhon/dotfiles/internal/tui"
)

func main() {
	inst, err := installer.NewInstaller()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--backup-only":
			path, err := inst.RunBackup()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Backup failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Backup completed: %s\n", path)
			return

		case "--version", "-v":
			fmt.Println("PACMAN Installer v1.0")
			return

		case "--help", "-h":
			fmt.Println("PACMAN Installer")
			fmt.Println("  Interactive TUI installer for dotfiles")
			fmt.Println("")
			fmt.Println("Flags:")
			fmt.Println("  --backup-only   Create backup and exit")
			fmt.Println("  --version, -v   Show version")
			fmt.Println("  --help, -h      Show this help")
			return
		}
	}

	model := tui.NewModel(inst)
	prog := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := prog.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
