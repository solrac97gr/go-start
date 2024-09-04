package templates

import (
	"fmt"

	"github.com/solrac97gr/go-start/utils"
)

var ServerTemplate = `package server

import (
	"github.com/%s/%s/internal/%s/domain/ports"

	"github.com/gofiber/fiber/v2"
	"fmt"
)

const (
	%sPath = "/%s"
)

type %sServer struct {
	App *fiber.App
	Path string
}

func New%sServer(hdl ports.%sHandler) (*%sServer, error) {
	if hdl == nil {
		return nil, fmt.Errorf("handler is nil")
	}

	app := fiber.New()
	// Here you can register your routes
	// app.Get("/example", hdl.Example)

	return &%sServer{
		App: app,
		Path: %sPath,
	}, nil
}

func (s *%sServer) GetPath() string {
	return s.Path
}

func (s *%sServer) GetApp() *fiber.App {
	return s.App
}

func (s *%sServer) Start(port string) error {
	return s.App.Listen(":"+port)
}
`

func NewServerTemplate(githubUser, projectName, serverName string) string {
	var entityLower = utils.Lowercase(serverName)
	var entityCapitalized = utils.Capitalize(serverName)

	return fmt.Sprintf(ServerTemplate,
		githubUser,
		projectName,
		entityLower,
		entityCapitalized,
		entityLower,
		entityCapitalized,
		entityCapitalized,
		entityCapitalized,
		entityCapitalized,
		entityCapitalized,
		entityCapitalized,
		entityCapitalized,
		entityCapitalized,
		entityCapitalized,
	)
}
