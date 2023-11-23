package fileserver

func (f fileServer) FilePath(params map[string]string) (file_path, file_area, err string) {
	// fmt.Println("parámetros FilePath recibidos: ", params)
	const this = "FilePath error "
	if len(params) != 1 {
		return "", "", this + "solo se puede recibir un parámetro para leer un archivo"
	}

	id, ok := params["id"]
	if !ok {
		return "", "", this + "parámetro id no enviado para leer archivo"
	}

	err = f.input_id.ValidateField(id, false)
	if err != "" {
		return "", "", this + err
	}

	data, err := f.ReadByID(id)
	if err != "" {
		return "", "", this + err
	}

	if len(data) != 1 {
		return "", "", this + "parámetros incorrectos al recuperar archivo"
	}

	file_path, file_area = f.BuildFilePath(data[0])
	// fmt.Println("AREA ARCHIVO: s, DB:" + file_area)
	// fmt.Println("PATH:" + file_path)

	return file_path, file_area, ""
}
