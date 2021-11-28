package internal

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func InitDB() (*sql.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	var dbString string

	if dbPwd != "NONE" {

		dbString = dbUser + ":" + dbPwd + "@tcp(" + dbHost + ")/" + dbName
	} else {
		dbString = dbUser + "@tcp(" + dbHost + ")/" + dbName
	}

	log.WithFields(log.Fields{
		"event": "Database connection",
		"topic": "DB",
	}).Infof("Connecting to databse %s on host %s as %s.\n", dbName, dbHost, dbUser)
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "Database connection",
			"topic": "DB",
		}).Fatal(err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.WithFields(log.Fields{
			"event": "Database connection",
			"topic": "DB",
		}).Fatal(err.Error())
		return nil, err
	}

	return db, err
}

func CreateTable(db *sql.DB) error {
	query := ""
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	result, err := db.ExecContext(ctx, query)

	if err != nil {
		log.WithFields(log.Fields{
			"event": "Tables creation",
			"topic": "DB",
		}).Fatal(err.Error())
	}

	log.WithFields(log.Fields{
		"event": "Tables creation",
		"topic": "DB",
	}).Infof("Table %s created.\nRows affected: %d", "Projects", result.RowsAffected)
	return err
}
