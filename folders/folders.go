package folders

import (
	"fmt"
	"os"
)

type Folders struct {
}

func NewFolders() *Folders {
	return &Folders{}
}

// Create a folder if it does not exist
func (f *Folders) CreateFolder(route string) {
	os.Mkdir(route, 0755)
}

func (f *Folders) CreateFolderStructure(projectName string, subAppName []string) {
	f.CreateFolder(projectName + "/internal")
	f.CreateFolder(projectName + "/cmd")
	f.CreateFolder(projectName + "/cmd/http")
	f.CreateFolder(projectName + "/configs")
	f.CreateFolder(projectName + "/scripts")
	f.CreateFolder(projectName + "/pkg")
	f.CreateFolder(projectName + "/infrastructure")
	f.CreateFolder(projectName + "/infrastructure/docker")

	for _, route := range subAppName {
		f.CreateFolder(projectName + "/internal/" + route)
		f.CreateFolder(projectName + "/internal/" + route + "/domain")
		f.CreateFolder(projectName + "/internal/" + route + "/domain/ports")
		f.CreateFolder(projectName + "/internal/" + route + "/application")
		f.CreateFolder(projectName + "/internal/" + route + "/infrastructure/repository")
		f.CreateFolder(projectName + "/internal/" + route + "/infrastructure/handler")
		f.CreateFolder(route)
	}

	fmt.Println("Folder structure created successfully!")
}
