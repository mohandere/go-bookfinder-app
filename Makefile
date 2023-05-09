BUILD_DIR := bin
TOOLS_DIR := tools

default: all

all: clean lint test build run

dep:
	go mod download
	go mod tidy

.PHONY: $(BUILD_DIR)/server
bin/server: cmd/server/*.go
	CGO_ENABLED=1 go build -ldflags="-s -w" -o ./bin/server ./cmd/server/

.PHONY: build
build: bin/server

.PHONY: run
run: build
	bin/server

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(TOOLS_DIR)

tools/golangci-lint/golangci-lint:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b tools/golangci-lint latest

.PHONY: lint
lint: $(TOOLS_DIR)/golangci-lint/golangci-lint
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run ./...

.PHONY: test
test:
	go test -race -cover -coverprofile=coverage.txt -covermode=atomic ./...
