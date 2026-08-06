goimports:=$(shell go env GOPATH)/bin/goimports

# Download the import optimizer (and code formatter)
$(goimports):
	@echo "Installing goimports"
	go install golang.org/x/tools/cmd/goimports@v0.41.0

build:
	@echo "Building Adyen Go API library"
	go build -o bin/main main.go

run:
	@echo "Building and running Adyen Go API library"
	go run main.go

# Our unit tests
test:
	@echo "Testing Adyen Go API library"
	go test ./...

# Requeries a .env file with ADYEN_API_KEY and ADYEN_MERCHANT
# Note that some tests might depended on permissions associated to the ADYEN_API_KEY
#
# These tests always run for release pipelines
integration-tests:
	@echo "Running integration tests"
	go test ./... -tags integration

clean:
	@echo "Cleaning binaries, build and test caches"
	@rm -rf bin/
	go clean -cache -testcache

fmt: $(goimports)
	goimports -w .

verify: fmt build run test
