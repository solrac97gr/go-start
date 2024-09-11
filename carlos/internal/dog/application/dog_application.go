package application

import (
	"github.com/solrac97gr/carlos/internal/dog/domain/ports"
)

type DogApp struct {
	repository	ports.DogRepository
}

func NewDogApp(repo ports.DogRepository) (*DogApp, error) {
	return &DogApp{
		repository: repo,
	}, nil
}
