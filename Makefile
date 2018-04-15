PWD=$(shell pwd)
BINARY=server
NAME=chat

GOPAHT_PATH=GOPATH=${PWD}/vendor:${PWD}
GOLANG_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
FLAGS=-installsuffix cgo -ldflags "-w -extld ld -extldflags -static -a
GET_GB=go get github.com/constabulary/gb/...

DOCKER_BASE_IMAGE=alpine3.3

VERSION ?= $(DOCKER_BASE_IMAGE)-$(shell git rev-parse --short HEAD)

IMAGE=${NAME}:${VERSION}

run: build
	bin/${BINARY}

static: check_vendor clean
	${GOPAHT_PATH} ${GOLANG_ENV} go build -o bin/${BINARY} ${BINARY}

build: check_vendor clean
	gb build -ldflags "$(BUILD_FLAGS)" server

docker:
	docker build -t ${IMAGE} .

check_vendor: check_gb_tool
	if [ ! -d "vendor/src" ]; then gb vendor restore; fi

check_gb_tool:
	if ! which gb > /dev/null; then echo "Need install gb. Run: go get github.com/constabulary/gb/..."; exit 1; fi

# Cleans our project: deletes binaries
clean:
	rm -f bin/*
