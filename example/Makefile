APP_NAME = code-sdk-go-example
APP_VERSION = 1.0
DOCKER_IMAGE = $(APP_NAME):$(APP_VERSION)

# Build the Go application directly
build:
	go build -o build/$(APP_NAME)

# Build the Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
docker-run:
	docker run -p 3000:3000 $(DOCKER_IMAGE)

# Clean up
clean:
	rm -rf build/

# Build and run the example
example: docker-build docker-run

.PHONY: build docker-build docker-run clean example
