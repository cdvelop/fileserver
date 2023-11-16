package fileserver

import (
	"os"
)

func (fileServer) FileGet(path string) (any, error) {

	archive, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return archive, nil
}

func (fileServer) FileDelete(path string) error {

	// fmt.Println("BORRANDO ARCHIVO EN EL SERVIDOR")
	// Borrar archivos desde hdd
	err := os.Remove(path)
	if err != nil {
		// fmt.Println("ERROR AL BORRAR ARCHIVO EN EL SERVIDOR", err)
		return err
	}

	return nil
}
