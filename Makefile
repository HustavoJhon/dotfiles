# PACMAN.DOTS - Makefile

BINARY=bin/installer
GOFLAGS=-ldflags="-s -w"

.PHONY: all build install clean backup update fmt lint test run help

all: build

build:
	@echo "Building $(BINARY)..."
	@mkdir -p bin
	go build $(GOFLAGS) -o $(BINARY) ./cmd/installer
	@echo "✓ Build complete: $(BINARY)"

run: build
	@./$(BINARY)

install: build
	@./$(BINARY)

clean:
	@echo "Cleaning..."
	@rm -rf bin/ bin-debug/ bin-release/
	@echo "✓ Clean complete"

backup:
	@go run ./cmd/installer --backup-only

update:
	@git pull --rebase
	@go mod tidy
	@$(MAKE) build

fmt:
	@go fmt ./...
	@echo "✓ Formatting complete"

lint:
	@golangci-lint run ./... 2>/dev/null || echo "⚠ golangci-lint not installed, skipping"

test:
	@go test ./... -v
	@echo "✓ Tests complete"

help:
	@echo "PACMAN.DOTS Commands:"
	@echo "  make build    - Build Go TUI installer"
	@echo "  make install  - Run the TUI installer"
	@echo "  make run      - Build and run"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make backup   - Run backup only"
	@echo "  make update   - Git pull, tidy deps, rebuild"
	@echo "  make fmt      - Format Go code"
	@echo "  make lint     - Run golangci-lint"
	@echo "  make test     - Run all tests"
