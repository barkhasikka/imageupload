package files

import (
	"log"
	"os"
)

func CreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Println(" -- error creating " + dir)
			return err
		}
	}
	return nil
}
