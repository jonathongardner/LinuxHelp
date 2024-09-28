package cli

import (
	"fmt"
	"os"

	"github.com/jonathongardner/linuxhelp/setup/app"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Run() error {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}

	app := &cli.App{
		Name:    "jelp",
		Version: app.Version,
		Usage:   "Command line tool to help setup linux just how i like it.",
		Commands: []*cli.Command{
			githubCommand,
			dotCommand,
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "logging level",
			},
			// &cli.StringFlag{
			// 	Name:    "path",
			// 	Aliases: []string{"p"},
			// 	Usage:   "Project Path",
			// },
		},
		Before: func(c *cli.Context) error {
			if c.Bool("verbose") {
				log.SetLevel(log.DebugLevel)
				log.Debug("Setting to debug...")
			}
			return nil
		},
	}
	return app.Run(os.Args)
}
