# PLAN DE REFACTORIZACIÓN — PACMAN

## RESUMEN EJECUTIVO

Transformar el repositorio actual de dotfiles (estructura plana con configuraciones duplicadas por distro) en una plataforma moderna con instalador TUI en Go, detección automática de distribución, sistema modular de paquetes y gestión inteligente de symlinks/backups.

---

## 1. DIAGNÓSTICO DEL ESTADO ACTUAL

### 1.1 Problemas Detectados

| Problema | Severidad | Archivos Afectados |
|---|---|---|
| **5 archivos .zshrc diferentes** sin coordinación | Alta | `/`, `Debian/`, `Manjaro/`, `Ubuntu/` (documentación), Windows |
| **Duplicación de Neovim settings.lua** (100% idéntico) | Alta | `Debian/nvim/`, `windows/nvim/` |
| **Script `_setup.sh` references `_install.sh` inexistente** | Alta | `_setup.sh` |
| **Usuario `sarah` hardcodeado en `_setup.sh`** | Alta | `_setup.sh:29` |
| **Typo `exa--icons` (falta espacio) duplicado** | Media | `Debian/.aliases`, `Manjaro/.zshrc` |
| **Typo path Windows `C:ProgramData...`** | Media | `windows/.config/user_profile.ps1` |
| **`.bashrc` referencia `~.bash_aliases` (sin `/`)** | Baja | `.bashrc` |
| **Sin gestor de dotfiles (ni Stow, chezmoi, yadm)** | Alta | Estructural |
| **Sin sistema de backup** | Alta | Estructural |
| **Sin detección de distro** | Alta | `_setup.sh` solo checkea apt |
| **Kitty config mínima, Tokyo Night activo, Nord inactivo** | Baja | `kitty/` |
| **Sin configuraciones Zellij, Wezterm, Ghostty, Starship, Fastfetch, Hyprland** | Media | README los menciona pero no existen |
| **Sin Makefile, sin Go, sin Python** | Alta | Estructural |
| **2 archivos keymaps.lua vacíos** | Baja | `Debian/`, `windows/` |
| **README desactualizado (template genérico)** | Media | `README.md` |

### 1.2 Activos a Preservar

| Activo | Valor | Nuevo Destino |
|---|---|---|
| **Modular ZSH config** (~/.dotfiles/zsh/) | Alto — mejor parte del repo | `configs/zsh/` sin cambios funcionales |
| **Kitty Tokyo Night config** | Alto | `configs/kitty/` |
| **Kitty Nord theme** | Bajo | `configs/kitty/themes/` |
| **Tmux config** (.tmux.conf) | Medio | `configs/tmux/tmux.conf` |
| **P10k config** (root/.p10k.zsh) | Medio | `configs/zsh/p10k.zsh` |
| **Banner ASCII** (banner.txt) | Medio — identidad visual | `assets/banners/` |
| **Ghost logo** (ghost.png) | Medio — branding | `assets/logo/` |
| **Wallpapers** | Medio | `assets/wallpapers/` |
| **Neovim config** (Debian) | Alto — NO TOCAR | `configs/nvim/` (preservado intacto) |
| **Windows configs** | Bajo — plataforma separada | `windows/` (igual, intacto) |
| **Fonts FiraCode** | Bajo | `assets/fonts/` |

---

## 2. NUEVA ARQUITECTURA

### 2.1 Árbol Completo Post-Refactorización

