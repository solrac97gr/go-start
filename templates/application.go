package templates

import (
	"fmt"
)

var ApplicationTemplate = `package application
type %sApp struct {
	repository	ports.%sRepository
}

func New%sApp(repo ports.%sRepository) *%sApp {
	return &%sApp{
		repository: repo,
	}
}
`

func NewApplicationTemplate(appName string) string {
	return fmt.Sprintf(ApplicationTemplate, appName, appName, appName, appName, appName, appName)
}

