package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/solrac97gr/go-start/files"
	"github.com/solrac97gr/go-start/folders"
)

func main() {
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

	fd := folders.NewFolders()
	err := fd.CreateFolderStructure(projectName, subAppsNamesList)
	if err != nil {
		fd.RemoveFolder(projectName)
		panic(err)
	}

	fl := files.NewFiles()
	majorMinor := strings.Split(strings.TrimPrefix(runtime.Version(), "go"), ".")
	goVersion := majorMinor[0] + "." + majorMinor[1]
	fl.CreateFilesStructure(githubUsername, projectName, goVersion, subAppsNamesList)
	fmt.Println("Project structure created successfully")

	After(projectName)
}

func After(projectName string) {
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
