package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Update(c *cli.Context) error {
	// TODO: Write 'update' command
	// NOTE: Downloads either stabel or bleeding edge from github and updates. (Might prompt to rerun setup)
	log.WithFields(log.Fields{
		"topic": "update",
	}).Warning("NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a")
	return nil
}
