package templates

import "fmt"

var PortsTemplate = `package ports

type %sRepository interface {
}

type %sAplication interface {
}

type %sHandler interface {
}
`

func NewPortsTemplate(name string) string {
	return fmt.Sprintf(PortsTemplate, name, name, name)
}
