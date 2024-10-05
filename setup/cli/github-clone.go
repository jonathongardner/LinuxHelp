package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var githubCloneCommand = &cli.Command{
	Name:  "clone",
	Usage: "Github clone this repo",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Aliases:     []string{"p"},
			Usage:       "Path to folder to clone into",
			DefaultText: "~/Projects",
		},
	},
	Action: func(c *cli.Context) error {
		path := c.String("path")
		if path == "" {
			homedir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("error getting home dir (%v)", err)
			}
			path = filepath.Join(homedir, "Projects")
		}
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory %v (%v)", path, err)
		}

		fmt.Println("cd %v && git clone git@github.com:jonathongardner/LinuxHelp.git", path)

		return nil
	},
}
