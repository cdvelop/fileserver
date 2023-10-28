package fileserver

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/cdvelop/model"
)

func fileStore(file multipart.File, f *model.FileNewToStore) error {

	err := os.MkdirAll(f.UploadFolder, os.ModePerm)
	if err != nil {
		return err
	}

	dst, err := os.Create(f.UploadFolder + "/" + f.FileNameOnDisk + f.Extension)
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
