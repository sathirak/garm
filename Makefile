build:
	@go build -o bin/garm

run: build
	@./bin/garm

dev:
	@go run main.go