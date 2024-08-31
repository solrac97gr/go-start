package files

import (
	"os"

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

func (f *Files) CreateFilesStructure(projectName string, subAppName []string) {
	f.CreateFile(projectName+"/go.mod", "module "+projectName)

	for _, route := range subAppName {
		route = utils.Lowercase(route)

		f.CreateFile(projectName+"/internal/"+route+"/domain/"+route+".go", "")
		f.CreateFile(projectName+"/internal/"+route+"/domain/ports/"+route+"_repository.go", "")
		f.CreateFile(projectName+"/internal/"+route+"/application/"+route+"_service.go", "")
		f.CreateFile(projectName+"/internal/"+route+"/infrastructure/repository/"+route+"_repository.go", "")
		f.CreateFile(projectName+"/internal/"+route+"/infrastructure/handler/"+route+"_handler.go", "")
	}
}
