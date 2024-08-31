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
func (f *Folders) CreateFolder(route string) error {
	if err := os.Mkdir(route, 0755); err != nil {
		if os.IsExist(err) {
			fmt.Println("Folder already exists: ", route)
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (f *Folders) CreateFolderStructure(projectName string, subAppName []string) error {
	var err error
	err = f.CreateFolder(projectName + "/internal")
	if err != nil {
		return err
	}
	err = f.CreateFolder(projectName + "/cmd")
	if err != nil {
		return err
	}
	err = f.CreateFolder(projectName + "/cmd/http")
	if err != nil {
		return err
	}
	err = f.CreateFolder(projectName + "/configs")
	if err != nil {
		return err
	}
	err = f.CreateFolder(projectName + "/scripts")
	if err != nil {
		return err
	}
	err = f.CreateFolder(projectName + "/pkg")
	if err != nil {
		return err
	}
	err = f.CreateFolder(projectName + "/infrastructure")
	if err != nil {
		return err
	}
	err = f.CreateFolder(projectName + "/infrastructure/docker")
	if err != nil {
		return err

	}
	for _, route := range subAppName {
		var err error
		err = f.CreateFolder(projectName + "/internal/" + route)
		if err != nil {
			return err
		}
		err = f.CreateFolder(projectName + "/internal/" + route + "/domain")
		if err != nil {
			return err
		}
		err = f.CreateFolder(projectName + "/internal/" + route + "/domain/ports")
		if err != nil {
			return err
		}
		err = f.CreateFolder(projectName + "/internal/" + route + "/application")
		if err != nil {
			return err
		}
		err = f.CreateFolder(projectName + "/internal/" + route + "/infrastructure/repository")
		if err != nil {
			return err
		}
		err = f.CreateFolder(projectName + "/internal/" + route + "/infrastructure/handler")
		if err != nil {
			return err
		}
		err = f.CreateFolder(route)
		if err != nil {
			return err
		}
	}

	fmt.Println("Folder structure created successfully!")
	return nil
}
