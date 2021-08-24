package defaultTemplates

import (
	_ "embed"
)

//go:embed templates/react-component.tsx.mustache
var reactComponentTsx string

var List = map[string]string{
	"react-component.tsx": reactComponentTsx,
}
