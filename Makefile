#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
ifeq ($(DETECTED_OS),)
  ifeq ($(OS),Windows_NT)
	  DETECTED_OS := windows
  else
	  UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),Darwin)
	    DETECTED_OS := mac
	  else
	    DETECTED_OS := linux
	  endif
  endif
endif
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

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=eurx \
		  -X github.com/cosmos/cosmos-sdk/version.ServerName=eurxd \
		  -X github.com/cosmos/cosmos-sdk/version.ClientName=eurxcli \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'


all: install

build: go.sum
ifeq ($(OS), Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/$(DETECTED_OS)/eurxd.exe ./cmd/eurxd
	go build -mod=readonly $(BUILD_FLAGS) -o build/$(DETECTED_OS)/eurxcli.exe ./cmd/eurxcli
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/$(DETECTED_OS)/eurxd ./cmd/eurxd
	go build -mod=readonly $(BUILD_FLAGS) -o build/$(DETECTED_OS)/eurxcli ./cmd/eurxcli
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 DETECTED_OS=linux $(MAKE) build

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/eurxd
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/eurxcli

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download
PHONY: go-mod-cache

go.sum: go.mod
	@echo "--> Ensuring dependencies have not been modified"
	@go mod verify

clean:
	rm -rf build/

########################################
### Linting

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify
.PHONY: lint

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/tendermint
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/cosmos/cosmos-sdk
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/lcnem/eurx
.PHONY: format

###############################################################################
###                                Localnet                                 ###
###############################################################################

build-docker-local-eurx:
	@$(MAKE) -C networks/local

# Run a 4-node testnet locally
localnet-start: build-linux localnet-stop
	@if ! [ -f build/node0/eurxd/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/eurxd:Z lcnem/eurxnode testnet --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test ; fi
	docker-compose up -d

localnet-stop:
	docker-compose down

########################################
### Testing

# TODO tidy up cli tests to use same -Enable flag as simulations, or the other way round
# TODO -mod=readonly ?
# build dependency needed for cli tests
test-all: build
	# basic app tests
	@go test ./app -v
	# basic simulation (seed "4" happens to not unbond all validators before reaching 100 blocks)
	@go test ./app -run TestFullAppSimulation        -Enabled -Commit -NumBlocks=100 -BlockSize=200 -Seed 4 -v -timeout 24h
	# other sim tests
	@go test ./app -run TestAppImportExport          -Enabled -Commit -NumBlocks=100 -BlockSize=200 -Seed 4 -v -timeout 24h
	@go test ./app -run TestAppSimulationAfterImport -Enabled -Commit -NumBlocks=100 -BlockSize=200 -Seed 4 -v -timeout 24h
	@# AppStateDeterminism does not use Seed flag
	@go test ./app -run TestAppStateDeterminism      -Enabled -Commit -NumBlocks=100 -BlockSize=200 -Seed 4 -v -timeout 24h

# run module tests and short simulations
test-basic: test
	@go test ./app -run TestFullAppSimulation        -Enabled -Commit -NumBlocks=5 -BlockSize=200 -Seed 4 -v -timeout 2m
	# other sim tests
	@go test ./app -run TestAppImportExport          -Enabled -Commit -NumBlocks=5 -BlockSize=200 -Seed 4 -v -timeout 2m
	@go test ./app -run TestAppSimulationAfterImport -Enabled -Commit -NumBlocks=5 -BlockSize=200 -Seed 4 -v -timeout 2m
	@# AppStateDeterminism does not use Seed flag
	@go test ./app -run TestAppStateDeterminism      -Enabled -Commit -NumBlocks=5 -BlockSize=200 -Seed 4 -v -timeout 2m

test:
	@go test ./...

test-rest:
	rest_test/./run_all_tests_from_make.sh

# Run cli integration tests
# `-p 4` to use 4 cores, `-tags cli_test` to tell go not to ignore the cli package
# These tests use the `eurxd` or `eurxcli` binaries in the build dir, or in `$BUILDDIR` if that env var is set.
test-cli: build
	@go test ./cli_test -tags cli_test -v -p 4

# Kick start lots of sims on an AWS cluster.
# This submits an AWS Batch job to run a lot of sims, each within a docker image. Results are uploaded to S3
start-remote-sims:
	# build the image used for running sims in, and tag it
	docker build -f simulations/Dockerfile -t lcnem/eurx-sim:master .
	# push that image to the hub
	docker push lcnem/eurx-sim:master
	# submit an array job on AWS Batch, using 1000 seeds, spot instances
	aws batch submit-job \
		-—job-name "master-$(VERSION)" \
		-—job-queue “simulation-1-queue-spot" \
		-—array-properties size=1000 \
		-—job-definition eurx-sim-master \
		-—container-override environment=[{SIM_NAME=master-$(VERSION)}]

.PHONY: all build-linux install clean build test test-cli test-all test-rest test-basic start-remote-sims

########################################
### Documentation

# Start docs site at localhost:8080
docs-develop:
	@cd docs && \
	npm install && \
	npm run serve

# Build the site into docs/.vuepress/dist
docs-build:
	@cd docs && \
	npm install && \
	npm run build