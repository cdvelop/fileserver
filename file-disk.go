package fileserver

import (
	"fmt"
	"os"
)

func (fileServer) FileGet(path string) (a any, err string) {

	archive, e := os.ReadFile(path)
	if e != nil {
		return nil, "FileGet error" + e.Error()
	}

	return archive, ""
}

func GetFile(path string) []byte {
	archive, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return archive
}

func (fileServer) FileDelete(path string) (err string) {

	// fmt.Println("BORRANDO ARCHIVO EN EL SERVIDOR:", path)
	// Borrar archivos desde hdd
	e := os.Remove(path)
	if e != nil {
		// fmt.Println("ERROR AL BORRAR ARCHIVO EN EL SERVIDOR", err)
		return e.Error()
	}
	// fmt.Println("BORRANDO ARCHIVO ok", e)

	return ""
}
