.PHONY: all build run test lint tidy

tidy:
	go mod tidy

build: tidy
	go build -o bin/market-data main.go

run: build
	./bin/market-data api

test: tidy
	go test -v ./...

lint:
	go vet ./...
