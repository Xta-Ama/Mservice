# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main main.go

# Run the application
run:
	@go run main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

# Live Reload
swagger-generate:
	@if command -v swag init > /dev/null; then \
	    swag init; \
	    echo "Successfully Generated RESTful API documentation!";\
	else \
	    read -p "Go's 'swaggo' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/swaggo/swag/cmd/swag@latest; \
	        swag init; \
	        echo "Watching...";\
	    else \
	        echo "Successfully Generated RESTful API documentation!"; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
