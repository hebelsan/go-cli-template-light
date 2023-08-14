BIN_NAME=go-template-cli
BIN_PATH=${GOPATH}/bin
IMAGE_NAME=${BIN_NAME}

## Locally run the golang test.
test:
	go test ./...

## Build locally the go project.
build:
	@echo "building ${BIN_NAME}"
	@echo "GOPATH=${GOPATH}"
	go build -o ${BIN_PATH}/${BIN_NAME}

lint:
	golangci-lint --exclude-use-default=false run ./...

## Compile optimized for alpine linux.
docker.build:
	@echo "building image ${IMAGE_NAME}"
	docker build -t $(IMAGE_NAME):latest .

## generate markdown documentation
doc:
	go run scripts/gen_docs.go