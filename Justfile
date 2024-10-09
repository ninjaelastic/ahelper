# List available commands
default:
    @just --list

# Build the project
build:
    go build -o bin/aich cmd/aich/main.go

# Run the project
run *ARGS:
    go run cmd/aich/main.go {{ARGS}}

# Run tests
test:
    go test ./...

# Run tests with coverage
test-coverage:
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out

# Format code
fmt:
    go fmt ./...

# Lint code
lint:
    golangci-lint run

# Clean build artifacts
clean:
    rm -rf bin

# Install dependencies
deps:
    go mod tidy

# Update dependencies
update-deps:
    go get -u ./...
    go mod tidy

# Generate documentation
docs:
    go doc -all > docs/api.txt

# Build and run
build-run *ARGS: build
    ./bin/aich {{ARGS}}

# Full check: format, lint, test
check: fmt lint test

# Setup project (install dependencies, build)
setup: deps build