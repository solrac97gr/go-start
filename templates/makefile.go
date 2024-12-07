package templates

import (
	"fmt"

	"github.com/solrac97gr/go-start/utils"
)

var MakefileTemplate = `
PROJECT_NAME=%s

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
`

func NewMakefileTemplate(projectName string) string {
	lProjectName := utils.Lowercase(projectName)
	return fmt.Sprintf(MakefileTemplate, lProjectName)
}
