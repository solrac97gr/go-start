package templates

import (
	"fmt"

	"github.com/solrac97gr/go-start/utils"
)

var HandlersTemplate = `package handlers

import (
	"github.com/%s/%s/internal/%s/domain/ports"
)

type %sHandler struct {
	application ports.%sAplication
}

func New%sHandler() *%sHandler {
	return &%sHandler{}
}
`

func NewHandlersTemplate(githubUser, projectName, handlerName string) string {
	handlerName = utils.Capitalize(handlerName)
	var lowerHandlerName = utils.Lowercase(handlerName)

	return fmt.Sprintf(
		HandlersTemplate,
		githubUser,
		projectName,
		lowerHandlerName,
		handlerName,
		handlerName,
		handlerName,
		handlerName,
		handlerName,
	)
}
