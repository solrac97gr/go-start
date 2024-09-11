package repository

type CatRepository struct {
}

func NewCatRepository() (*CatRepository, error) {
	return &CatRepository{}, nil
}
