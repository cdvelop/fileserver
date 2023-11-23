package fileserver

import "github.com/cdvelop/strings"

func (f fileServer) BuildIDFileNameAndDescription(header_name string) (id, description string) {

	// cortar el nombre del archivo para eliminar la extensión antes de almacenarlo
	// if len(t.Description) > 5 {
	// 	t.Description = t.Description[:len(t.Description)-len(t.Extension)]
	// }

	//1- si contiene un guion bajo _ viene con descripción el nombre de archivo
	// ej: 125552.1_gatito malo
	if strings.Contains(header_name, "_") == 1 {
		res := strings.Split(header_name, "_")
		if len(res) > 1 {
			id = res[0]
			description = res[1]
		}
	} else {
		// de lo contrario es un id normal
		id = header_name
		description = header_name
	}

	// 2- validamos si el id obtenido esta ok
	err := f.input_id.Validate.ValidateField(id, false)
	if err != "" {
		// si no es valido generamos un id nuevo
		id = f.GetNewID()
	}

	return

}
