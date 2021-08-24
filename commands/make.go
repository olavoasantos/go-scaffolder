package commands

import (
	"fmt"
	"olavoasantos/scaffolder/configuration"
	fileManager "olavoasantos/scaffolder/file-manager"
	"olavoasantos/scaffolder/templates"
	"olavoasantos/scaffolder/utilities"
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
	Name:      "make",
	Aliases:   []string{"m"},
	Usage:     "Make a new file based on a template",
	ArgsUsage: "{template name | relative path to template} {relative output path}",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "Name of the given component. Defaults to the file name or the directory name.",
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   configuration.DefaultConfigPath,
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
		config := configuration.Load(configFlag)

		// Initialize template manager
		templates := templates.NewManager(config.Templates)

		// Get template content
		template, err := templates.Get(templatePath)
		utilities.Check(err)

		// Render file contents
		result, err := mustache.Render(template, Variables{
			NAME: utilities.VariationsOf(name),
			PATH: output,
		})
		utilities.Check(err)

		// Write file
		err = fileManager.Write(output, result)
		utilities.Check(err)

		fmt.Println("Created file", output)

		return nil
	},
}
