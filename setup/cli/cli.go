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

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "logging level",
		},
	}

	app := &cli.App{
		Name:    "starter",
		Version: app.Version,
		Usage:   "Example starter app for cli tools!",
		Commands: []*cli.Command{
			githubCommand,
			dotCommand,
		},
		Flags: flags,
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
