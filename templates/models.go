package templates

import (
	"fmt"

	"github.com/solrac97gr/go-start/utils"
)

var ModelsTemplate = `package models

type %s struct {
}

func New%s() *%s {
	return &%s{}
}
`

func NewModelsTemplate(modelName string) string {
	modelName = utils.Capitalize(modelName)

	return fmt.Sprintf(
		ModelsTemplate,
		modelName,
		modelName,
		modelName,
		modelName,
	)
}
