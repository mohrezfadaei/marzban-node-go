BIN = marzban-node-go
BIN_DIR = build/bin
BIN_PATH = build/bin/$(BIN)

build:
		echo "Building the Go application..."
		go build -o $(BIN_PATH) cmd/main.go

build-linux:
		echo "Building the Go application for Linux..."
		GOOS=linux GOARCH=amd64 go build -o $(BIN_PATH)-linux-amd64 cmd/main.go
		GOOS=linux GOARCH=arm64 go build -o $(BIN_PATH)-linux-arm64 cmd/main.go

clean:
		echo "Cleaning up..."
		rm -f $(BIN_PATH)

lint:
		echo "Running linter..."
		golangci-lint run

fmt:
		echo "Formatting the Go code..."
		go fmt ./...


help:
		echo "Usage:"
		echo "  make build            Build the Go application for the host architecture"
		echo "  make build-linux      Build the Go application for Linux"
		echo "  make clean            Clean the build artifacts"
		echo "  make lint             Run Go linter"
		echo "  make fmt              Format the Go code"
		echo "  make help             Display this help message"