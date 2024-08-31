package templates

import "fmt"

var RepositoryTemplate = `package respository

type %sRepository struct {
}

func New%sRepository() *%sRepository {
	return &%sRepository{}
}
`

func NewRepositoryTemplate(repoName string) string {
	return fmt.Sprintf(RepositoryTemplate, repoName, repoName, repoName, repoName)
}
