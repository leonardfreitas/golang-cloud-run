.PHONY: build test

build: 
	@echo "Building docker image $(IMAGE_NAME)..."
	docker build -t golang-image:latest .

test:
	go test ./... -v
