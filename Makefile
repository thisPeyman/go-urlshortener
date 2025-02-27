# Paths
PROTO_SRC := api/src
PROTO_OUT := api/
CMD_DIR := cmd
SERVICE := id-gen

# Tools
PROTOC := protoc
PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc

# Check if protoc-gen-go and protoc-gen-go-grpc are installed
check-protoc-tools:
	@if ! [ -x "$(PROTOC_GEN_GO)" ]; then \
		echo "Error: protoc-gen-go is not installed. Run: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; exit 1; \
	fi
	@if ! [ -x "$(PROTOC_GEN_GO_GRPC)" ]; then \
		echo "Error: protoc-gen-go-grpc is not installed. Run: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; exit 1; \
	fi

# Generate Protobuf files
proto: check-protoc-tools
	@$(PROTOC) \
		--go_out=$(PROTO_OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative \
		-I $(PROTO_SRC) $(PROTO_SRC)/*.proto
	@echo "✅ Protobuf files generated in $(PROTO_OUT)/"

tidy:
	@echo "🛠 Running go mod tidy..."
	@go mod tidy
	@echo "✅ go.mod and go.sum are updated!"

# Build the service
build: tidy
	@echo "🔨 Building $(SERVICE)..."
	@go build -o bin/$(SERVICE) $(CMD_DIR)/$(SERVICE)/main.go
	@echo "✅ Built: bin/$(SERVICE)"

# Run the service
run:
	@echo "🚀 Running $(filter-out $@,$(MAKECMDGOALS))..."
	@go run cmd/$(filter-out $@,$(MAKECMDGOALS))/main.go

# Clean generated files
clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(PROTO_OUT)/*.pb.go
	@rm -rf bin/
	@echo "✅ Cleaned!"

# Format code
fmt:
	@echo "🛠 Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted!"

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test ./... -v
	@echo "✅ Tests completed!"

# Help message
help:
	@echo "Available commands:"
	@echo "  make proto     - Generate Protobuf files"
	@echo "  make build     - Build the service binary"
	@echo "  make run       - Run the service"
	@echo "  make clean     - Remove compiled binaries and generated Protobuf files"
	@echo "  make fmt       - Format Go code"
	@echo "  make test      - Run tests"

