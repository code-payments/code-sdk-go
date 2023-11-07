DOCKER_IMAGE = code-sdk-go:latest

# Build the Docker image
docker-build:
	@docker build -t $(DOCKER_IMAGE) .

# Run tests in Docker image
docker-test: docker-build
	@docker run --rm -v $(PWD):/sdk $(DOCKER_IMAGE) go test -cover ./...

# Run tests
test:
	@go test -cover ./...

.PHONY: docker-build docker-test test
