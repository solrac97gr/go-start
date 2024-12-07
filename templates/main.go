package templates

import (
	"fmt"
	"strings"

	"github.com/solrac97gr/go-start/utils"
)

var MainTemplate = `package main

import (
	"log"
	sserver "github.com/%s/%s/pkg/server"
	"github.com/%s/%s/pkg/factory"
)

func main() {
	// Sub apps initialize
%s
	// Build Super server
%s

	srv := sserver.NewServer(subApps)
	if err:=srv.Initialize("8080"); err != nil {
		log.Panic(err)
	}
}
`

func BuildServerSection(apps []string) string {
	var builder strings.Builder
	for _, app := range apps {
		cApp := utils.Capitalize(app)
		builder.WriteString(fmt.Sprintf("	%s:=factory.Build%sServer()\n", app, cApp))
	}
	return builder.String()
}

func BuildSuperServerSection(apps []string) string {
	var builder strings.Builder
	builder.WriteString("	subApps := []sserver.SubApp{\n")
	for _, app := range apps {
		builder.WriteString(fmt.Sprintf("		%s,\n", app))
	}
	builder.WriteString("	}")
	return builder.String()
}

func NewMainTemplate(githubUserName string, projectName string, subApps []string) string {
	lprojectName := utils.Lowercase(projectName)
	return fmt.Sprintf(MainTemplate,
		githubUserName,
		lprojectName,
		githubUserName,
		lprojectName,
		BuildServerSection(subApps),
		BuildSuperServerSection(subApps),
	)
}
