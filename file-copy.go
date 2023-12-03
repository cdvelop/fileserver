package fileserver

import (
	"io"
	"os"
)

func CopyFile(src string, dest string) (err string) {
	const this = "CopyFile error "
	// Abrir el archivo origen
	srcFile, er := os.Open(src)
	if er != nil {
		return this + er.Error()
	}
	defer srcFile.Close()

	// Crear el archivo destino
	destFile, er := os.Create(dest)
	if er != nil {
		return this + er.Error()
	}
	defer destFile.Close()

	// Copiar el contenido del archivo origen al archivo destino
	_, er = io.Copy(destFile, srcFile)
	if er != nil {
		err = this + er.Error()
	}

	return
}
