package files

import (
	"os"

	"github.com/solrac97gr/go-start/templates"
	"github.com/solrac97gr/go-start/utils"
)

type Files struct {
}

func NewFiles() *Files {
	return &Files{}
}

// Create a file with content if it does not exist
func (f *Files) CreateFile(route string, content string) {
	os.Create(route)
	os.WriteFile(route, []byte(content), 0644)
}

func (f *Files) CreateFilesStructure(githubName string, projectName string, goVersion string, subAppName []string) {
	f.CreateFile(projectName+"/go.mod", templates.NewGoModTemplate(githubName, projectName, goVersion))
	f.CreateFile((projectName + "/cmd/http/main.go"), templates.NewMainTemplate(githubName, projectName, subAppName))
	f.CreateFile(projectName+"/pkg/server/super_server.go", templates.NewSuperServerTemplate(projectName))
	f.CreateFile(projectName+"/Makefile", templates.NewMakefileTemplate(projectName))

	for _, route := range subAppName {
		route = utils.Lowercase(route)

		// Create models file
		f.CreateFile(
			projectName+"/internal/"+route+"/domain/models/"+route+".go",
			templates.NewModelsTemplate(route),
		)
		// Create Ports file
		f.CreateFile(
			projectName+"/internal/"+route+"/domain/ports/"+route+"_repository.go",
			templates.NewPortsTemplate(route),
		)
		// Create Application file
		f.CreateFile(
			projectName+"/internal/"+route+"/application/"+route+"_application.go",
			templates.NewApplicationTemplate(githubName, projectName, route),
		)
		// Create Repository file
		f.CreateFile(
			projectName+"/internal/"+route+"/infrastructure/repository/"+route+"_repository.go",
			templates.NewRepositoryTemplate(route),
		)
		// Create handler file
		f.CreateFile(
			projectName+"/internal/"+route+"/infrastructure/handler/"+route+"_handler.go",
			templates.NewHandlersTemplate(githubName, projectName, route),
		)
		// Create Factory file
		f.CreateFile(
			projectName+"/pkg/factory/"+route+"_factory.go",
			templates.NewFactoryTemplate(githubName, projectName, route),
		)
		// Create Server file
		f.CreateFile(
			projectName+"/internal/"+route+"/infrastructure/server/"+route+"_server.go",
			templates.NewServerTemplate(githubName, projectName, route),
		)
	}
}
