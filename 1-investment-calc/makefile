# Binary name
BINARY_NAME=investment_calculator

# Go related variables
GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin

# Main entry point
MAIN_FILE=investment_calculator.go

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.DEFAULT_GOAL := help

## Build: Build the application
build:
	@echo "Building..."
	go build -o $(BINARY_NAME) $(MAIN_FILE)

## Run: Run the application
run:
	@echo "Running..."
	go run $(MAIN_FILE)

## Clean: Remove binary
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f $(BINARY_NAME).exe

## Test: Run tests
test:
	@echo "Testing..."
	go test ./...

## Fmt: Format code
fmt:
	@echo "Formatting..."
	go fmt ./...

## Vet: Run go vet
vet:
	@echo "Vetting..."
	go vet ./...

## Build-and-run: Build and run the application
build-and-run: build
	@echo "Running built binary..."
	./$(BINARY_NAME)

## Help: Display this help screen
help:
	@echo "Usage:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build run clean test fmt vet build-and-run help