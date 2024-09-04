package templates

import (
	"fmt"

	"github.com/solrac97gr/go-start/utils"
)

var PortsTemplate = `package ports

type %sRepository interface {
}

type %sAplication interface {
}

type %sHandler interface {
}
`

func NewPortsTemplate(name string) string {
	lower := utils.Lowercase(name)
	capitalized := utils.Capitalize(lower)
	return fmt.Sprintf(PortsTemplate, capitalized, capitalized, capitalized)
}
