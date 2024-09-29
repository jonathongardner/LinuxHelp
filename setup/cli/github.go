package cli

import (
	// log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var githubSSHCommand = &cli.Command{
	Name:  "ssh",
	Usage: "Github ssh stuff",
	Subcommands: []*cli.Command{
		githubSSHPullCommand,
		githubSSHPushCommand,
	},
}

var githubCommand = &cli.Command{
	Name:  "github",
	Usage: "Github setup stuff",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Usage: "Github host (Must be compatible with application/vnd.github.v3+json)",
			Value: "api.github.com",
		},
		&cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u"},
			Required: true,
			Usage:    "User to pull ssh for",
			EnvVars:  []string{"GH_USER"},
		},
		&cli.StringFlag{
			Name:    "token",
			Usage:   "Authentication token",
			EnvVars: []string{"GH_TOKEN"},
		},
	},
	Subcommands: []*cli.Command{
		githubSSHCommand,
	},
}
