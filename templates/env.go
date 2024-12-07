package templates

var EnvTemplate = `MONGO_USER=root
MONGO_PASSWORD=password
MONGO_DB_NAME=db`

func NewEnvTemplate() string {
	return EnvTemplate
}
