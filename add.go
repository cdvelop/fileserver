package fileserver

import (
	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/model"
	"github.com/cdvelop/unixid"
)

func AddFileApi(l model.Logger, db model.DataBaseAdapter, hdd model.FileDiskRW, root_folder ...string) (*fileServer, error) {

	fh, err := filehandler.Add(l, db, hdd, root_folder...)
	if err != nil {
		return nil, err
	}

	fs := &fileServer{
		FileHandler: fh,
		input_id:    unixid.InputPK(),
	}

	return fs, nil
}
