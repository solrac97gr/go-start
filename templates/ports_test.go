package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, expected, result)
	}

}
