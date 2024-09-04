package templates

import "fmt"

var GoModTemplate = `module github.com/%s/%s

go %s

require (
    github.com/gofiber/fiber/v2 v2.17.0
)
`

func NewGoModTemplate(githubUser string, projectName string, goVersion string) string {
	return fmt.Sprintf(GoModTemplate, githubUser, projectName, goVersion)
}
