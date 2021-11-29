package main

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/tomkys144/gitea-teaset/cmd"
)

func main() {
	// Logrus setup
	logFormatter := new(logrus.TextFormatter)
	logFormatter.FullTimestamp = true
	logrus.SetFormatter(logFormatter)

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
				Action: cmd.EnvGen,
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err.Error())
	}
}
