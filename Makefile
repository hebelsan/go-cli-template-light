BIN_NAME=go-cli-template
BIN_PATH=${GOPATH}/bin
IMAGE_NAME=${BIN_NAME}

## Locally run the golang test.
test:
	golangci-lint run ./...
	go test ./...

## Build locally the go project.
build:
	@echo "building ${BIN_NAME}"
	@echo "GOPATH=${GOPATH}"
	go build -o ${BIN_PATH}/${BIN_NAME}

## Compile optimized for alpine linux.
docker.build:
	@echo "building image ${IMAGE_NAME}"
	docker build -t $(IMAGE_NAME):latest .

## generate markdown documentation
doc:
	go run scripts/gen_docs.go