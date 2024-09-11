package templates

import (
	"fmt"

	"github.com/solrac97gr/go-start/utils"
)

var RepositoryTemplate = `package repository

type %sRepository struct {
}

func New%sRepository() (*%sRepository, error) {
	return &%sRepository{}, nil
}
`

func NewRepositoryTemplate(repoName string) string {
	lowerName := utils.Lowercase(repoName)
	capitalizedName := utils.Capitalize(lowerName)
	return fmt.Sprintf(RepositoryTemplate, capitalizedName, capitalizedName, capitalizedName, capitalizedName)
}
