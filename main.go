package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/solrac97gr/go-start/files"
	"github.com/solrac97gr/go-start/folders"
	"github.com/solrac97gr/go-start/generator"
	"github.com/solrac97gr/go-start/utils"
)

var option *string

func init() {
	option = flag.String("flow", "default", "select the generator you wanna run")
}

func main() {
	flag.Parse()
	fd := folders.NewFolderService()
	fl := files.NewFilesService()
	gn := generator.NewGeneratorService(fl, fd)
	flow(*option, fd, fl, gn)
}

func flow(flow string, fd *folders.FoldersService, fl *files.FilesService, gn *generator.GeneratorService) {
	switch flow {
	case "default":
		defaultFlow(fd, fl)
	case "app":
		addAppFlow(gn)
	}
}

func defaultFlow(fd *folders.FoldersService, fl *files.FilesService) {
	fmt.Println("Welcome to the project generator")

	var projectName string
	fmt.Println("Enter the project name 📁: ")
	fmt.Scanln(&projectName)
	if projectName == "" {
		panic("Project name cannot be empty")
	}

	var githubUsername string
	fmt.Println("Enter the github username 🖥️: ")
	fmt.Scanln(&githubUsername)
	if githubUsername == "" {
		panic("Github username cannot be empty")
	}

	// Get the list of subapps
	var subAppsNames string
	fmt.Println("Enter the subapps names separated by comma 📦: ")
	fmt.Scanln(&subAppsNames)
	if subAppsNames == "" {
		panic("Subapps names cannot be empty")
	}
	// remove empty spaces from the subAppsNames
	subAppsNames = strings.TrimSpace(subAppsNames)
	// Get the list of subAppsNames
	subAppsNamesList := strings.Split(subAppsNames, ",")
	// Print the project name and github Username
	fmt.Println("Project Name: ", projectName)
	fmt.Println("Github Username: ", githubUsername)
	// Print the list of subAppsNamesList
	fmt.Println("SubApps Names: ", subAppsNamesList)

	if err := fd.CreateFolderStructure(projectName, subAppsNamesList); err != nil {
		fd.RemoveFolder(projectName)
		panic(err)
	}

	majorMinor := strings.Split(strings.TrimPrefix(runtime.Version(), "go"), ".")
	goVersion := majorMinor[0] + "." + majorMinor[1]
	if err := fl.CreateFilesStructure(githubUsername, projectName, goVersion, subAppsNamesList); err != nil {
		panic(err)
	}
	fmt.Println("Project structure created successfully")
	// Execute the go mod download
	os.Chdir(projectName)
	fmt.Println("Downloading dependencies 📥")
	cmdGoModDownload := exec.Command("go", "mod", "download")
	cmdGoModTidy := exec.Command("go", "mod", "tidy")
	cmdGitInit := exec.Command("git", "init")

	cmdGoModDownload.Stdout = os.Stdout
	cmdGoModTidy.Stdout = os.Stdout
	cmdGitInit.Stdout = os.Stdout

	cmdGoModDownload.Stderr = os.Stderr
	cmdGoModTidy.Stderr = os.Stderr
	cmdGitInit.Stderr = os.Stderr

	if err := cmdGoModDownload.Run(); err != nil {
		panic(err)
	}
	if err := cmdGoModTidy.Run(); err != nil {
		panic(err)
	}
	fmt.Println("Initializing Repository 🛸")
	if err := cmdGitInit.Run(); err != nil {
		panic(err)
	}
}

func addAppFlow(gn *generator.GeneratorService) {
	var appName string
	fmt.Println("Enter the app name 📁: ")
	fmt.Scanln(&appName)
	if appName == "" {
		panic("app name cannot be empty")
	}
	if err := gn.GenerateNewApp(utils.Lowercase(appName)); err != nil {
		panic(err)
	}
}
