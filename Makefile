ALL_PKGS := $(shell go list ./... | grep -v /vendor)

GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOLANGCI-LINT=golangci-lint

.PHONY: run test lint
# DON'T EDIT BELOW

all: test_ci

build:
	$(GOCMD) build -o ./release/gct-ta ./cmd/gct-ta/. 
test_ci:
	$(GOCMD) test -cover -race -v $(ALL_PKGS)

test_all:
	$(GOCMD) test -race -v $(ALL_PKGS)

lint:
	$(GOLANGCI-LINT) run -verbose
