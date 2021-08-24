package templates

import (
	_ "embed"
)

//go:embed default-templates/react-component.tsx.mustache
var reactComponentTsx string

var List = map[string]string{
	"react-component.tsx": reactComponentTsx,
}
