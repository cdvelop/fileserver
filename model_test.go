package fileserver_test

import (
	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/testools"
)

type dataTest struct {
	testools.Request
	files []string //ej: "gatito.jpg, perro.png"
	filehandler.FileSetting

	// pk_name string //ej: id_file
}
