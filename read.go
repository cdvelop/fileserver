package fileserver

import (
	"github.com/cdvelop/model"
)

func (f fileServer) FilePath(params map[string]string) (file_path, file_area string, err error) {
	// fmt.Println("parámetros FilePath recibidos: ", params)

	if len(params) != 1 {
		return "", "", model.Error("solo se puede recibir un parámetro para leer un archivo")
	}

	id, ok := params["id"]
	if !ok {
		return "", "", model.Error("parámetro id no enviado para leer archivo")
	}

	err = f.input_id.ValidateField(id, false)
	if err != nil {
		return "", "", err
	}

	data, err := f.ReadByID(id)
	if err != nil {
		return "", "", err
	}

	if len(data) != 1 {
		return "", "", model.Error("parámetros incorrectos al recuperar archivo")
	}

	file_path, file_area = f.BuildFilePath(data[0])
	// fmt.Println("AREA ARCHIVO: s, DB:" + file_area)
	// fmt.Println("PATH:" + file_path)

	return file_path, file_area, nil
}
