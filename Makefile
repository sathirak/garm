# Directory of makefile
PROJECT_ROOT = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

export GOBIN ?= $(PROJECT_ROOT)bin
export PATH := $(GOBIN):$(PATH)

MODULE_DIRS = .

build:
	@go build -o bin/garm

run: build
	@./bin/garm

dev:
	@clear
	@staticcheck .
	@goimports -w .
	@go vet
	@go build -o bin/garm
	@./bin/garm

lint:
	@staticcheck .
	@goimports -w .
	@go vet
	@go build -o bin/garm

.PHONY: tidy
tidy:
	@$(foreach dir,$(MODULE_DIRS), \
		(cd $(dir) && go mod tidy) &&) true

.PHONY: govulncheck
vulncheck:
	@govulncheck ./...

.PHONY: test
test:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go test -race ./...) &&) true