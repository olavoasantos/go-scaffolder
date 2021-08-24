package configuration

import (
	"encoding/json"
	"fmt"
	fileManager "olavoasantos/scaffolder/file-manager"
	"olavoasantos/scaffolder/utilities"
)

type ScaffolderConfig struct {
	Templates map[string]string
}

var DefaultConfigPath string = "config.json"

func Load(relativePath string) ScaffolderConfig {
	config := ScaffolderConfig{}

	configPath, err := fileManager.PathTo(relativePath)
	utilities.Check(err)

	configJson, err := fileManager.GetContentsOf(configPath)
	if err == nil {
		err = json.Unmarshal([]byte(configJson), &config)
		utilities.Check(err)
	} else {
		if relativePath != DefaultConfigPath {
			fmt.Println("Config file  \"" + relativePath + "\" was not found")
		}
	}

	return config
}
