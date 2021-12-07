package internal

import (
	"archive/zip"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	err = publicSetup()
	if err != nil {
		return err
	}

	return nil
}

func publicSetup() error {
	Url := "https://github.com/tomkys144/teaset/releases/latest/download/frontend.zip"
	resp, err := http.Get(Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	path := os.Getenv("PATH_HOME") + "/tmp.zip"
	out, err := os.Create(path)
	if err != nil {
		println("here")
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	archive, err := zip.OpenReader(path)
	if err != nil {
		return nil
	}
	defer archive.Close()

	dst := os.Getenv("PATH_HOME") + "/public"

	for _, f := range archive.File {
		fp := filepath.Join(dst, f.Name)

		if !strings.HasPrefix(fp, filepath.Clean(dst)+string(os.PathSeparator)) {
			return errors.New("invalid file path")
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fp, os.ModePerm)
			continue
		}

		err := os.MkdirAll(filepath.Dir(fp), os.ModePerm)
		if err != nil {
			return err
		}

		dstFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		srcFile, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}

		dstFile.Close()
		srcFile.Close()
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