```
dotfiles/
│
├── cmd/
│   └── installer/
│       └── main.go                    # Entrypoint TUI installer
│
├── internal/
│   ├── distro/
│   │   ├── distro.go                  # Distro detection interface
│   │   ├── arch.go                    # ArchAdapter
│   │   ├── debian.go                  # DebianAdapter
│   │   └── fedora.go                  # FedoraAdapter
│   │
│   ├── installer/
│   │   ├── installer.go               # Core install logic
│   │   ├── modules.go                 # Module definitions
│   │   └── steps.go                   # Install steps
│   │
│   ├── packages/
│   │   ├── packages.go                # Package manager abstraction
│   │   └── lists.go                   # Package lists per component
│   │
│   ├── backup/
│   │   └── backup.go                  # Backup system
│   │
│   ├── symlink/
│   │   └── symlink.go                 # Symlink management
│   │
│   ├── tui/
│   │   ├── styles.go                  # Lip Gloss styles
│   │   ├── welcome.go                 # Welcome screen
│   │   ├── menu.go                    # Component selection
│   │   ├── progress.go                # Install progress
│   │   └── model.go                   # Bubble Tea model
│   │
│   ├── config/
│   │   └── config.go                  # Config loading
│   │
│   └── assets/
│       └── ghost.go                   # Pacman Ghost ASCII
│
├── configs/
│   ├── zsh/
│   │   ├── .zshrc                     # Modular zshrc (fusionado)
│   │   ├── aliases.zsh                # De aliases actual
│   │   ├── exports.zsh                # De exports actual
│   │   ├── functions.zsh              # De functions actual
│   │   ├── paths.zsh                  # De paths actual
│   │   └── p10k.zsh                   # De root/.p10k.zsh
│   │
│   ├── zellij/
│   │   ├── config.kdl                 # Nueva configuración
│   │   └── layouts/
│   │       └── main.kdl
│   │
│   ├── kitty/
│   │   ├── kitty.conf                 # Actual + mejoras
│   │   ├── color.ini                  # Tokyo Night (actual)
│   │   └── themes/
│   │       └── nord.ini              # Nord (migrado)
│   │
│   ├── wezterm/
│   │   └── .wezterm.lua              # Nueva configuración
│   │
│   ├── ghostty/
│   │   └── config                    # Nueva configuración
│   │
│   ├── starship/
│   │   └── starship.toml             # Nueva configuración
│   │
│   ├── fastfetch/
│   │   └── config.jsonc              # Nueva configuración
│   │
│   ├── hypr/
│   │   ├── hyprland.conf             # Nueva configuración
│   │   ├── hyprlock.conf
│   │   └── hypridle.conf
│   │
│   ├── tmux/
│   │   └── .tmux.conf                # Actual migrado
│   │
│   └── nvim/                          # ← INTACTO (solo copiado)
│       ├── init.lua
│       └── lua/
│           ├── keymaps.lua
│           ├── plugins.lua
│           └── settings.lua
│
├── scripts/
│   ├── install/
│   │   ├── bootstrap.sh              # Bootstrap mínimo (Go install)
│   │   └── post-install.sh           # Post-install hooks
│   │
│   ├── python/
│   │   ├── update_fonts.py           # Font cache updater
│   │   ├── wallpaper_manager.py      # Wallpaper setter/rotator
│   │   ├── theme_switcher.py         # Theme swapper
│   │   └── package_audit.py          # Package audit/report
│   │
│   └── maintenance/
│       ├── cleanup.sh                # Temp file cleanup
│       └── healthcheck.sh            # System health check
│
├── shaders/
│   └── glsl/
│       ├── kitty-blur.glsl           # Optional kitty shader
│       └── hypr-luma.glsl            # Hyprland shader
│
├── assets/
│   ├── logo/
│   │   └── ghost.png                 # Actual (migrado)
│   ├── banners/
│   │   └── banner.txt                # Actual (migrado)
│   ├── wallpapers/                   # Actuales (migrados)
│   └── fonts/
│       └── FiraCode/                 # Actuales (migrados)
│
├── pkg/                              # Reserved for Go packages
│
├── docs/
│   ├── ARCHITECTURE.md
│   ├── MODULES.md
│   └── ROADMAP.md
│
├── windows/                 # ← INTACTO (sin cambios)
│
├── go.mod
├── go.sum
├── Makefile
├── install.sh
├── PLAN.md                  ← (este archivo)
└── README.md
```

### 2.2 Archivos a Eliminar

| Archivo | Razón |
|---|---|
| `Distributions/` | Reemplazado por sistema modular de distros |
| `root/` | Contenido migrado a `configs/zsh/` |
| `doc/` (contenido no arquitectural) | Migrar assets, descartar docs redundantes |
| `_setup.sh` | Reemplazado por `install.sh` + Go TUI |
| `.bash_aliases` | Legacy, no usado con ZSH |
| `.bashrc` | Legacy |
| `system.md` | Notas personales, mover a `docs/` si aplica |

