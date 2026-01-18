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

func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		log.Print("Error deleting file:", err)
		return err
	}
	return nil
}
