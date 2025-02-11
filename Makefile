# Project Variables
PROTO_DIR=api
SERVER_DIR=server

# Generate protobuf code using buf
generate:
	@echo "Generating protobuf files..."
	rm -rf pkg/api/* ; cd $(PROTO_DIR) ; buf dep update ; buf lint ; buf generate
	
