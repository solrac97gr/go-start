package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
	"github.com/stretchr/testify/assert"
)

func Test_NewHandlersTemplate(t *testing.T) {
	var expected = `package handler

import (
	"github.com/user/project/internal/test/domain/ports"
)

type TestHandler struct {
	application ports.TestAplication
}

func NewTestHandler(app ports.TestAplication) (*TestHandler, error) {
	return &TestHandler{
		application: app,
	}, nil
}
`
	var result = templates.NewHandlersTemplate("user", "project", "Test")
	assert.Equal(t, expected, result)
}
