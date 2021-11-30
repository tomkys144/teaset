package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Start(c *cli.Context) error {
	// TODO: Write 'start' command
	// NOTE: Starts main server task in either detached mode or in running command line
	log.WithFields(log.Fields{
		"topic": "start",
	}).Warning("NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a")

	return nil
}
