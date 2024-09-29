package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jonathongardner/linuxhelp/setup/github"
	"github.com/jonathongardner/linuxhelp/setup/ssh"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var githubSSHPullCommand = &cli.Command{
	Name:  "pull",
	Usage: "Github pull ssh keys for user",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Aliases:     []string{"p"},
			Usage:       "Path to public key",
			DefaultText: "~/.ssh/authorized_keys",
		},
		&cli.BoolFlag{
			Name:  "dry-run",
			Usage: "Dont save to a file",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "clear",
			Usage: "Clear old authorized keys",
			Value: false,
		},
		&cli.StringFlag{
			Name:  "key-id",
			Usage: "Github keyid to copy",
		},
		&cli.StringFlag{
			Name:  "fingerprint",
			Usage: "Sha256 to copy",
		},
	},
	Action: func(c *cli.Context) error {
		path := c.String("path")
		if path == "" {
			homedir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("error getting home dir (%v)", err)
			}
			path = filepath.Join(homedir, ".ssh/authorized_keys")
		}
		authKey, err := ssh.NewAuthKey(path, c.Bool("dry-run"))
		if err != nil {
			return err
		}
		defer authKey.Close()

		githubApi := github.NewApi(c.String("host"), c.String("user"), c.String("token"))
		keys, err := githubApi.SSHKeys()
		if err != nil {
			return err
		}

		if c.Bool("clear") {
			err := authKey.Clear()
			if err != nil {
				return err
			}
		}

		err = authKey.Parse()
		if err != nil {
			return err
		}

		fingerprint := c.String("fingerprint")
		keyId := c.String("key-id")
		log.Infof("Found %v", len(keys))
		for _, key := range keys {
			if fingerprint != "" && fingerprint != key.Fingerprint {
				continue
			}
			if keyId != "" && keyId == key.GetID() {
				continue
			}

			added, err := authKey.Add(key)
			if err != nil {
				return err
			}
			if added {
				log.Infof("Added Key %v (%v)", key.ID, key.Fingerprint)
			}
		}

		return nil
	},
}
