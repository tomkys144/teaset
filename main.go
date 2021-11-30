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
				Flags:  setup_flags,
				Action: cmd.Setup,
			},
			{
				Name:   "env_gen",
				Usage:  "Setups .env file (Don't need to be run for setup command)",
				Flags:  env_gen_flags,
				Action: cmd.EnvGen,
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err.Error())
	}
}
