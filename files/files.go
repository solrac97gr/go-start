package files

import (
	"fmt"
	"os"
	"sync"

	"github.com/solrac97gr/go-start/templates"
	"github.com/solrac97gr/go-start/utils"
)

// fileConfig helper struct
type fileConfig struct {
	path     string
	template func(githubName string, projectName string, appName string) string
}

// Template Generator Params type helper
// with this helper we can add new params without to modify all the functions in the map
type TemplateGeneratorParams struct {
	githubUser  string
	projectName string
	goVersion   string
	subAppName  []string
}

// Template Generator function type helper
type TemplateGenerator func(params *TemplateGeneratorParams) string

// AppFileNameGenerator function type helper
type AppFileNameGenerator func(projectName string, route string) string

// FilesService file service generator
type FilesService struct {
	projectFiles map[string]TemplateGenerator
	appFiles     map[string]fileConfig
}

// NewFilesService create a server for generate the files for the project with their content
func NewFilesService() *FilesService {
	return &FilesService{
		projectFiles: initializeProjectFiles(),
		appFiles:     initializeAppFiles(),
	}
}

// CreateFilesStructure creates the files of the project
func (f *FilesService) CreateFilesStructure(githubName string, projectName string, goVersion string, subAppName []string) error {
	templateGenerator := &TemplateGeneratorParams{
		githubUser:  githubName,
		projectName: projectName,
		goVersion:   goVersion,
		subAppName:  subAppName,
	}

	// Initialize WaitGroup
	var wg sync.WaitGroup
	errChan := make(chan error, 1) // Buffered to avoid blocking

	// Iterate through the map of project files and generate them
	for file, generateContent := range f.projectFiles {
		wg.Add(1)
		// 	Problem: Loop variables are reused, causing all goroutines to share the same variables.
		//	Solution: Capture the loop variables into new variables local to the current iteration.
		//	Result: Each goroutine has its own copy of the variables, ensuring correct behavior.
		file := file
		generateContent := generateContent

		go func(file string, generateContent func(*TemplateGeneratorParams) string) {
			defer wg.Done()
			content := generateContent(templateGenerator)
			if err := f.CreateFile(projectName+file, content); err != nil {
				errChan <- err
			}
		}(file, generateContent)
	}

	// Iterate through the array of sub apps and generate their files
	for _, appName := range subAppName {
		wg.Add(1)

		// Capture the current loop variable to avoid closure issues
		appName := utils.Lowercase(appName)

		go func(appName string) {
			defer wg.Done()
			if err := f.createApp(githubName, projectName, appName); err != nil {
				errChan <- err
			}
		}(appName)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Check for errors
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// createApp create the app files
func (f *FilesService) createApp(githubName, projectName, appName string) error {
	for name, file := range f.appFiles {
		var fullPath string
		fullPath = fmt.Sprintf("%s/"+file.path, projectName, appName, appName)
		if name == "factory" {
			fullPath = fmt.Sprintf("%s/"+file.path, projectName, appName)
		}
		if err := f.CreateFile(fullPath, file.template(githubName, projectName, appName)); err != nil {
			return fmt.Errorf("failed to create %s file: %w", name, err)
		}
	}
	return nil
}

// Create a file with content if it does not exist
func (f *FilesService) CreateFile(route string, content string) error {
	_, err := os.Create(route)
	if err != nil {
		return err
	}
	if err := os.WriteFile(route, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

func initializeProjectFiles() map[string]TemplateGenerator {
	return map[string]TemplateGenerator{
		"/go.mod": func(params *TemplateGeneratorParams) string {
			return templates.NewGoModTemplate(params.githubUser, params.projectName, params.goVersion)
		},
		"/cmd/http/main.go": func(params *TemplateGeneratorParams) string {
			return templates.NewMainTemplate(params.githubUser, params.projectName, params.subAppName)
		},
		"/pkg/server/super_server.go": func(params *TemplateGeneratorParams) string {
			return templates.NewSuperServerTemplate(params.projectName)
		},
		"/Makefile": func(params *TemplateGeneratorParams) string {
			return templates.NewMakefileTemplate(params.projectName)
		},
		"/infrastructure/docker/Dockerfile": func(params *TemplateGeneratorParams) string {
			return templates.NewDockerfileTemplate(params.goVersion)
		},
		"/docker-compose.yaml": func(params *TemplateGeneratorParams) string {
			return templates.NewDockerComposeTemplate()
		},
		"/.env.example": func(params *TemplateGeneratorParams) string {
			return templates.NewEnvTemplate()
		},
		"/.gitignore": func(params *TemplateGeneratorParams) string {
			return templates.NewGitIgnoreTemplate()
		},
	}
}

func initializeAppFiles() map[string]fileConfig {
	return map[string]fileConfig{
		"models": {
			path: "internal/%s/domain/models/%s.go",
			template: func(githubName, projectName, appName string) string {
				return templates.NewModelsTemplate(appName)
			},
		},
		"ports": {
			path: "internal/%s/domain/ports/%s.go",
			template: func(githubName, projectName, appName string) string {
				return templates.NewPortsTemplate(appName)
			},
		},
		"application": {
			path: "internal/%s/application/%s_application.go",
			template: func(githubName, projectName, appName string) string {
				return templates.NewApplicationTemplate(githubName, projectName, appName)
			},
		},
		"repository": {
			path: "internal/%s/infrastructure/repository/%s_repository.go",
			template: func(githubName, projectName, appName string) string {
				return templates.NewRepositoryTemplate(appName)
			},
		},
		"handler": {
			path: "internal/%s/infrastructure/handler/%s_handler.go",
			template: func(githubName, projectName, appName string) string {
				return templates.NewHandlersTemplate(githubName, projectName, appName)
			},
		},
		"factory": {
			path: "pkg/factory/%s_factory.go",
			template: func(githubName, projectName, appName string) string {
				return templates.NewFactoryTemplate(githubName, projectName, appName)
			},
		},
		"server": {
			path: "internal/%s/infrastructure/server/%s_server.go",
			template: func(githubName, projectName, appName string) string {
				return templates.NewServerTemplate(githubName, projectName, appName)
			},
		},
	}
}
