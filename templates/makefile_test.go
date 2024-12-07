package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
	"github.com/stretchr/testify/assert"
)

func Test_NewMakefileTemplate(t *testing.T) {
	expected := `
PROJECT_NAME=test

all: build

build:
	go build -o bin/$(PROJECT_NAME) cmd/http/main.go

run:
	./bin/$(PROJECT_NAME)

test:
	go test -v ./...

clean:
	rm -rf bin/$(PROJECT_NAME)
	rm -rf bin

.PHONY: build run clean
`
	result := templates.NewMakefileTemplate("test")
	if expected != result {
		assert.Equal(t, expected, result)
	}
}
