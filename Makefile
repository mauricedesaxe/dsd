# the default ARCH for tailwindcss is macos-x64
ARCH ?= macos-arm64

# Setup the project for the first time
setup:
	@echo "Setting up project..."

ENV ?= dev

kill:
	@echo "Killing process on port 3000..."
	@kill -9 $(lsof -i:3000 -t) || echo "No process running on port 3000."

dev:
	@echo "Starting development server..."
	@air

build:
	@echo "Building Go binary..."
	@go build -o bin/app
	@chmod +x ./bin/app
	@echo "Go binary built."

	@echo "Project built."

run:
	@echo "Running project..."
	@./bin/app || echo "Failed to run the application. Check if the binary exists and has execution permissions."
