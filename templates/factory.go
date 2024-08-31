package templates

import (
	"fmt"

	"github.com/solrac97gr/go-start/utils"
)

var FactoryTemplate = `package factory

import (
	"github.com/%s/%s/internal/%s/application"
	"github.com/%s/%s/internal/%s/infrastructure/repository"
	"github.com/%s/%s/internal/%s/infrastructure/handler"
	"github.com/%s/%s/internal/%s/infrastructure/server"
)

func Build%sServer() *server.%sServer {
	repo, err := repository.New%sRepository()
	if err != nil {
		panic(err)
	}
	app, err := application.New%sApp(repo)
	if err != nil {
		panic(err)
	}
	handler, err := handler.New%sHandler(app)
	if err != nil {
		panic(err)
	}
	server, err := server.New%sServer(handler)
	if err != nil {
		panic(err)
	}
	return server
}
`

func NewFactoryTemplate(githubUser, projectName, subAppName string) string {
	subAppName = utils.Capitalize(subAppName)
	return fmt.Sprintf(
		FactoryTemplate,
		githubUser,
		projectName,
		utils.Lowercase(subAppName),
		githubUser,
		projectName,
		utils.Lowercase(subAppName),
		githubUser,
		projectName,
		utils.Lowercase(subAppName),
		githubUser,
		projectName,
		utils.Lowercase(subAppName),
		subAppName,
		subAppName,
		subAppName,
		subAppName,
		subAppName,
		subAppName,
	)
}
