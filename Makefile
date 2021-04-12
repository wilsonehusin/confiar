GOROOT?=$(shell go env GOROOT)
GIT_TAG?=$(shell git describe --always --dirty --tags)
GIT_SHA?=$(shell git rev-parse --verify HEAD)
CONTAINER_REGISTRY?=ghcr.io/wilsonehusin
GOX_FLAGS?=-osarch="linux/amd64 linux/arm linux/arm64 darwin/amd64 darwin/arm64"
OUT_DIR?=_output
CGO_ENABLED?=0

GOTARGET=github.com/wilsonehusin/confiar
GOVERSION=$(shell go env GOVERSION)

REPO_ROOT=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
TARGET_IMAGES=$(shell docker image ls --all --format "{{.Repository}}:{{.Tag}}" | grep confiar)

VERSION_FLAG=-X=$(GOTARGET)/internal.Version=$(GIT_TAG)
GIT_SHA_FLAG=-X=$(GOTARGET)/internal.GitSHA=$(GIT_SHA)
COMPILER_FLAG=-X=$(GOTARGET)/internal.Go=$(GOVERSION)

BUILD_FLAGS=-ldflags '$(VERSION_FLAG) $(GIT_SHA_FLAG) $(COMPILER_FLAG)'
TEST_BUILD_FLAGS=-ldflags '$(VERSION_FLAG) $(GIT_SHA_FLAG) $(COMPILER_FLAG)'
DOCKER_TAGS=--tag $(CONTAINER_REGISTRY)/confiar:latest --tag $(CONTAINER_REGISTRY)/confiar:$(GIT_TAG)

.PHONY: all
all: build

.PHONY: testbuild
testbuild: GIT_TAG:=zz_test-$(GIT_TAG)
testbuild: BUILD_FLAGS:=$(TEST_BUILD_FLAGS)
testbuild: build

.PHONY: testmultibuild
testmultibuild: GIT_TAG:=zz_test-$(GIT_TAG)
testmultibuild: BUILD_FLAGS:=-ldflags '$(VERSION_FLAG) $(GIT_SHA_FLAG) $(COMPILER_FLAG)'
testmultibuild: multibuild

.PHONY: testcontainer
testcontainer: testbuild
	docker build $(DOCKER_TAGS) .

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(OUT_DIR)/ $(BUILD_FLAGS) .

.PHONY: multibuild
multibuild:
	CGO_ENABLED=$(CGO_ENABLED) gox -output="$(OUT_DIR)/{{.Dir}}_{{.OS}}_{{.Arch}}" $(GOX_FLAGS) $(BUILD_FLAGS)

.PHONY: container
container: build
	docker build $(DOCKER_TAGS) .

.PHONY: release
release: container
	docker push $(DOCKER_TAGS)

.PHONY: clean
clean:
	rm -rf $(REPO_ROOT)/$(OUT_DIR)
	@docker image rm $(confiar_IMAGES) 2>/dev/null || true
