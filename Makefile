
PACKAGES = $(shell go list ./... | grep -v '/vendor/')
VERSION = $(shell git rev-parse --short HEAD)
COMMIT = $(shell git log -1 --format='%H')

BUILD_TAGS = netgo
# BUILD_FLAGS = -tags "${BUILD_TAGS}" -ldflags \
# 	"-X github.com/ibc-marketplace/version.VERSION=${VERSION} \
# 	-X github.com/ibc-marketplace/version.COMMIT=${COMMIT} \
# 	-X github.com/ibc-marketplace/version.BuildTags=${BUILD_TAGS} \
# 	-s -w"

all: build install

build: go.sum
	GO111MODULE=on go build  $(BUILD_FLAGS)  -o build/relayerd ./cmd/relayerd/
	GO111MODULE=on go build   $(BUILD_FLAGS) -o build/relayercli ./cmd/relayercli/

install: go.sum
	go install -tags "$(BUILD_FLAGS)" ./cmd/relayercli
	go install -tags "$(BUILD_FLAGS)" ./cmd/relayerd

test:
	@go test -cover $(PACKAGES)

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		go mod verify

.PHONY: all build install go.sum test