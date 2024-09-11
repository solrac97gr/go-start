package application

import (
	"github.com/solrac97gr/carlos/internal/cat/domain/ports"
)

type CatApp struct {
	repository	ports.CatRepository
}

func NewCatApp(repo ports.CatRepository) (*CatApp, error) {
	return &CatApp{
		repository: repo,
	}, nil
}
