run: build
	@./bin/mansys

build:
	@go build -o ./bin/mansys main.go

test:
	@go test -v ./...

.PHONY: run build test

all: run

help:
	@echo "run: run the mansys"
	@echo "build: go build -0 ./bin/mansys"
	@echo "test: go test -v"
	@echo "all:input make but with nothing will do,default run"


