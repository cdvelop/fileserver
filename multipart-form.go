package fileserver

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/cdvelop/model"
)

// path_files ej: ./test_files
// "files" ej: "files", "endoscopia", "voucher", "foto_mascota", "foto_usuario"
// files_name ej: "gatito.jpg, perro.png"
func MultiPartFileForm(path_files string, x_files any, form map[string]string) (body []byte, boundary string, err error) {

	var files map[string]string

	if map_files, ok := x_files.(map[string]string); ok {
		files = map_files
	} else {
		return nil, "", model.Error("error MultiPartFileForm files map[string]string no ingresado")
	}

	body_buf := &bytes.Buffer{}
	writer := multipart.NewWriter(body_buf)

	for file_name, nominated_name := range files {

		// abrimos el archivo local para la prueba
		File, err := os.Open(filepath.Join(path_files, file_name))
		if err != nil {
			return nil, "", err
		}
		defer File.Close()

		if nominated_name != "" {
			file_name = nominated_name
		}

		// creamos formato de envió de archivo
		part, err := writer.CreateFormFile("files", file_name)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = io.Copy(part, File)
		if err != nil {
			return nil, "", err
		}

		// escribimos en memoria el campo del formulario
		err = writer.WriteField("files", file_name)
		if err != nil {
			return nil, "", err
		}
	}

	// Agregamos los parámetros al formulario
	for key, value := range form {

		err := writer.WriteField(key, value)
		if err != nil {
			return nil, "", err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return body_buf.Bytes(), writer.FormDataContentType(), nil
}

// https://matt.aimonetti.net/posts/2013-07-golang-multipart-File-upload-example/
