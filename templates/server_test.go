package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
	"github.com/stretchr/testify/assert"
)

func Test_NewServerTemplate(t *testing.T) {
	var expected = `package server

import (
	"github.com/user/project/internal/domain/ports"

	"github.com/gofiber/fiber/v2"
)

const (
	TestPath = "/test"
)

type TestServer struct {
	App *fiber.App
	Path string
}

func NewTestServer(hdl ports.TestHandler) (*TestServer, error) {
	if hdl == nil {
		return nil, fmt.Errorf("handler is nil")
	}

	app := fiber.New()
	// Here you can register your routes
	// app.Get("/example", hdl.Example)

	return &TestServer{
		App: app,
		Path: TestPath,
	}, nil
}

func (s *TestServer) GetPath() string {
	return s.Path
}

func (s *TestServer) GetApp() *fiber.App {
	return s.App
}

func (s *TestServer) Start(port string) error {
	return s.App.Listen(":"+port)
}
`
	var result = templates.NewServerTemplate("user", "project", "Test")
	assert.Equal(t, expected, result)
}
