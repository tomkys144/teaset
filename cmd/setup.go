package cmd

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/tomkys144/gitea-teaset/internal"
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

	db, err = internal.InitDB()
	if err != nil {
		return err
	}
	return err
}
