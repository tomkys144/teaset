package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Stop(c *cli.Context) error {
	// TODO: Write 'stop' command
	// NOTE: Sends stop signal to detached process
	log.WithFields(log.Fields{
		"topic": "stop",
	}).Warning("NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a")
	return nil
}
