package main

import (
	_ "embed"

	"github.com/jonathongardner/linuxhelp/setup/cli"

	log "github.com/sirupsen/logrus"
)

//go:embed version.txt
var version string

func main() {
	err := cli.Run(version)
	if err != nil {
		log.Fatal(err)
	}
}
