.PHONY: build run clean test

build:
	go build -o pharmacy-api main.go

run:
	go run main.go

clean:
	rm -f pharmacy-api

test:
	go test ./...

deps:
	go mod download

tidy:
	go mod tidy