### 2.3 Archivos a Preservar Intactos

| Archivo | Condición |
|---|---|
| `windows/` | 100% intacto |
| `configs/nvim/` | 100% intacto (contenido copiado de `Debian/nvim/`) |
| `LICENSE` | Intacto |

---

## 3. FASES DE IMPLEMENTACIÓN

### FASE 1: Go TUI Installer (Días 1-2)

#### 3.1 Estructura del Módulo Go

```
cmd/installer/main.go
internal/tui/
  ├── model.go          # Bubble Tea model principal
  ├── styles.go         # Lip Gloss theme (morado/cyan/negro)
  ├── welcome.go        # Pacman Ghost ASCII + logo
  ├── menu.go           # Bubbles multiselect
  └── progress.go       # Barra de progreso
internal/distro/
  ├── distro.go         # Interfaz DistroAdapter
  ├── arch.go           # pacman -S
  ├── debian.go         # apt-get install
  └── fedora.go         # dnf install
internal/installer/
  ├── installer.go      # Orquestador
  ├── modules.go        # Definición de módulos
  └── steps.go          # Pasos de instalación
internal/backup/
  └── backup.go         # ~/.dotfiles-backup/timestamp/
internal/symlink/
  └── symlink.go        # ln -sf wrapper
internal/config/
  └── config.go         # Carga de configuración
internal/assets/
  └── ghost.go          # ASCII art del Pacman Ghost
```

#### 3.2 Flujo TUI

```
[Welcome Screen]
  ├── Mostrar PACMAN GHOST ASCII
  ├── Mostrar "PACMAN INSTALLER v1.0"
  └── Press Enter to continue

[Distro Detection]
  ├── Leer /etc/os-release
  ├── Detectar: Arch / Debian / Fedora / CachyOS / EndeavourOS / Ubuntu
  ├── Mostrar: "Detected: Arch Linux"
  └── Press Enter to continue

[Component Selection]  ← Bubbles multiselect
  ├── [x] ZSH
  ├── [x] Zellij
  ├── [ ] Kitty
  ├── [ ] WezTerm
  ├── [ ] Ghostty
  ├── [x] Fastfetch
  ├── [x] Starship
  ├── [ ] Hyprland
  ├── [ ] Themes
  └── [ ] Wallpapers

[Backup Prompt]
  └── "Create backup before installing? [Y/n]"

[Install Progress]
  ├── [████████░░] Installing ZSH...
  ├── [████████░░] Installing Zellij...
  └── Done!

[Summary]
  ├── ✓ ZSH installed
  ├── ✓ Zellij installed
  ├── ✗ Kitty skipped
  └── Press Enter to exit
```

#### 3.3 Tema Visual

```go
// internal/tui/styles.go
var (
    purple   = lipgloss.Color("#7C3AED")   // Morado principal
    cyan     = lipgloss.Color("#06B6D4")   // Cyan acento
    darkBg   = lipgloss.Color("#0F0F0F")   // Negro fondo
    white    = lipgloss.Color("#FFFFFF")
    gray     = lipgloss.Color("#6B7280")
    
    titleStyle    = lipgloss.NewStyle().Foreground(purple).Bold(true)
    accentStyle   = lipgloss.NewStyle().Foreground(cyan)
    selectedStyle = lipgloss.NewStyle().Foreground(purple).Bold(true)
    progressStyle = lipgloss.NewStyle().Foreground(cyan)
)
```

#### 3.4 Mascota: Pacman Ghost

```
         .--.
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
  |_____|_______|_____|
```

---

### FASE 2: Sistema Modular de Distros (Día 2)

#### 4.1 Interfaz DistroAdapter

```go
type DistroAdapter interface {
    Name() string
    InstallPackage(pkg string) error
    InstallPackages(pkgs []string) error
    RemovePackage(pkg string) error
    UpdateSystem() error
    PackageManager() string
}
```

#### 4.2 Implementaciones

