package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
)

func Test_ApplicationTemplate(t *testing.T) {
	expected := `package application

import (
	"github.com/solrac97gr/go-start/internal/test/domain/ports"
)

type TestApp struct {
	repository	ports.TestRepository
}

func NewTestApp(repo ports.TestRepository) *TestApp {
	return &TestApp{
		repository: repo,
	}
}
`
	result := templates.NewApplicationTemplate("solrac97gr",  "go-start" , "Test")
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
