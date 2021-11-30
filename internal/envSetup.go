package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func EnvSetup(source string, dest string) error {
	sourceAbs, _ := filepath.Abs(source)
	var destAbs string
	if dest == "" {
		destAbs, _ = filepath.Abs("./")
	} else {
		destAbs, _ = filepath.Abs(dest)
	}

	log.WithFields(log.Fields{
		"event": "generation",
		"topic": ".env",
	}).Infof("Creating .env file from \033[35m%s\033[0m in \033[35m%s\033[0m", sourceAbs, destAbs)

	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		dest_abs, _ := filepath.Abs(dest)
		log.WithFields(log.Fields{
			"event": "generation",
			"topic": ".env",
		}).Errorf("Could not create/open directory \033[35m%s\033[0m", dest_abs)
		return err
	}

	dest += "/.env"
	destAbs, _ = filepath.Abs(dest)
	env, _ := godotenv.Read(dest)
	envNew, _ := godotenv.Read(source)

	for k, v := range envNew {
		env[k] = v
	}

	f, err := os.Create(dest)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "generation",
			"topic": ".env",
		}).Errorf("Could not create/open \033[35m%s\033[0m", destAbs)
		return err
	}
	defer f.Close()

	// ---- SETUP ------------------
	// ---- FILES -----
	fmt.Printf("Do you want to setup values for paths?\nIf not values provided in .env file will be used [y/n] (y): ")
	res := "y"
	fmt.Scanln(&res)

	properSetUp := !(env == nil || env["PATH_HOME"] == "" || env["PATH_CUSTOM"] == "")

	if res != "y" && !properSetUp {
		res = "n"
		fmt.Printf("Do you really want to skip paths setup? Currently DB is not set up properly/not set up at all? [y/n] (n): ")
		fmt.Scanln(&res)

		if res == "y" {
			res = "n"
		} else {
			res = "y"
		}
	}

	if res == "y" {
		f.WriteString("# PATHS\n")
		var value string = ""
		for value == "" {
			fmt.Println("Path to root directory:")
			fmt.Scanln(&value)
			f.WriteString("PATH_HOME=\"" + value + "\"\n")
		}
		value = ""

		for value == "" {
			fmt.Println("Path to custom directory (used to store customized files):")
			fmt.Scanln(&value)
			f.WriteString("PATH_CUSTOM=\"" + value + "\"\n")
		}
	} else {
		f.WriteString("# PATHS\n")
		f.WriteString("PATH_HOME=\"" + env["PATH_HOME"] + "\"\n")
		f.WriteString("PATH_CUSTOM=\"" + env["PATH_CUSTOM"] + "\"\n")
	}

	// ---- DB ------
	fmt.Printf("Do you want to setup values for database?\nIf not values provided in .env file will be used [y/n] (y): ")
	res = "y"
	fmt.Scanln(&res)

	properSetUp = !(env == nil || env["DB_HOST"] == "" || env["DB_NAME"] == "" || env["DB_USER"] == "" || env["DB_PWD"] == "")

	if res != "y" && !properSetUp {
		res = "n"
		fmt.Printf("Do you really want to skip DB setup? Currently DB is not set up properly/not set up at all? [y/n] (n): ")
		fmt.Scanln(&res)

		if res == "y" {
			res = "n"
		} else {
			res = "y"
		}
	}

	if res == "y" {
		f.WriteString("\n# DB\n")
		var value string = ""
		for value == "" {
			fmt.Println("Database host:")
			fmt.Scanln(&value)
			f.WriteString("DB_HOST=\"" + value + "\"\n")
		}
		value = ""

		for value == "" {
			fmt.Println("Database name:")
			fmt.Scanln(&value)
			f.WriteString("DB_NAME=\"" + value + "\"\n")
		}
		value = ""

		for value == "" {
			fmt.Println("Database user:")
			fmt.Scanln(&value)
			f.WriteString("DB_USER=\"" + value + "\"\n")
		}
		value = ""

		for value == "" {
			fmt.Println("Database password (if none wite \"NONE\"):")
			fmt.Scanln(&value)
			f.WriteString("DB_PWD=\"" + value + "\"\n")
		}
	} else {
		f.WriteString("\n# DB\n")
		f.WriteString("DB_HOST=\"" + env["DB_HOST"] + "\"\n")
		f.WriteString("DB_NAME=\"" + env["DB_NAME"] + "\"\n")
		f.WriteString("DB_USER=\"" + env["DB_USER"] + "\"\n")
		f.WriteString("DB_PWD=\"" + env["DB_PWD"] + "\"\n")
	}

	// SYNC
	log.WithFields(log.Fields{
		"event": "generation",
		"topic": ".env",
	}).Info("Writing data to .env file")

	err = f.Sync()
	if err != nil {
		log.WithFields(log.Fields{
			"event": "generation",
			"topic": ".env",
		}).Errorf("Could not write to \033[35m%s\033[0m", destAbs)
		return err
	}

	err = godotenv.Load(dest)
	if err != nil {
		log.WithFields(log.Fields{
			"event": ".env loading",
			"topic": ".env",
		}).Errorf("Could not read from \033[35m%s\033[0m", destAbs)
		return err
	}
	return err
}

func EnvSilentSetup(source string, dest string) error {

	sourceAbs, _ := filepath.Abs(source)
	var destAbs string
	if dest == "" {
		destAbs, _ = filepath.Abs("./")
	} else {
		destAbs, _ = filepath.Abs(dest)
	}

	log.WithFields(log.Fields{
		"event": "generation",
		"topic": ".env",
	}).Infof("Creating .env file from \033[35m%s\033[0m in \033[35m%s\033[0m", sourceAbs, destAbs)

	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "generation",
			"topic": ".env",
		}).Errorf("Could not create/open directory \033[35m%s\033[0m", destAbs)
		return err
	}
	dest += "/.env"
	destAbs, _ = filepath.Abs(dest)
	env, _ := godotenv.Read(dest)
	envNew, _ := godotenv.Read(source)

	for k, v := range envNew {
		env[k] = v
	}

	f, err := os.Create(dest)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "generation",
			"topic": ".env",
		}).Errorf("Could not create/open \033[35m%s\033[0m", destAbs)
		return err
	}
	defer f.Close()

	err = godotenv.Write(env, dest)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "generation",
			"topic": ".env",
		}).Errorf("Could not write to \033[35m%s\033[0m", destAbs)
		return err
	}

	err = godotenv.Load(dest)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "loading",
			"topic": ".env",
		}).Errorf("Could not read from \033[35m%s\033[0m", destAbs)
		return err
	}

	log.WithFields(log.Fields{
		"event": "loading",
		"topic": ".env",
	}).Infof("Loaded .env file from \033[35m%s\033[0m", destAbs)
	return nil
}
