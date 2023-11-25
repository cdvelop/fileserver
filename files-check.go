package fileserver

import (
	"io/fs"
	"os"
)

func FileCheck(dir string, files_names ...string) (out []fs.DirEntry, err string) {

	if dir == "" {
		return nil, "el parámetro dir no pueden estar vacío"
	}

	if len(files_names) == 0 {
		return nil, "no hay nombre de archivos para revisar"
	} else {
		for _, file_name := range files_names {
			if file_name == "" {
				return nil, "el parámetro file_name no pueden estar vacío"
			}
		}
	}

	files, er := os.ReadDir(dir)
	if er != nil {
		return nil, er.Error()
	}

	return files, ""
}
