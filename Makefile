# Directory of
SERVICE = garm

PROJECT_ROOT = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

export GOBIN ?= $(PROJECT_ROOT)/bin
export PATH := $(GOBIN):$(PATH)

MODULE_DIRS = .

build:
	@go build -o bin/$(SERVICE)

run: build
	@./bin/$(SERVICE)

dev:
	@clear
	@staticcheck .
	@goimports -w .
	@go vet .
	@go build -o bin/$(SERVICE) .
	@./bin/$(SERVICE)

check:
	@clear
	@govulncheck
	@make lint-ci
	@staticcheck .
	@goimports -w .
	@go vet .

.PHONY: lint-ci
lint-ci: golangci-lint tidy-lint

.PHONY: golangci-lint
golangci-lint:
	@$(foreach mod,$(MODULE_DIRS), \
		(cd $(mod) && \
		echo "[lint] golangci-lint: $(mod)" && \
		golangci-lint run --path-prefix $(mod) ./...) &&) true

lint:
	@staticcheck .
	@goimports -w .
	@go vet
	@go build -o bin/$(SERVICE)

.PHONY: tidy
tidy:
	@$(foreach dir,$(MODULE_DIRS), \
		(cd $(dir) && go mod tidy) &&) true

.PHONY: tidy-lint
tidy-lint:
	@$(foreach mod,$(MODULE_DIRS), \
	(cd $(mod) && \
	echo "[lint] tidy: $(mod)" && \
	go mod tidy && \
	git diff --exit-code -- go.mod go.sum)) || true


.PHONY: govulncheck
vulncheck:
	@govulncheck ./...

.PHONY: test
test:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go test -race ./...) &&) true
