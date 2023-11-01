package fileserver

import (
	"os"
)

func (FileServer) ReadFileFrom(path string) ([]byte, error) {

	archive, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return archive, nil
}
