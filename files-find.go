package fileserver

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func FindFilesWithNonZeroSize(dir string, filenames []string) error {

	// Esperar
	time.Sleep(50 * time.Millisecond)

	// Crea un mapa para hacer un seguimiento de los archivos encontrados
	found := make(map[string]bool)
	for _, filename := range filenames {
		found[filename] = false
	}

	// Recorre el directorio en busca de archivos
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Comprueba si el archivo actual es uno de los que estamos buscando
		filename := filepath.Base(path)
		if _, ok := found[filename]; ok && info.Size() > 0 {
			found[filename] = true
		}

		return nil
	})

	if err != nil {
		return err
	}

	// Verifica que se encontraron todos los archivos y que tienen tamaño mayor que cero
	for filename, ok := range found {
		if !ok {
			return fmt.Errorf("no se encontró el archivo %s", filename)
		}

	}

	return nil
}

func FindFile(dir, file_name string) (content string, err error) {

	files, err := FileCheck(dir, file_name)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() && file.Name() == file_name {

			file_path := filepath.Join(dir, file.Name())

			content, err := os.ReadFile(file_path)
			if err != nil {
				return "", err
			}
			return string(content), nil
		}
	}

	return "", fmt.Errorf("archivo %v no encontrado", file_name)
}

func FindFirstFileWithExtension(dir, extension string) (content string, err error) {

	files, err := FileCheck(dir, extension)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if extension == filepath.Ext(file.Name()) {
			file_path := filepath.Join(dir, file.Name())

			content, err := os.ReadFile(file_path)
			if err != nil {
				return "", err
			}
			return string(content), nil

		}

	}

	return "", fmt.Errorf("extension %v no encontrada", extension)
}
