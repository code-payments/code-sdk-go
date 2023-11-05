DOCKER_IMAGE = code-sdk-go:latest
PROTOC = docker run --rm -v $(PWD):/sdk $(DOCKER_IMAGE) protoc
PROTO_DIR = ./proto
PROTO_OUT_DIR = ./genproto

# Build the Docker image
docker-build:
	@docker build -t $(DOCKER_IMAGE) .

# Run tests in Docker image
docker-test: docker-build
	@docker run --rm -v $(PWD):/sdk $(DOCKER_IMAGE) go test -cover ./...

# Run tests
test:
	@go test -cover ./...

# Generate Go code from Protobuf files in Docker image
generate: docker-build
	$(PROTOC) $(PROTO_DIR)/messages.proto --go_out=$(PROTO_OUT_DIR)

# Clean generated files
clean:
	@rm -r $(PROTO_OUT_DIR)/*

.PHONY: docker-build docker-test test generate clean
