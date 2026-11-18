# PACMAN - Makefile

BINARY  := bin/installer
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: all build run install dry-run status list omarchy omarchy-sync backup update fmt lint test e2e clean help

all: build

build: ## Compilar el instalador
	@mkdir -p bin
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/installer
	@echo "✓ $(BINARY) ($(VERSION))"

run install: build ## Abrir el instalador TUI
	@./$(BINARY)

dry-run: build ## TUI sin modificar nada
	@./$(BINARY) --dry-run

status: build ## Estado de los enlaces
	@./$(BINARY) --status

list: build ## Componentes disponibles en este sistema
	@./$(BINARY) --list

omarchy: build ## Instalar solo el módulo Omarchy (sin TUI)
	@./$(BINARY) --install omarchy

omarchy-sync: ## Copiar la config viva de Omarchy al repo
	@./scripts/omarchy/sync.sh

backup: build ## Backup de configs actuales
	@./$(BINARY) --backup-only

update: ## git pull + rebuild
	git pull --rebase
	go mod tidy
	@$(MAKE) build

fmt: ## Formatear código Go
	go fmt ./...

lint: ## go vet + golangci-lint (si existe)
	go vet ./...
	@command -v golangci-lint >/dev/null && golangci-lint run ./... || echo "⚠ golangci-lint no instalado, solo go vet"

test: ## Tests Go
	go test ./...

e2e: ## Tests e2e en Docker
	docker build -t pacman-test -f Dockerfile.test .
	docker run --rm pacman-test

clean: ## Borrar binarios
	rm -rf bin/

help: ## Esta ayuda
	@grep -hE '^[a-z -]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  \033[33m%-14s\033[0m %s\n", $$1, $$2}'
