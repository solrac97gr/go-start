package handler

import (
	"github.com/solrac97gr/carlos/internal/cat/domain/ports"
)

type CatHandler struct {
	application ports.CatAplication
}

func NewCatHandler(app ports.CatAplication) (*CatHandler, error) {
	return &CatHandler{
		application: app,
	}, nil
}
