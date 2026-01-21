# Makefile for auth-service

# Variables
APP_NAME = auth-service
MAIN_FILE = cmd/main.go
BINARY = bin/${APP_NAME}

# Build the application
build:
	go build -o ${BINARY} ${MAIN_FILE}

# Run the application
run:
	go run ${MAIN_FILE}

# Build and run the application
build-run: build
	./${BINARY}

# Install dependencies
deps:
	go mod tidy

# Test the application
test:
	go test -v ./...

# Test with coverage
test-cover:
	go test -cover ./...

# Clean build files
clean:
	rm -f ${BINARY}

# Docker build
docker-build:
	docker build -t ${APP_NAME} .

# Docker run
docker-run:
	docker run -p 8080:8080 ${APP_NAME}

# Docker compose up
docker-up:
	docker-compose up -d

# Docker compose down
docker-down:
	docker-compose down

# Air live reload
air-run:
	air

# Help
help:
	@echo "Available commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  build-run   - Build and run the application"
	@echo "  deps        - Install dependencies"
	@echo "  test        - Test the application"
	@echo "  test-cover  - Test with coverage"
	@echo "  clean       - Clean build files"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run  - Run Docker container"
	@echo "  docker-up   - Start Docker Compose"
	@echo "  docker-down - Stop Docker Compose"
	@echo "  air-run     - Run with Air live reload"
	@echo "  help        - Show this help"

.PHONY: build run build-run deps test test-cover clean docker-build docker-run docker-up docker-down air-run help