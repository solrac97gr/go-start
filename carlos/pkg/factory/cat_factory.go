package factory

import (
	"github.com/solrac97gr/carlos/internal/cat/application"
	"github.com/solrac97gr/carlos/internal/cat/infrastructure/repository"
	"github.com/solrac97gr/carlos/internal/cat/infrastructure/handler"
	"github.com/solrac97gr/carlos/internal/cat/infrastructure/server"
)

func BuildCatServer() *server.CatServer {
	repo, err := repository.NewCatRepository()
	if err != nil {
		panic(err)
	}
	app, err := application.NewCatApp(repo)
	if err != nil {
		panic(err)
	}
	handler, err := handler.NewCatHandler(app)
	if err != nil {
		panic(err)
	}
	server, err := server.NewCatServer(handler)
	if err != nil {
		panic(err)
	}
	return server
}
