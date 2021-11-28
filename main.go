package main

import (
	"os"

	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := &cli.App{
		Name: "Teaset",
		Commands: []*cli.Command{
			{
				Name:   "setup",
				Usage:  "Setups Teaset environment",
				Flags:  setup_flags,
				Action: cmd.Setup,
			},
		},
	}

	err = app.Run(os.Args)

	if err != nil {
		log.Fatal(err.Error())
	}
}
