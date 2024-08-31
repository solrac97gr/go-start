package templates

import "fmt"

var GoModTemplate = `module %s

go %s
`

func NewGoModTemplate(projectName string, goVersion string) string {
	return fmt.Sprintf(GoModTemplate, projectName, goVersion)
}
