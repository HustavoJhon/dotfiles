package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/hustavojhon/dotfiles/internal/installer"
	"github.com/hustavojhon/dotfiles/internal/tui"
)

// version is set at build time: -ldflags "-X main.version=..."
var version = "dev"

func main() {
	var (
		backupOnly  = flag.Bool("backup-only", false, "crear backup y salir")
		dryRun      = flag.Bool("dry-run", false, "mostrar qué se haría sin tocar nada")
		list        = flag.Bool("list", false, "listar componentes disponibles en este sistema")
		status      = flag.Bool("status", false, "mostrar el estado de los enlaces")
		install     = flag.String("install", "", "instalar sin TUI: ids separados por coma, o \"all\"")
		noPackages  = flag.Bool("no-packages", false, "solo crear enlaces, no instalar paquetes")
		noBackup    = flag.Bool("no-backup", false, "con --install: no crear backup completo")
		showVersion = flag.Bool("version", false, "mostrar versión")
	)
	flag.BoolVar(showVersion, "v", false, "mostrar versión")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "PACMAN dotfiles installer %s\n\nUso: installer [flags]\n\n", version)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nSin flags abre el instalador interactivo (TUI).")
	}
	flag.Parse()

	if *showVersion {
		fmt.Println("PACMAN Installer " + version)
		return
	}

	inst, err := installer.NewInstaller()
	if err != nil {
		fail(err)
	}
	inst.DryRun = *dryRun
	inst.SkipPackages = *noPackages

	switch {
	case *backupOnly:
		path, err := inst.RunBackup(os.Stdout)
		if err != nil {
			fail(fmt.Errorf("backup falló: %w", err))
		}
		fmt.Println("Backup completado:", path)

	case *list:
		platform := inst.Distro.Name()
		if inst.Omarchy {
			platform += " (Omarchy)"
		}
		fmt.Printf("Sistema: %s · repo: %s\n\n", platform, inst.DotfilesDir)
		for _, c := range inst.AvailableComponents() {
			def := " "
			if c.Default {
				def = "*"
			}
			fmt.Printf(" %s %-10s %s\n", def, c.ID, c.Description)
		}
		fmt.Println("\n * = seleccionado por defecto")

	case *status:
		inst.Status(os.Stdout)

	case *install != "":
		os.Exit(runHeadless(inst, strings.Split(*install, ","), !*noBackup))

	default:
		prog := tea.NewProgram(tui.NewModel(inst, version), tea.WithAltScreen())
		if _, err := prog.Run(); err != nil {
			fail(err)
		}
	}
}

func runHeadless(inst *installer.Installer, ids []string, doBackup bool) int {
	comps, err := inst.FindComponents(ids)
	if err != nil {
		fail(err)
	}

	if doBackup && !inst.DryRun {
		if _, err := inst.RunBackup(os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "⚠ backup:", err)
		}
	}

	if missing := inst.MissingPackages(comps); len(missing) > 0 && !inst.SkipPackages && !inst.DryRun {
		fmt.Println("→ Paquetes a instalar:", strings.Join(missing, " "))
		cmd := installer.SudoCmd()
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, "⚠ sudo no disponible; solo se crearán enlaces")
			inst.SkipPackages = true
		}
	}

	failures := 0
	for _, c := range comps {
		fmt.Printf("\n▸ %s\n", c.Name)
		r := inst.InstallComponent(c)
		for _, line := range r.Log {
			fmt.Println(line)
		}
		for _, w := range r.Warnings {
			fmt.Println("  ⚠", w)
		}
		if r.Err != nil {
			failures++
			fmt.Fprintln(os.Stderr, "  ✗", r.Err)
		}
	}
	return min(failures, 1)
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(1)
}
