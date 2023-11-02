package fileserver

import (
	"os"
)

func (FileServer) ReadFile(path string) ([]byte, error) {

	archive, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return archive, nil
}

func (FileServer) DeleteFile(path string) error {

	// Borrar archivos desde hdd
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
