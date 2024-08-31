package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
)

func Test_NewHandlersTemplate(t *testing.T) {
	var expected = `package handlers

import (
	"github.com/user/project/internal/test/domain/ports"
)

type TestHandler struct {
	application ports.TestAplication
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}
`
	var result = templates.NewHandlersTemplate("user", "project", "Test")
	if result != expected {
		t.Errorf("Expected: %s\nGot: %s", expected, result)
	}
}
