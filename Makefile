build:
	@go build -o bin/garm

run: build
	@./bin/garm

dev:
	@clear
	@go build -o bin/garm
	@./bin/garm