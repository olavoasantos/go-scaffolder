package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"olavoasantos/scaffolder/fileManager"
	"olavoasantos/scaffolder/templateManager"
	"olavoasantos/scaffolder/utilities"
	"os"
	"path/filepath"

	"github.com/cbroglie/mustache"
	"github.com/urfave/cli/v2"
)

type Variables struct {
	PATH string
	NAME utilities.Variations
}

type ScaffolderConfig struct {
	Templates map[string]string
}

var MakeCommand = &cli.Command{
	Name:    "make",
	Aliases: []string{"m"},
	Usage:   "make a new file based on a template",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "Name of the given component. Defaults to the file name or the directory name.",
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "config.json",
			Usage:   "Load configuration from `FILE` relative path",
		},
	},
	Action: func(c *cli.Context) error {
		var err error

		// Get CLI variables
		templatePath := c.Args().Get(0)
		output := c.Args().Get(1)
		configFlag := c.String("config")

		// Get fallback values for "name"
		name := c.String("name")
		if name == "" {
			name = fileManager.FileName(filepath.Base(output))
			if name == "index" {
				name = filepath.Base(filepath.Dir(output))
			}
		}

		// Parse config
		config := ScaffolderConfig{}
		configPath, err := fileManager.PathTo(configFlag)
		utilities.Check(err)
		configJson, err := fileManager.GetContentsOf(configPath)
		if err == nil {
			err = json.Unmarshal([]byte(configJson), &config)
			utilities.Check(err)
		} else {
			if configFlag != "config.json" {
				fmt.Println("Config file  \"" + configFlag + "\" was not found")
			}
		}

		// Initialize template manager
		templates := templateManager.New(config.Templates)

		// Get template content
		template, err := templates.Get(templatePath)
		utilities.Check(err)

		// Render file contents
		result, err := mustache.Render(template, Variables{NAME: utilities.VariationsOf(name), PATH: output})
		utilities.Check(err)

		// Write file
		outputPath, err := fileManager.PathTo(output)
		utilities.Check(err)
		err = os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
		utilities.Check(err)
		err = ioutil.WriteFile(outputPath, []byte(result), os.ModePerm)
		utilities.Check(err)

		fmt.Println("Created file", outputPath)

		return nil
	},
}
