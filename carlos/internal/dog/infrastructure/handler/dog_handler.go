package handler

import (
	"github.com/solrac97gr/carlos/internal/dog/domain/ports"
)

type DogHandler struct {
	application ports.DogAplication
}

func NewDogHandler(app ports.DogAplication) (*DogHandler, error) {
	return &DogHandler{
		application: app,
	}, nil
}
