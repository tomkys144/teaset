// Teaset
//
// Teaset is service desk for gitea
package main

import (
	"os"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/tomkys144/gitea-teaset/cmd"
)

// Main function. Sets up logrus formatter and cli app
func main() {
	// Logrus setup
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"topic", "event"},
		ShowFullLevel:   true,
		TrimMessages:    true,
		TimestampFormat: time.RFC850,
	})

	app := &cli.App{
		Name:                 "Teaset",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:   "setup",
				Usage:  "Setups Teaset environment",
				Flags:  SetupFlags,
				Action: cmd.Setup,
			},
			{
				Name:    "env_gen",
				Aliases: []string{"env"},
				Usage:   "Setups .env file (Don't need to be run for setup command)",
				Flags:   Env_genFlags,
				Action:  cmd.EnvGen,
			},
			{
				Name:   "start",
				Usage:  "Starts Gitea server | NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a",
				Flags:  StartFlags,
				Action: cmd.Start,
			},
			{
				Name:   "stop",
				Usage:  "Stops Gitea server | NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a",
				Flags:  StopFlags,
				Action: cmd.Stop,
			},
			{
				Name:    "update",
				Aliases: []string{"up"},
				Usage:   "Updates Teaset to newest version | NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a",
				Flags:   UpdateFlags,
				Action:  cmd.Update,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Displays current Teaset version | NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a",
				Flags:   VersionFlags,
				Action:  cmd.Version,
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err.Error())
	}
}