| Adapter | PM | Comando |
|---|---|---|
| ArchAdapter | pacman | `sudo pacman -S --noconfirm` |
| DebianAdapter | apt | `sudo apt-get install -y` |
| FedoraAdapter | dnf | `sudo dnf install -y` |

Detección vía:
```go
func DetectDistro() (DistroAdapter, error) {
    data, _ := os.ReadFile("/etc/os-release")
    idRE := regexp.MustCompile(`^ID=["']?(\w+)["']?$`)
    idLikeRE := regexp.MustCompile(`^ID_LIKE=["']?(\w+)["']?$`)
    // match and return appropriate adapter
}
```

---

### FASE 3: Migración de Configuraciones (Día 2-3)

#### 5.1 ZSH — Fusión de los 5 .zshrc en Uno Modular

**Estrategia:** Tomar `Distributions/Debian/.zshrc` como base (es el más completo y moderno) y mejorarlo.

**Cambios específicos:**
1. Mantener la carga modular (~/.dotfiles/zsh/*.zsh) exactamente igual
2. Agregar detección automática de distro para Homebrew paths
3. Migrar `root/.p10k.zsh` a `configs/zsh/p10k.zsh`
4. El instalador copiará `configs/zsh/` a `~/.dotfiles/`

**Archivos finales en `configs/zsh/`:**
- `.zshrc` — Carga modular (basado en Debian)
- `aliases.zsh` — Actual + fixes (typo exa, mejor organización)
- `exports.zsh` — Actual + mejoras
- `functions.zsh` — Actual + mejoras
- `paths.zsh` — Actual + mejoras
- `p10k.zsh` — Migrado de `root/.p10k.zsh`

#### 5.2 Kitty — Migrar + Mejorar

**Actual:** `/.config/kitty/kitty.conf` + `color.ini`
**Nuevo:** `configs/kitty/kitty.conf` + `color.ini` + `themes/nord.ini`

Mejoras:
- Agregar blur (Linux soportado)
- Agregar Catppuccin Mocha como opción temática
- Mantener Tokyo Night como default
- Optimizar rendimiento (sync_to_monitor, repaint_delay)

#### 5.3 Tmux — Migrar Directamente

`configs/tmux/.tmux.conf` — exactamente el mismo contenido, solo movido.

#### 5.4 Neovim — COPIAR, NO TOCAR

Copiar `Distributions/Debian/nvim/` → `configs/nvim/` literalmente.
El Neovim queda intacto.

---

### FASE 4: Gestión de Backups y Symlinks (Día 3)

#### 6.1 Sistema de Backup

```go
// internal/backup/backup.go
func CreateBackup() (string, error) {
    timestamp := time.Now().Format("20060102_150405")
    backupDir := filepath.Join(homeDir, ".dotfiles-backup", timestamp)
    
    // Backup paths
    paths := []string{
        ".zshrc",
        ".config/kitty",
        ".tmux.conf",
        ".p10k.zsh",
        ".local/share/zsh",
    }
    
    for _, p := range paths {
        src := filepath.Join(homeDir, p)
        if exists(src) {
            copyTree(src, filepath.Join(backupDir, p))
        }
    }
    
    return backupDir, nil
}
```

#### 6.2 Sistema de Symlinks

```go
// internal/symlink/symlink.go
type SymlinkManager struct {
    dotfilesDir string
    dryRun      bool
}

