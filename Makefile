# Variables
APP_NAME := sus-gopher-facts
IMAGE_TAG := latest
REPO_URL := diegoalveslv/$(APP_NAME)

# Show help
help:
	@echo "Available commands:"
	@echo "  make help         - Show this help message"
	@echo "  make go-build     - Build the Go binary"
	@echo "  make go-clean     - Clean up the Go binary"
	@echo "  make docker-build - Build the Docker image"
	@echo "  make docker-run   - Run the Docker container"
	@echo "  make docker-tag   - Tag the Docker image for the remote repository"
	@echo "  make docker-push  - Push the Docker image to the remote repository"
	@echo "  make docker-delete- Delete the Docker image"

# Build the Go binary
go-build:
	CGO_ENABLED=0 go build -o build/$(APP_NAME)

# Clean up the Go binary
go-clean:
	go clean
	rm -f $(APP_NAME)

# Build the Docker image
docker-build: go-build
	docker build -t $(APP_NAME):$(IMAGE_TAG) .

# Run the Docker container
docker-run:
	docker run -p 8080:8080 $(APP_NAME):$(IMAGE_TAG)

# Tag the Docker image for the remote repository
docker-tag:
	docker tag $(APP_NAME):$(IMAGE_TAG) $(REPO_URL):$(IMAGE_TAG)

# Push the Docker image to the remote repository
docker-push: docker-tag
	docker push $(REPO_URL):$(IMAGE_TAG)

# Delete the Docker image
docker-delete:
	docker rmi $(APP_NAME):$(IMAGE_TAG) $(REPO_URL):$(IMAGE_TAG)