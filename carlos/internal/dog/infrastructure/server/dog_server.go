package server

import (
	"github.com/solrac97gr/carlos/internal/dog/domain/ports"

	"github.com/gofiber/fiber/v2"
	"fmt"
)

const (
	DogPath = "/dog"
)

type DogServer struct {
	App *fiber.App
	Path string
}

func NewDogServer(hdl ports.DogHandler) (*DogServer, error) {
	if hdl == nil {
		return nil, fmt.Errorf("handler is nil")
	}

	app := fiber.New()
	// Here you can register your routes
	// app.Get("/example", hdl.Example)

	return &DogServer{
		App: app,
		Path: DogPath,
	}, nil
}

func (s *DogServer) GetPath() string {
	return s.Path
}

func (s *DogServer) GetApp() *fiber.App {
	return s.App
}

func (s *DogServer) Start(port string) error {
	return s.App.Listen(":"+port)
}
