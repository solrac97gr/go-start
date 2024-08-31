package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
)

func Test_NewGoModTemplate(t *testing.T) {
	projectName := "test"
	goVersion := "1.14"
	expected := `module test

go 1.14
`
	if got := templates.NewGoModTemplate(projectName, goVersion); got != expected {
		t.Errorf("NewGoModTemplate() = %v, want %v", got, expected)
	}
}
