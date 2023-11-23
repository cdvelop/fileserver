package fileserver

import (
	"os"
	"path/filepath"
)

// ej: gotools.DeleteFilesByExtension(main_folder\files, []string{".js", ".css", ".wasm"})
func DeleteFilesByExtension(dir string, exts []string) (err string) {
	files, e := os.ReadDir(dir)
	if e != nil {
		return e.Error()
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		for _, ext := range exts {
			if ext == filepath.Ext(file.Name()) {
				path := filepath.Join(dir, file.Name())
				e := os.Remove(path)
				if e != nil {
					return "Error deleting file: " + e.Error()
				}
				break
			}
		}
	}

	return ""
}
