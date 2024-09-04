
PROJECT_NAME=go-start

all: build

build:
	go build -o bin/$(PROJECT_NAME) main.go

run:
	./bin/$(PROJECT_NAME)

test:
	go test -v ./...

clean:
	rm -rf bin/$(PROJECT_NAME)
	rm -rf bin

.PHONY: build run clean
