APP_NAME = inno_cart
DOCKER_COMPOSE_FILE = docker-compose.yml

GO_FILES := $(shell find . -type f -name '*.go')

test:
	@echo "Running tests..."
	@go test ./...
	@echo "Tests passed!"

build: test
	@echo "Building the Go binary..."
	@go build -o $(APP_NAME) ./cmd
	@echo "Build successful!"


docker-up: build
	@echo "Starting the Docker containers..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "Docker containers are up and running!"

docker-down:
	@echo "Stopping the Docker containers..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down
	@echo "Docker containers are stopped!"

clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)
	@echo "Cleaned up!"

deploy: docker-up

all: test build docker-up

.PHONY: test build docker-build docker-up docker-down clean deploy all
