package fileserver

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/cdvelop/filehandler"
)

func fileStoreInHDD(file multipart.File, upload_folder string, f *filehandler.File) (err string) {
	const this = "fileStoreInHDD "
	e := os.MkdirAll(upload_folder, os.ModePerm)
	if e != nil {
		return this + e.Error()
	}

	dst, er := os.Create(upload_folder + "/" + f.Id_file + f.Extension)
	if er != nil {
		return this + er.Error()
	}
	defer dst.Close()

	_, er = io.Copy(dst, file)
	if er != nil {
		return this + er.Error()
	}

	return ""
}
