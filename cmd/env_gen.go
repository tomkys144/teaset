package cmd

import (
	"github.com/tomkys144/gitea-teaset/internal"
	"github.com/urfave/cli/v2"
)

func EnvGen(c *cli.Context) error {
	err := internal.EnvSetup()
	return err
}
