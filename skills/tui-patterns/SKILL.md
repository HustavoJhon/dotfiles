# TUI Patterns (Bubble Tea)

## Estructura

El instalador usa Bubble Tea con el patrón Model-Update-View:

- `model.go` — Estado de la aplicación, screens, datos
- `update.go` — Manejo de eventos y navegación
- `view.go` — Renderizado de cada pantalla
- `styles.go` — Estilos Lip Gloss (tema gruvbox)

## Cómo agregar una nueva pantalla

1. Agregar constante `Screen` en `model.go`
2. Agregar estado al `Model` struct si es necesario
3. Agregar handler en `update.go`
4. Agregar caso en `view.go`
5. Agregar título en `GetScreenTitle()`
