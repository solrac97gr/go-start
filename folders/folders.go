package folders

import (
	"fmt"
	"os"
)

type FoldersService struct {
	projectFolders []string
	appFolders     []string
}

// NewFolderService create a new folder service
func NewFolderService() *FoldersService {
	projectFolders, appFolders := initializeFolders()
	return &FoldersService{
		projectFolders: projectFolders,
		appFolders:     appFolders,
	}
}

// CreateFolderStructure create the folder structure necessary for the project
func (f *FoldersService) CreateFolderStructure(projectName string, subAppName []string) error {
	// Create the project principal folder
	if err := f.createPrincipalFolders(projectName); err != nil {
		return err
	}

	// Create the subapp folders
	for _, route := range subAppName {
		if err := f.CreateAppsFolders(projectName, route); err != nil {
			return err
		}
	}
	return nil
}

// createPrincipalFolders create the folders of the root of the project
func (f *FoldersService) createPrincipalFolders(projectName string) error {
	// Create the project folders
	for _, folder := range f.projectFolders {
		if err := f.CreateFolder(projectName + folder); err != nil {
			return err
		}
	}
	return nil
}

// CreateAppsFolders create the folders of an app
func (f *FoldersService) CreateAppsFolders(projectName string, appName string) error {
	for _, folder := range f.appFolders {
		if err := f.CreateFolder(f.buildAppRoutePath(projectName, appName) + folder); err != nil {
			return err
		}
	}
	return nil
}

// buildAppRoutePath Build the route of the App
func (f *FoldersService) buildAppRoutePath(projectName string, route string) string {
	return fmt.Sprintf("%s/internal/%s", projectName, route)
}

// Create a folder if it does not exist
func (f *FoldersService) CreateFolder(route string) error {
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

// RemoveFolder delete folder
func (f *FoldersService) RemoveFolder(route string) error {
	if err := os.RemoveAll(route); err != nil {
		return err
	}
	return nil
}

func initializeFolders() (projectFolders []string, appFolders []string) {
	return []string{
			"/", // for create the container folder
			"/internal",
			"/cmd",
			"/cmd/http",
			"/pkg",
			"/pkg/factory",
			"/pkg/server",
			"/infrastructure",
			"/infrastructure/docker",
		},
		[]string{
			"/", // for create the container folder
			"/domain",
			"/domain/models",
			"/domain/ports",
			"/application",
			"/infrastructure",
			"/infrastructure/repository",
			"/infrastructure/handler",
			"/infrastructure/server",
		}
}
