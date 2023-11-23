package fileserver

import (
	"fmt"
	"io/fs"
	"os"
)

func FileCheck(dir string, files_names ...string) ([]fs.DirEntry, error) {

	if dir == "" {
		return nil, fmt.Errorf("el parámetro dir no pueden estar vacío")
	}

	if len(files_names) == 0 {
		return nil, fmt.Errorf("no hay nombre de archivos para revisar")
	} else {
		for _, file_name := range files_names {
			if file_name == "" {
				return nil, fmt.Errorf("el parámetro file_name no pueden estar vacío")
			}
		}
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return files, nil
}
