package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jonathongardner/linuxhelp/setup/dot"

	"github.com/urfave/cli/v2"
)

var dotCommand = &cli.Command{
	Name:  "dot",
	Usage: "Export/Copy dot files",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "project",
			Aliases:     []string{"p"},
			Usage:       "Path to linux help",
			DefaultText: "~/Projects/LinuxHelp",
		},
	},
	Action: func(c *cli.Context) error {
		project := c.String("project")
		homedir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("error getting home dir (%v)", err)
		}
		if project == "" {
			project = filepath.Join(homedir, "/Projects/LinuxHelp")
		}
		return dot.Export(homedir, project)
	},
}
