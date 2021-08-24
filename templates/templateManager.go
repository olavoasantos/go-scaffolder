package templates

import (
	"errors"
	"olavoasantos/scaffolder/file-manager"
	"olavoasantos/scaffolder/utilities"
	"os"
)

type TemplateManager struct {
	list map[string]string
}

func NewManager(template ...map[string]string) *TemplateManager {
	templates := TemplateManager{list: List}

	if len(template) >= 1 {
		templates.Set(template[0])
	}

	return &templates
}

func (t *TemplateManager) All() map[string]string {
	return t.list
}

func (t *TemplateManager) Set(templates map[string]string) *TemplateManager {
	for key, value := range templates {
		t.list[key] = value
	}

	return t
}

func (t *TemplateManager) Has(template string) bool {
	if _, ok := t.list[template]; ok {
		return true
	}

	return false
}

func (t *TemplateManager) Get(template string) (string, error) {
	if t.Has(template) {
		path, err := fileManager.PathTo(t.list[template])
		if err != nil {
			return "", err
		}

		content, err := fileManager.GetContentsOf(path)
		if err != nil {
			return t.list[template], nil
		}

		return content, nil
	}

	path, err := fileManager.PathTo(template)
	utilities.Check(err)

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fileManager.GetContentsOf(path)
	}

	return "", errors.New("template \"" + template + "\" does not exist. Make sure to register this on the config file or that this is a valid relative path to the template file.")
}
