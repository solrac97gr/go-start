package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
	"github.com/stretchr/testify/assert"
)

func Test_ApplicationTemplate(t *testing.T) {
	expected := `package application

import (
	"github.com/solrac97gr/go-start/internal/test/domain/ports"
)

type TestApp struct {
	repository	ports.TestRepository
}

func NewTestApp(repo ports.TestRepository) (*TestApp, error) {
	return &TestApp{
		repository: repo,
	}, nil
}
`
	result := templates.NewApplicationTemplate("solrac97gr", "go-start", "Test")
	assert.Equal(t, expected, result)
}
