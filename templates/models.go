package templates

import "fmt"

var ModelsTemplate = `package models

type %s struct {
}

func New%s() *%s {
	return &%s{}
}
`

func NewModelsTemplate(modelName string) string {
	return fmt.Sprintf(ModelsTemplate, modelName, modelName, modelName, modelName)
}
