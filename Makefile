.PHONY: test lint build check-coverage coverage

# Test with coverage
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	golangci-lint run

build:
	go build -v ./...

# Check test coverage for specific packages
check-coverage:
	@echo "Checking test coverage..."
	@./scripts/check_coverage.sh

# CI target that runs everything
ci: lint test build check-coverage

# Helper target to generate coverage report
coverage: test
	go tool cover -html=coverage.txt -o coverage.html
