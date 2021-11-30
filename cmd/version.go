package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Version(c *cli.Context) error {
	// TODO: Write 'version' command
	// NOTE: Returns current version (might show newest available)
	log.WithFields(log.Fields{
		"topic": "version",
	}).Warning("NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a")
	return nil
}
