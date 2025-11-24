#!/usr/bin/make -f

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)
CONTRACTS_DIR = ./x/evm/contracts

# don't override user values of COMMIT and VERSION
ifeq (,$(COMMIT))
  COMMIT := $(shell git log -1 --format='%H')
endif

ifeq (,$(VERSION))
  VERSION := $(shell git describe --tags)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

TM_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
	ifeq ($(OS),Windows_NT)
		GCCEXE = $(shell where gcc.exe 2> NUL)
		ifeq ($(GCCEXE),)
			$(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
		else
			build_tags += ledger
		endif
	else
		UNAME_S = $(shell uname -s)
		ifeq ($(UNAME_S),OpenBSD)
			$(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
		else
			GCC = $(shell command -v gcc 2> /dev/null)
			ifeq ($(GCC),)
				$(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
			else
				build_tags += ledger
			endif
		endif
	endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc
endif
# handle rocksdb
define ROCKSDB_INSTRUCTIONS

################################################################
RocksDB support requires the RocksDB shared library and headers.
macOS (Homebrew):
  brew install rocksdb
  export CGO_CFLAGS="-I/usr/local/opt/rocksdb/include"
  export CGO_LDFLAGS="-L/usr/local/opt/rocksdb/lib"
See https://github.com/rockset/rocksdb-cloud/blob/master/INSTALL.md for custom setups.
################################################################

endef

ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
  $(info $(ROCKSDB_INSTRUCTIONS))
  CGO_ENABLED ?= 1
  build_tags += rocksdb grocksdb_clean_link
endif
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += boltdb
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=minievm \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=minitiad \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
			-X github.com/cometbft/cometbft/version.TMCoreSemVer=$(TM_VERSION)

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# The below include contains the tools and runsim targets.
include contrib/devtools/Makefile

all: tools install lint test

build: go.sum
ifeq ($(OS),Windows_NT)
	exit 1
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/minitiad ./cmd/minitiad
endif

build-linux:
	mkdir -p $(BUILDDIR)
	docker build --no-cache --tag initia/minievm ./
	docker create --name temp initia/minievm:latest --env VERSION=$(VERSION)
	docker cp temp:/usr/local/bin/minitiad $(BUILDDIR)/
	docker rm temp

install: go.sum 
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/minitiad

update-swagger-docs: statik
	$(BINDIR)/statik -src=client/docs/swagger-ui -dest=client/docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
        exit 1;\
    else \
        echo "\033[92mSwagger docs are in sync\033[0m";\
    fi

.PHONY: build build-linux install update-swagger-docs

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

proto-swagger-gen:
	@echo "Generating Swagger files"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh
	$(MAKE) update-swagger-docs

proto-pulsar-gen:
	@echo "Generating Dep-Inj Protobuf files"
	@$(protoImage) sh ./scripts/protocgen-pulsar.sh

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec buf format {} -w \;

proto-lint:
	@$(protoImage) buf lint --error-format=json ./proto

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main

.PHONY: proto-all proto-gen proto-swagger-gen proto-pulsar-gen proto-format proto-lint proto-check-breaking


###############################################################################
###                           Tests 
###############################################################################

test: contracts-gen test-unit test-integration

test-all: contracts-gen test-unit test-race test-cover test-integration

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock test' ./...

test-integration:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock test' ./integration-tests/...

test-race:
	@VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock test' ./...

test-cover:
	@go test -mod=readonly -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock test' ./...

contracts-gen: $(CONTRACTS_DIR)/*
	@bash ./scripts/contractsgen.sh

benchmark:
	@go test -timeout 20m -mod=readonly -bench=. ./... 

fuzz:
	@go test --timeout 2m -mod=readonly -fuzz=Fuzz ./x/evm/keeper

.PHONY: test test-all test-cover test-unit test-race benchmark contracts-gen

###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	golangci-lint run --timeout=15m --tests=false

lint-fix:
	golangci-lint run --fix --timeout=15m --tests=false

.PHONY: lint lint-fix

###############################################################################
###                                Testnet                                  ###
###############################################################################

testnet-initialize:
	sh ./scripts/testnet.sh
.PHONY: testnet-initialize
