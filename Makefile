# Project Variables
PROTO_DIR=api
SERVER_DIR=server

# Generate protobuf code using buf
generate:
	@echo "Generating protobuf files..."
	cd $(PROTO_DIR) && buf generate
