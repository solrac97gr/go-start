package templates

var GitignoreTemplate = `.env
./bin`

func NewGitIgnoreTemplate() string {
	return GitignoreTemplate
}
