package fileserver

import (
	"os"
	"path/filepath"
	"time"
)

func FindFilesWithNonZeroSize(dir string, filenames []string) (err string) {

	// Esperar
	time.Sleep(50 * time.Millisecond)

	// Crea un mapa para hacer un seguimiento de los archivos encontrados
	found := make(map[string]bool)
	for _, filename := range filenames {
		found[filename] = false
	}

	// Recorre el directorio en busca de archivos
	er := filepath.Walk(dir, func(path string, info os.FileInfo, er error) error {
		if er != nil {
			return er
		}

		// Comprueba si el archivo actual es uno de los que estamos buscando
		filename := filepath.Base(path)
		if _, ok := found[filename]; ok && info.Size() > 0 {
			found[filename] = true
		}

		return nil
	})

	if er != nil {
		return er.Error()
	}

	// Verifica que se encontraron todos los archivos y que tienen tamaño mayor que cero
	for filename, ok := range found {
		if !ok {
			return "no se encontró el archivo " + filename
		}

	}

	return ""
}

func FindFile(dir, file_name string) (content, err string) {

	files, err := FileCheck(dir, file_name)
	if err != "" {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() && file.Name() == file_name {

			file_path := filepath.Join(dir, file.Name())

			content, er := os.ReadFile(file_path)
			if er != nil {
				return "", er.Error()
			}
			return string(content), ""
		}
	}

	return "", "archivo " + file_name + " no encontrado"
}

func FindFirstFileWithExtension(dir, extension string) (content, err string) {

	files, err := FileCheck(dir, extension)
	if err != "" {
		return "", err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if extension == filepath.Ext(file.Name()) {
			file_path := filepath.Join(dir, file.Name())

			content, er := os.ReadFile(file_path)
			if er != nil {
				return "", er.Error()
			}
			return string(content), ""

		}

	}

	return "", "extension " + extension + " no encontrada"
}
