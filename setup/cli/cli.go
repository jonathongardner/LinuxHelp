package cli

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Run(version string) error {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(version)
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}

	app := &cli.App{
		Name:    "jelp",
		Version: version,
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
