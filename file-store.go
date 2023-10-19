package fileserver

import (
	"io"
	"mime/multipart"
	"os"
)

func fileStore(file multipart.File, upload_folder, file_name, extension string) error {

	err := os.MkdirAll(upload_folder, os.ModePerm)
	if err != nil {
		return err
	}

	dst, err := os.Create(upload_folder + "/" + file_name + extension)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}
