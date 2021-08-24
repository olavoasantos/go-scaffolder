package fileManager

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func PathTo(relative string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", errors.New("could not determine the command working directory")
	}

	return filepath.Join(cwd, relative), nil
}

func GetContentsOf(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New("file \"" + path + "\" was not found")
	}

	return string(content), nil
}

func FileName(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
