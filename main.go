package main

import (
	"fmt"
	"os"
	"os/exec"
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
	fl.CreateFilesStructure(githubUsername, projectName, "1.19", subAppsNamesList)
	fmt.Println("Project structure created successfully")

	// Execute the go mod download
	os.Chdir(projectName)
	fmt.Println("Downloading dependencies 📥")
	cmd := exec.Command("go", "mod", "download")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

}
