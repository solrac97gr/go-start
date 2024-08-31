package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
)

func Test_ApplicationTemplate(t *testing.T) {
	expected := `package application
type TestApp struct {
	repository	ports.TestRepository
}

func NewTestApp(repo ports.TestRepository) *TestApp {
	return &TestApp{
		repository: repo,
	}
}
`
	result := templates.NewApplicationTemplate("Test")
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
