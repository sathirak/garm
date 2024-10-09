build:
	@go build -o bin/garm

run: build
	@./bin/garm

dev:
	@clear
	@goimports -w .
	@go vet
	@go build -o bin/garm
	@./bin/garm

lint:
	@goimports -w .
	@go vet
	@go build -o bin/garm