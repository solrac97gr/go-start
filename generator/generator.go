package generator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/solrac97gr/go-start/files"
	"github.com/solrac97gr/go-start/folders"
	"github.com/solrac97gr/go-start/utils"
)

type GeneratorService struct {
	fileService   *files.FilesService
	folderService *folders.FoldersService
}

func NewGeneratorService(fileService *files.FilesService, folderService *folders.FoldersService) *GeneratorService {
	return &GeneratorService{
		fileService:   fileService,
		folderService: folderService,
	}
}

// GenerateNewApp generate new app for the project
func (g *GeneratorService) GenerateNewApp(appName string) error {
	projectName, err := g.GetCurrentProjectName()
	if err != nil {
		return err
	}
	githubUserName, err := g.GetGithubUserName()
	if err != nil {
		return err
	}
	// Save the current working directory
	originalPath, err := os.Getwd()
	if err != nil {
		return err
	}

	// Navigate one directory up
	parentPath := filepath.Dir(originalPath)
	if err := os.Chdir(parentPath); err != nil {
		return err
	}

	// Ensure we navigate back to the original directory
	defer func() {
		_ = os.Chdir(originalPath) // Return to the original path, ignore error
	}()

	if err := g.folderService.CreateAppsFolders(projectName, appName); err != nil {
		return err
	}
	if err := g.fileService.CreateApp(githubUserName, projectName, appName); err != nil {
		return err
	}
	if err := g.integrateAppToMain(projectName, appName); err != nil {
		return err
	}
	return nil
}

func (g *GeneratorService) integrateAppToMain(projectName, appName string) error {
	sectionComment := "Server()"
	newServerLine := fmt.Sprintf("	%s := factory.Build%sServer()", appName, utils.Capitalize(appName))
	subAppsDeclaration := "subApps := []sserver.SubApp{"
	newSubAppLine := fmt.Sprintf("		%s,", appName)

	// Get the current path
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	// Construct the path to the main file
	mainPath := filepath.Join(currentPath, projectName, "cmd", "http", "main.go")

	// Read the main file
	content, err := os.ReadFile(mainPath)
	if err != nil {
		return err
	}

	// Split the content into lines
	lines := strings.Split(string(content), "\n")

	// Process the lines to insert new lines
	updatedLines := []string{}
	insertServerLine := false
	insertSubAppLine := false
	for _, line := range lines {
		updatedLines = append(updatedLines, line)

		// Insert newServerLine after the section comment
		if strings.Contains(line, sectionComment) && !insertServerLine {
			updatedLines = append(updatedLines, newServerLine)
			insertServerLine = true
		}

		// Insert newSubAppLine after the subApps declaration
		if strings.Contains(line, subAppsDeclaration) && !insertSubAppLine {
			updatedLines = append(updatedLines, newSubAppLine)
			insertSubAppLine = true
		}
	}

	// Write the updated content back to the main file
	err = os.WriteFile(mainPath, []byte(strings.Join(updatedLines, "\n")), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (g *GeneratorService) GetCurrentProjectName() (string, error) {
	// Get the current path
	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	projectName := strings.Split(currentPath, "/")
	return projectName[len(projectName)-1], nil
}

func (g *GeneratorService) GetGithubUserName() (string, error) {
	// Get the current path
	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Construct the path to the go.mod file
	goModulePath := filepath.Join(currentPath, "go.mod")

	// Open the file
	file, err := os.Open(goModulePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Use bufio.Scanner to read the first line
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		// Return the first line
		line := scanner.Text()
		return strings.Split(strings.Split(strings.TrimSpace(line), " ")[1], "/")[1], nil
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", nil // Return an empty string if the file is empty
}
