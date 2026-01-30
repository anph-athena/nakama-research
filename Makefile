.PHONY: help build run test clean fmt vet docker-up docker-down docker-logs

help:
	@echo "Available commands:"
	@echo "  make build          - Build the Nakama plugin (main.so)"
	@echo "  make run            - Build and run the plugin"
	@echo "  make test           - Run all tests"
	@echo "  make fmt            - Format code"
	@echo "  make vet            - Run go vet"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make docker-up      - Start Docker services"
	@echo "  make docker-down    - Stop Docker services"
	@echo "  make docker-logs    - View Nakama logs"
	@echo "  make docker-restart - Restart Nakama service"

build:
	go build -buildmode=plugin -trimpath -o /tmp/main.so main.go

run: build
	@echo "Plugin built at /tmp/main.so"

test:
	go test -v -race -coverprofile=coverage.out ./...

test-coverage: test
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: fmt vet
	@echo "Linting complete"

clean:
	rm -f /tmp/main.so
	rm -f coverage.out coverage.html

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f nakama

docker-restart:
	docker-compose restart nakama

tidy:
	go mod tidy

deps:
	go mod download

setup: tidy docker-up
	@echo "Setup complete. Starting services..."

all: lint build
	@echo "Build complete"
