package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/nohara-cloud/nohara-node/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
