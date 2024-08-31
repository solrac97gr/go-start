package templates

import "fmt"

var HandlersTemplate = `package handlers

import (
	"github.com/%s/%s/domain/ports"
)

type %sHandler struct {
	application ports.%sAplication
}

func New%sHandler() *%sHandler {
	return &%sHandler{}
}
`

func NewHandlersTemplate(githubUser, projectName, handlerName string) string {
	return fmt.Sprintf(HandlersTemplate, githubUser, projectName, handlerName, handlerName, handlerName, handlerName, handlerName)
}
