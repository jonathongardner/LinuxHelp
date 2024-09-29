package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jonathongardner/linuxhelp/setup/github"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

var githubSSHPushCommand = &cli.Command{
	Name:  "push",
	Usage: "Github push public ssh keys for user",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Aliases:     []string{"p"},
			Usage:       "Path to public key",
			DefaultText: "~/.ssh/id_rsa.pub",
		},
		&cli.StringFlag{
			Name:  "title",
			Usage: "A description name for the key",
		},
	},
	Action: func(c *cli.Context) error {
		path := c.String("path")
		if path == "" {
			homedir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("error getting home dir (%v)", err)
			}
			path = filepath.Join(homedir, ".ssh/id_rsa.pub")
		}

		githubApi := github.NewApi(c.String("host"), c.String("user"), c.String("token"))

		key, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error opening file %v (%v)", path, err)
		}

		fingerprint, err := githubApi.SSHSave(string(key), c.String("title"))
		if err != nil {
			return fmt.Errorf("error adding %v (%v)", fingerprint, err)
		}
		log.Infof("Added %v ssh", fingerprint)

		return nil
	},
}
