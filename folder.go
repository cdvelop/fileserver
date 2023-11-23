package fileserver

import (
	"os"
	"path/filepath"
)

func DeleteIfFolderSizeExceeds(folder_path string, maxSizeMB int64) (err string) {
	var size int64

	e := filepath.Walk(folder_path, func(path string, info os.FileInfo, er error) error {
		if er != nil {
			return er
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	if e != nil {
		return e.Error()
	}

	sizeMB := size / 1024 / 1024
	if sizeMB > maxSizeMB {
		e := os.RemoveAll(folder_path)
		if e != nil {
			return e.Error()
		}
	}

	return ""
}

func CreateFolderIfNotExist(folder_path string) (err string) {
	// Verificar si el directorio ya existe
	_, e := os.Stat(folder_path)
	if e == nil {
		// El directorio ya existe, no es necesario crearlo
		return ""
	}

	// Intentar crear el directorio
	e = os.MkdirAll(folder_path, os.ModePerm)
	if e != nil {
		// Ocurrió un error al crear el directorio
		return e.Error()
	}

	// Directorio creado exitosamente
	return ""
}

const DeleteEmptyFolderResponse = "directorio con archivos. no fue eliminado"

func DeleteEmptyFolder(folder_path string) (err string) {
	// fmt.Println("Verificar si el directorio existe")
	fileInfo, e := os.Stat(folder_path)
	if e != nil {
		// log.Println(err)
		// fmt.Println("directorio no existe nada que eliminar")
		return ""
	}

	// fmt.Println("Verificar si es un directorio")
	if !fileInfo.IsDir() {
		return os.ErrInvalid.Error()
	}

	// fmt.Println("Obtener la lista de archivos y directorios en el directorio")
	fileList, e := os.ReadDir(folder_path)
	if e != nil {
		return e.Error()
	}

	// fmt.Println("Verificar si el directorio está vacío")
	if len(fileList) == 0 {
		// fmt.Println("Eliminar el directorio vacío")
		e := os.Remove(folder_path)
		if e != nil {
			return e.Error()
		}
	} else {
		return DeleteEmptyFolderResponse
	}

	return ""
}
