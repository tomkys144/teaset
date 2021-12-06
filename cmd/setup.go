package cmd

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/tomkys144/teaset/internal"
)

var db *sql.DB
var err error

func Setup(c *cli.Context) error {
	// log level setup
	if c.Bool("verbose") {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	// files setup
	if c.Bool("interactive") {
		err = internal.EnvSetup(c.Path("source"), c.Path("destination"))
	} else {
		err = internal.EnvSilentSetup(c.Path("source"), c.Path("destination"))
	}
	if err != nil {
		return err
	}

	err = internal.FileSetup()
	if err != nil {
		return err
	}
	// db setup
	db, err = internal.InitDB()
	if err != nil {
		return err
	}

	err = internal.CreateTable(db)
	if err != nil {
		return err
	}

	return err
}
