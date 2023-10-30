package fileserver_test

import (
	"github.com/cdvelop/fileinput"
	"github.com/cdvelop/testools"
)

type dataTest struct {
	file *fileinput.File

	pk_name string //ej: id_file

	field_name string   //ej: endoscopia, voucher, foto_mascota, foto_usuario
	files      []string //ej: "gatito.jpg, perro.png"
	file_type  string   //ej: imagen,video,document,pdf
	max_files  int64
	max_size   int64

	expected string

	*testools.Request
}
