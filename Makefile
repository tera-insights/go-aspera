# Copyright 2025 Tera Insights, LLC. All rights reserved.

VERSION := $(shell grep 'current_version =' ./cli/.bumpversion.toml | sed 's/^[[:space:]]*current_version = //'| tr -d '"')
VERSION_WO_DASHES = $(subst -,.,$(VERSION))

GOLANG_PKG := github.com/tera-insights/go-aspera
GOTAGS_TEST := -gcflags="all=-N -l" -tags no_openssl
GOTAGS := -tags no_openssl
GO_LDFLAGS = -ldflags '-X "main.Version=$(VERSION)" -X "main.BuildDate=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")" -X "main.GitHash=$(shell git rev-parse HEAD)" $(GO_EXTRA_LDFLAGS)'


## Building
clean:
	rm -rf bin

bin/: clean
	mkdir -p bin/

bin/debug/: clean
	mkdir -p bin/debug/

bin/go-aspera: bin/
	go build $(GO_LDFLAGS) $(GOTAGS) -o bin/ti-ascli ./cli

bin/debug/go-aspera: bin/debug/
	go build $(GO_LDFLAGS) $(GOTAGS_TEST) -o bin/debug/ti-ascli ./cli

build: bin/go-aspera

build_debug: bin/debug/go-aspera

test:
	go test ./... -cover

test-coverage:
	go test ./... -coverprofile=coverage.out -v
	go tool cover -html=coverage.out