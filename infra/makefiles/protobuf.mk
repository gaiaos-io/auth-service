# Versions
PROTOC_GEN_GO_VERSION := v1.36.11
PROTOC_GEN_GO_GRPC_VERSION := v1.6.0

# Paths
PROTO_BASE_DIR := proto
PROTO_VERSION := v1
PROTO_DIR := $(PROTO_BASE_DIR)/$(PROTO_VERSION)
PROTO_FILE := $(PROTO_DIR)/auth.proto

.PHONY: proto-install proto-gen proto-clean

# Install protoc plugins
proto-install:
	@echo "Installing protoc plugins..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

# Generate Go code
proto-gen: proto-install
	@echo "Generating Go code from proto files..."
	protoc \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)

# Clean generated files
proto-clean:
	@echo "Cleaning generated files..."
	rm -f $(PROTO_DIR)/*_pb.go $(PROTO_DIR)/*_grpc.pb.go

