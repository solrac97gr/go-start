package server

import (
	"github.com/solrac97gr/carlos/internal/cat/domain/ports"

	"github.com/gofiber/fiber/v2"
	"fmt"
)

const (
	CatPath = "/cat"
)

type CatServer struct {
	App *fiber.App
	Path string
}

func NewCatServer(hdl ports.CatHandler) (*CatServer, error) {
	if hdl == nil {
		return nil, fmt.Errorf("handler is nil")
	}

	app := fiber.New()
	// Here you can register your routes
	// app.Get("/example", hdl.Example)

	return &CatServer{
		App: app,
		Path: CatPath,
	}, nil
}

func (s *CatServer) GetPath() string {
	return s.Path
}

func (s *CatServer) GetApp() *fiber.App {
	return s.App
}

func (s *CatServer) Start(port string) error {
	return s.App.Listen(":"+port)
}
