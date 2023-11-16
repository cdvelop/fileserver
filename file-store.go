package fileserver

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/cdvelop/filehandler"
)

func fileStoreInHDD(file multipart.File, upload_folder string, f *filehandler.File) error {

	err := os.MkdirAll(upload_folder, os.ModePerm)
	if err != nil {
		return err
	}

	dst, err := os.Create(upload_folder + "/" + f.Id_file + f.Extension)
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
