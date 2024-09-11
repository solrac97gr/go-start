package repository

type DogRepository struct {
}

func NewDogRepository() (*DogRepository, error) {
	return &DogRepository{}, nil
}
