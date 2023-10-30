package fileserver

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

// path_files ej: ./test_files
// field_name ej: endoscopia, voucher, foto_mascota, foto_usuario
// files_name ej: "gatito.jpg, perro.png"
func MultiPartFileForm(path_files, field_name string, files_name []string, form map[string]string) (body *bytes.Buffer, content_type string, err error) {

	body = &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, file_name := range files_name {

		// abrimos el archivo local para la prueba
		File, err := os.Open(filepath.Join(path_files, file_name))
		if err != nil {
			return nil, "", err
		}
		defer File.Close()

		// creamos formato de envió de archivo
		part, err := writer.CreateFormFile(field_name, file_name)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = io.Copy(part, File)
		if err != nil {
			return nil, "", err
		}

		// escribimos en memoria el campo del formulario
		err = writer.WriteField(field_name, file_name)
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

	return body, writer.FormDataContentType(), nil
}

// https://matt.aimonetti.net/posts/2013-07-golang-multipart-File-upload-example/
