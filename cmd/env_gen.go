package cmd

import (
	"github.com/tomkys144/teaset/internal"
	"github.com/urfave/cli/v2"
)

func EnvGen(c *cli.Context) error {
	err := internal.EnvSetup(c.Path("source"), c.Path("destination"))
	return err
}
