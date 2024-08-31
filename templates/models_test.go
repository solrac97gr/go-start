package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
)

func Test_ModelsTemplate(t *testing.T) {
	expected := `package models

type Test struct {
}

func NewTest() *Test {
	return &Test{}
}
`
	result := templates.NewModelsTemplate("Test")
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
