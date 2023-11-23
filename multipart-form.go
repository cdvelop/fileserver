package fileserver

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// path_files ej: ./test_files
// "files" ej: "files", "endoscopia", "voucher", "foto_mascota", "foto_usuario"
// files_name ej: "gatito.jpg, perro.png"
func MultiPartFileForm(path_files string, x_files any, form map[string]string) (body []byte, boundary string, err string) {
	const this = "MultiPartFileForm error "
	var files map[string]string

	if map_files, ok := x_files.(map[string]string); ok {
		files = map_files
	} else {
		return nil, "", this + "files map[string]string no ingresado"
	}

	body_buf := &bytes.Buffer{}
	writer := multipart.NewWriter(body_buf)

	for file_name, nominated_name := range files {

		// abrimos el archivo local para la prueba
		File, e := os.Open(filepath.Join(path_files, file_name))
		if e != nil {
			return nil, "", this + "Open " + e.Error()
		}
		defer File.Close()

		if nominated_name != "" {
			file_name = nominated_name
		}

		// creamos formato de envió de archivo
		part, e := writer.CreateFormFile("files", file_name)
		if e != nil {
			return nil, "", this + "CreateFormFile " + e.Error()
		}
		_, e = io.Copy(part, File)
		if e != nil {
			return nil, "", this + "Copy " + e.Error()
		}

		// escribimos en memoria el campo del formulario
		e = writer.WriteField("files", file_name)
		if e != nil {
			return nil, "", this + "WriteField files" + e.Error()
		}
	}

	// Agregamos los parámetros al formulario
	for key, value := range form {

		e := writer.WriteField(key, value)
		if e != nil {
			return nil, "", this + "WriteField form " + e.Error()
		}
	}

	e := writer.Close()
	if e != nil {
		return nil, "", this + "Close " + e.Error()
	}

	return body_buf.Bytes(), writer.FormDataContentType(), ""
}

// https://matt.aimonetti.net/posts/2013-07-golang-multipart-File-upload-example/