func (sm *SymlinkManager) Link(src, dest string) error {
    // dest ~/.zshrc -> src ~/dotfiles/configs/zsh/.zshrc
    absSrc := filepath.Join(sm.dotfilesDir, src)
    absDest := filepath.Join(homeDir, dest)
    
    backup existing if not symlink
    os.Remove(absDest)
    os.Symlink(absSrc, absDest)
}
```

---

### FASE 5: Scripts Auxiliares (Día 3)

#### 7.1 Python Scripts

**`scripts/python/update_fonts.py`**: 
- Escanea `assets/fonts/`
- Corre `fc-cache` 
- Verifica Nerd Font patches

**`scripts/python/wallpaper_manager.py`**:
- Establece wallpapers desde `assets/wallpapers/`
- Soporte: Hyprpaper, feh, gsettings
- Rotación automática

**`scripts/python/theme_switcher.py`**:
- Cambia entre Tokyo Night / Nord / Catppuccin
- Actualiza Kitty, Wezterm, Bat, FZF simultáneamente

**`scripts/python/package_audit.py`**:
- Compara paquetes instalados vs lista deseada
- Reporta faltantes y extraños

---

### FASE 6: GLSL Shaders (Día 4)

#### 8.1 Shaders

**`shaders/glsl/kitty-blur.glsl`**:
```
// Gaussian blur passthrough for kitty
```

**`shaders/glsl/hypr-luma.glsl`**:
```
// Luminance-based effect for Hyprland
```

Son archivos decorativos, desacoplados del core del instalador.

---

### FASE 7: Makefile y README (Día 4)

#### 9.1 Comandos Makefile

```makefile
build:        # go build ./cmd/installer/
install:      # go run ./cmd/installer/ (o binary)
clean:        # rm -rf bin/ go.sum
backup:       # go run ./cmd/installer/ --backup-only
update:       # git pull + rebuild
fmt:          # go fmt ./...
lint:         # golangci-lint run ./...
test:         # go test ./...
```

#### 9.2 README.md

README profesional con:
- Logo + shields
- Instalación (curl | bash one-liner)
- Screenshots TUI
- Arquitectura
- Módulos disponibles
- Roadmap
- Contributing

---

## 4. CAMBIOS ESPECÍFICOS POR ARCHIVO

### 4.1 Archivos a Crear

| Archivo | Propósito | Líneas Estimadas |
|---|---|---|
| `cmd/installer/main.go` | Entrypoint | 30 |
| `internal/tui/model.go` | Bubble Tea model | 150 |
| `internal/tui/styles.go` | Lip Gloss theme | 80 |
| `internal/tui/welcome.go` | Welcome + Pacman Ghost | 100 |
| `internal/tui/menu.go` | Component multiselect | 120 |
| `internal/tui/progress.go` | Progress bars | 80 |
| `internal/distro/distro.go` | Interface | 40 |
| `internal/distro/arch.go` | Arch adapter | 80 |
| `internal/distro/debian.go` | Debian adapter | 80 |
| `internal/distro/fedora.go` | Fedora adapter | 80 |
| `internal/installer/installer.go` | Orchestrator | 120 |
| `internal/installer/modules.go` | Module definitions | 80 |
| `internal/installer/steps.go` | Install steps | 100 |
| `internal/backup/backup.go` | Backup system | 100 |
| `internal/symlink/symlink.go` | Symlink manager | 80 |
| `internal/config/config.go` | Config loader | 60 |
| `internal/assets/ghost.go` | Pacman Ghost ASCII | 50 |
| `configs/zellij/config.kdl` | Zellij config | 80 |
| `configs/zellij/layouts/main.kdl` | Layout | 40 |
| `configs/wezterm/.wezterm.lua` | Wezterm config | 80 |
| `configs/ghostty/config` | Ghostty config | 30 |
| `configs/starship/starship.toml` | Starship prompt | 60 |
| `configs/fastfetch/config.jsonc` | Fastfetch config | 40 |
| `configs/hypr/hyprland.conf` | Hyprland WM | 100 |
| `configs/hypr/hyprlock.conf` | Hyprlock | 40 |
| `configs/hypr/hypridle.conf` | Hypridle | 30 |
| `scripts/python/update_fonts.py` | Font updater | 60 |
| `scripts/python/wallpaper_manager.py` | Wallpaper manager | 80 |
| `scripts/python/theme_switcher.py` | Theme swapper | 100 |
| `scripts/python/package_audit.py` | Package audit | 80 |
| `shaders/glsl/kitty-blur.glsl` | Kitty blur shader | 30 |
| `shaders/glsl/hypr-luma.glsl` | Hyprland shader | 30 |
| `docs/ARCHITECTURE.md` | Architecture docs | 100 |
| `docs/MODULES.md` | Module documentation | 80 |
| `docs/ROADMAP.md` | Roadmap | 60 |
| `Makefile` | Build system | 40 |
| `install.sh` | Bootstrap installer | 50 |
| `go.mod` | Go module | 5 |

### 4.2 Archivos a Modificar

| Archivo | Cambio |
|---|---|
| `.gitignore` | Agregar `bin/`, `*.out`, `__pycache__/` |
| `README.md` | Reescribir completamente |

### 4.3 Archivos a Mover

| Origen | Destino |
|---|---|
| `Distributions/Debian/.zshrc` | `configs/zsh/.zshrc` (version mejorada) |
| `Distributions/Debian/.dotfiles/zsh/*` | `configs/zsh/*` |
| `Distributions/Debian/.dotfiles/banner.txt` | `assets/banners/banner.txt` |
| `Distributions/Debian/nvim/` | `configs/nvim/` (intacto) |
| `root/.p10k.zsh` | `configs/zsh/p10k.zsh` |
| `.tmux.conf` | `configs/tmux/.tmux.conf` |
| `.config/kitty/` | `configs/kitty/` |
| `doc/src/images/ghost.png` | `assets/logo/ghost.png` |
| `doc/Fonts/FiraCode/` | `assets/fonts/FiraCode/` |
| `doc/Wallpaper/` | `assets/wallpapers/` |

### 4.4 Archivos a Eliminar

| Archivo | Razón |
|---|---|
| `Distributions/` | Reemplazado por sistema modular |
| `root/` | Contenido migrado |
| `_setup.sh` | Reemplazado por Go TUI + install.sh |
| `.bash_aliases` | Legacy |
| `.bashrc` | Legacy |
| `.config/` (root) | Migrado a configs/ |
| `doc/` (estructural) | Contenido migrado a assets/ y docs/ |
| `system.md` | Notas personales no arquitecturales |

---

## 5. LO QUE NO CAMBIA

### 5.1 Intocable

| Componente | Razón |
|---|---|
| **Configuración de Neovim** | Solicitud explícita del usuario |
| **Plugins de Neovim** | Solicitud explícita del usuario |
| **Estructura interna de Neovim** | Solicitud explícita del usuario |
| **Dotfiles Windows** | Plataforma separada, intacta |
| **Licencia** | Legal, se mantiene |

### 5.2 Preservación de Personalización

| Elemento | Preservación |
|---|---|
| Aliases personalizados | Migrados 1:1 |
| Exports de entorno | Migrados 1:1 |
| Funciones ZSH | Migradas 1:1 |
| Colores Tokyo Night | Default mantenido |
| Banner ASCII | Migrado a assets/ |
| Logo ghost.png | Migrado a assets/logo/ |

---

## 6. VERIFICACIÓN

### 6.1 Tests

```bash
go test ./...           # Go unit tests
make lint               # golangci-lint
make build              # go build
```

### 6.2 Verificación Manual

```bash
./install.sh            # Bootstrap → Go TUI
make backup             # Backup-only mode
make clean              # Clean artifacts
```

---

## 7. RIESGOS Y MITIGACIONES

| Riesgo | Mitigación |
|---|---|
| Ruptura de config existente | Backup automático pre-instalación |
| Go no instalado en sistema | Bootstrap con install.sh (curl Go) |
| Pacman Ghost no gusta | El usuario lo pidió explícitamente |
| Duplicación en migración | Symlinks, no copias |
| Dependencias rotas | Instalación idempotente |

---

## 8. TIMELINE ESTIMADO

| Fase | Duración | Dependencias |
|---|---|---|
| Análisis | COMPLETADO | — |
| Plan | COMPLETADO | Análisis |
| Go TUI scaffolding | 2 sesiones | Plan |
| Distro detection | 1 sesión | Go scaffolding |
| Config migration | 2 sesiones | Plan |
| Backup/symlink | 1 sesión | Go scaffolding |
| Python scripts | 1 sesión | — |
| GLSL shaders | 0.5 sesión | — |
| Makefile/README | 0.5 sesión | Todo lo anterior |
| Testing/QA | 1 sesión | Todo lo anterior |

**Total estimado: ~9 sesiones de trabajo**
