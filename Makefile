APP_NAME = inno_cart
DOCKER_COMPOSE_FILE = docker-compose.yml

test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests passed!"

build: test
	@echo "Building the Go binary..."
	@go build -o ./bin/${APP_NAME} ./cmd
	@echo "Build successful!"

docker-up: build
	@echo "Starting the Docker containers..."
	@docker-compose -f ${DOCKER_COMPOSE_FILE} up --build -d
	@echo "Docker containers are up and running!"

docker-down:
	@echo "Stopping the Docker containers..."
	@docker-compose -f ${DOCKER_COMPOSE_FILE} down
	@echo "Docker containers are stopped!"

clean:
	@echo "Cleaning up..."
	@rm -f ./bin/${APP_NAME}
	@rm -rf ./bin
	@echo "Cleaned up!"

deploy: docker-up

all: test build docker-up

.PHONY: test build docker-build docker-up docker-down clean deploy all
