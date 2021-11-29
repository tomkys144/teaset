package internal

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func EnvSetup() error {
	env, _ := godotenv.Read("./.env")

	f, err := os.Create("./.env")
	if err != nil {
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

	f.Sync()

	err = godotenv.Load("./.env")
	return err
}
