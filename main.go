package main

import (
	"olavoasantos/scaffolder/commands"
	"olavoasantos/scaffolder/utilities"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	// Register commands on the CLI instance
	app := &cli.App{
		Commands: []*cli.Command{
			commands.MakeCommand,
		},
	}

	// Order flags and commands for the help block
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	// Execute the CLI and panic if anything goes wrong on the execution
	utilities.Check(app.Run(os.Args))
}
