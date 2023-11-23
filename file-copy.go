package fileserver

import (
	"io"
	"os"
)

func CopyFile(src string, dest string) error {
	// Abrir el archivo origen
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Crear el archivo destino
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copiar el contenido del archivo origen al archivo destino
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
