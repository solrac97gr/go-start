package factory

import (
	"github.com/solrac97gr/carlos/internal/dog/application"
	"github.com/solrac97gr/carlos/internal/dog/infrastructure/repository"
	"github.com/solrac97gr/carlos/internal/dog/infrastructure/handler"
	"github.com/solrac97gr/carlos/internal/dog/infrastructure/server"
)

func BuildDogServer() *server.DogServer {
	repo, err := repository.NewDogRepository()
	if err != nil {
		panic(err)
	}
	app, err := application.NewDogApp(repo)
	if err != nil {
		panic(err)
	}
	handler, err := handler.NewDogHandler(app)
	if err != nil {
		panic(err)
	}
	server, err := server.NewDogServer(handler)
	if err != nil {
		panic(err)
	}
	return server
}
