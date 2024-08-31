package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
	"github.com/stretchr/testify/assert"
)

func Test_RepositoryTemplate(t *testing.T) {
	expected := `package respository

type TestRepository struct {
}

func NewTestRepository() (*TestRepository, error) {
	return &TestRepository{}, nil
}
`
	result := templates.NewRepositoryTemplate("Test")
	assert.Equal(t, expected, result)
}
