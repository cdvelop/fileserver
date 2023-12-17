package fileserver

import (
	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/model"
	"github.com/cdvelop/unixid"
)

func AddFileApi(h *model.MainHandler) (fs *fileServer, err string) {

	fs = &fileServer{
		FileHandler: nil,
		input_id:    unixid.InputPK(),
	}
	h.FileApi = fs
	h.FileDiskRW = fs

	fh, err := filehandler.Add(h)
	if err != "" {
		return nil, err
	}
	fs.FileHandler = fh

	return fs, ""
}
