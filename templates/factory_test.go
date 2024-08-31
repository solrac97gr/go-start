package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
	"github.com/stretchr/testify/assert"
)

func Test_NewFactoryTemplate(t *testing.T) {
	var expected = `package factory

import (
	"github.com/user/project/internal/test/application"
	"github.com/user/project/internal/test/infrastructure/repository"
	"github.com/user/project/internal/test/infrastructure/handler"
	"github.com/user/project/internal/test/infrastructure/server"
)

func BuildTestServer() *server.TestServer {
	repo, err := repository.NewTestRepository()
	if err != nil {
		panic(err)
	}
	app, err := application.NewTestApp(repo)
	if err != nil {
		panic(err)
	}
	handler, err := handler.NewTestHandler(app)
	if err != nil {
		panic(err)
	}
	server, err := server.NewTestServer(handler)
	if err != nil {
		panic(err)
	}
	return server
}
`
	var result = templates.NewFactoryTemplate("user", "project", "test")
	assert.Equal(t, expected, result)
}
