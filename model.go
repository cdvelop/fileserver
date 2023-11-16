package fileserver

import (
	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/model"
)

type fileServer struct {
	*filehandler.FileHandler

	input_id *model.Input
}
