package templates

import (
	"fmt"

	"github.com/solrac97gr/go-start/utils"
)

var SuperServerTemplate = `package sserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type SubApp interface {
	GetPath() string
	GetApp() *fiber.App
	Start(string) error
}

type Server struct {
	Apps	[]SubApp
}

func NewServer(apps []SubApp) *Server {
	s := &Server{
		Apps:	apps,
	}
	if len(s.Apps) == 0 {
		fmt.Println("0 app set in the server")
	}
	return s
}

func (s *Server) Initialize(port string) error {
	app := fiber.New(fiber.Config{
		AppName: "%s",
	})

	app.Use(cors.New())

	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		},
	))

	baseURL := app.Group("%s/v1")
	// Stack Endpoint
	baseURL.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})
	// Ping Endpoint
	baseURL.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	for _, subapp := range s.Apps {
		baseURL.Mount(subapp.GetPath(), subapp.GetApp())
	}

	err := app.Listen(":" + port)
	if err != nil {
		return err
	}
	return nil
}
`

func NewSuperServerTemplate(projectName string) string {
	lowerCase := utils.Lowercase(projectName)
	return fmt.Sprintf(SuperServerTemplate, lowerCase, lowerCase)
}
