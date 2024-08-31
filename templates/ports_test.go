package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
)

func Test_NewPortsTemplate(t *testing.T) {
	var expected = `package ports

type TestRepository interface {
}

type TestAplication interface {
}

type TestHandler interface {
}
`
	var name = "Test"
	var result = templates.NewPortsTemplate(name)
	if result != expected {
		t.Errorf("Expected: %s\nGot: %s", expected, result)
	}

}
