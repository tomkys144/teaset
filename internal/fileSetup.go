package internal

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func FileSetup() error {
	// generate bin directory
	err := binSetup()
	if err != nil {
		return err
	}
	return nil
}

func binSetup() error {
	ex, err := os.Executable()
	if err != nil {
		log.WithFields(log.Fields{
			"topic": "file generation",
			"event": "binaries generation",
		}).Fatalln(err.Error())
		return err
	}

	src, err := os.Open(ex)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "copying Teaset executable to bin",
			"topic": "file generation",
		}).Fatalln(err.Error())
		return err
	}
	defer src.Close()

	err = os.MkdirAll(os.Getenv("PATH_HOME")+"/bin", os.ModePerm)
	if err != nil {
		log.WithFields(log.Fields{
			"topic": "file generation",
			"event": "binaries generation",
		}).Fatalln(err.Error())
		return err
	}

	dst, err := os.Create(os.Getenv("PATH_HOME") + "/bin/teaset")
	if err != nil {
		log.WithFields(log.Fields{
			"topic": "file generation",
			"event": "Creating teaset binary",
		}).Fatalln(err.Error())
		return err
	}
	defer dst.Close()

	err = dst.Chmod(0755)
	if err != nil {
		log.WithFields(log.Fields{
			"topic": "file generation",
			"event": "Creating teaset binary",
		}).Fatalln(err.Error())
		return err
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		log.WithFields(log.Fields{
			"topic": "file generation",
			"event": "Creating teaset binary",
		}).Fatalln(err.Error())
		return err
	}

	return nil
}
