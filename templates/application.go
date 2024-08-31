package templates

import (
	"fmt"

	"github.com/solrac97gr/go-start/utils"
)

var ApplicationTemplate = `package application

import (
	"github.com/%s/%s/internal/%s/domain/ports"
)

type %sApp struct {
	repository	ports.%sRepository
}

func New%sApp(repo ports.%sRepository) (*%sApp, error) {
	return &%sApp{
		repository: repo,
	}, nil
}
`

func NewApplicationTemplate(githubName, projectName, appName string) string {
	appName = utils.Capitalize(appName)
	lowerAppName := utils.Lowercase(appName)

	return fmt.Sprintf(
		ApplicationTemplate,
		githubName,
		projectName,
		lowerAppName,
		appName,
		appName,
		appName,
		appName,
		appName,
		appName,
	)
}
