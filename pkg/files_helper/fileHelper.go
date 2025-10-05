package files_helper

import (
	"bytes"
	"log"
	"os"
)

func CreateFile(path string, evaluatedTemplate *bytes.Buffer) error {

	err := os.WriteFile(path, evaluatedTemplate.Bytes(), 0644)
	if err != nil {
		log.Print("Error creating file:", err)
		os.Remove(path)
		return err
	}

	return nil
}
