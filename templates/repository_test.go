package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
)

func Test_RepositoryTemplate(t *testing.T) {
	expected := `package respository

type TestRepository struct {
}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}
`
	result := templates.NewRepositoryTemplate("Test")
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
