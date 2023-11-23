package fileserver

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// file_name ej: "theme/index.html"
func FileWrite(file_name string, data *bytes.Buffer) error {

	dst, err := os.Create(file_name)
	if err != nil {
		return fmt.Errorf("error al crear archivo %v", err)
	}
	defer dst.Close()

	// Copy the uploaded File to the filesystem at the specified destination
	_, err = io.Copy(dst, data)
	if err != nil {
		return fmt.Errorf("error no se logro escribir el archivo %v en el destino %v", file_name, err)
	}

	return nil
}
