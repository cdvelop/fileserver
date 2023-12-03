package fileserver

import (
	"bytes"
	"io"
	"os"
)

// file_name ej: "theme/index.html"
func FileWrite(file_name string, data *bytes.Buffer) (err string) {
	const this = "FileWrite error "
	dst, er := os.Create(file_name)
	if er != nil {
		return this + "al crear archivo " + er.Error()
	}
	defer dst.Close()

	// Copy the uploaded File to the filesystem at the specified destination
	_, er = io.Copy(dst, data)
	if er != nil {
		return this + "no se logro escribir el archivo " + file_name + " en el destino " + er.Error()
	}

	return ""
}
