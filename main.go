package main

import (
	"os"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/tomkys144/gitea-teaset/cmd"
)

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
				Flags:  setupFlags,
				Action: cmd.Setup,
			},
			{
				Name:    "env_gen",
				Aliases: []string{"env"},
				Usage:   "Setups .env file (Don't need to be run for setup command)",
				Flags:   env_genFlags,
				Action:  cmd.EnvGen,
			},
			{
				Name:   "start",
				Usage:  "Starts Gitea server | NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a",
				Flags:  startFlags,
				Action: cmd.Start,
			},
			{
				Name:   "stop",
				Usage:  "Stops Gitea server | NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a",
				Flags:  stopFlags,
				Action: cmd.Stop,
			},
			{
				Name:    "update",
				Aliases: []string{"up"},
				Usage:   "Updates Teaset to newest version | NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a",
				Flags:   updateFlags,
				Action:  cmd.Update,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Displays current Teaset version | NOT IMPLEMENTED YET \xf0\x9f\x98\xa2\x0a",
				Flags:   versionFlags,
				Action:  cmd.Version,
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err.Error())
	}
}
