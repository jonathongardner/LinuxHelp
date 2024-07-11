package main

import (
	"github.com/jonathongardner/linuxhelp/setup/cli"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := cli.Run()
	if err != nil {
		log.Fatal(err)
	}
}