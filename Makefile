.PHONY: build test

build: 
	@echo "Building docker image $(IMAGE_NAME)..."
	docker build -t $(golang-image:latest) -f Dockerfile.prod .

test:
	go test ./... -v
