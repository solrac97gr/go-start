package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
)

func Test_NewGoModTemplate(t *testing.T) {
	githubUser := "user"
	projectName := "test"
	goVersion := "1.14"
	expected := `module github.com/user/test

go 1.14

require (
    github.com/gofiber/fiber/v2 v2.17.0
)
`
	if got := templates.NewGoModTemplate(githubUser, projectName, goVersion); got != expected {
		t.Errorf("NewGoModTemplate() = %v, want %v", got, expected)
	}
}
