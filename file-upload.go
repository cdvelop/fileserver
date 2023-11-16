package fileserver

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/model"
)

func (f *fileServer) FileUpload(object_name, area_file string, request ...any) ([]map[string]string, error) {
	var r *http.Request
	var w http.ResponseWriter
	for _, v := range request {
		if rq, ok := v.(*http.Request); ok {
			r = rq
		}
		if wr, ok := v.(http.ResponseWriter); ok {
			w = wr
		}
	}

	if r == nil || w == nil {
		return nil, model.Error("r *http.Request o w http.ResponseWriter no enviado en FileUpload")
	}

	x, err := f.GetFileSettings(object_name)
	if err != nil {
		return nil, err
	}

	fmt.Println("CONFIGURACIONES:", x)

	form_data, err := multipartFormDataFile(x, r, w)
	if err != nil {
		return nil, err
	}

	fmt.Println("FORMULARIO ENVIADO:", form_data)

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		return nil, model.Error("error no hay archivos detectados en el formulario")
	}

	if len(files) > int(x.MaximumFilesAllowed) {
		return nil, model.Error("error se pretende subir", len(files), "archivos, pero el máximo permitido es:", x.MaximumFilesAllowed)
	}

	data_out := []map[string]string{}

	// verificar si el id del objeto fue enviado
	object_id, exist := form_data["object_id"]
	if !exist {
		return nil, model.Error("valor campo object_id:", x.FieldNameWithObjectID, "no enviado en formulario archivo")
	}

	upload_folder := f.UploadFolderPath(form_data)

	for _, fileHeader := range files {
		if fileHeader.Size > x.GetMaximumFileSize() {
			return nil, model.Error("error archivo(s) excede(n) el tamaño admitido de:", x.MaximumKbSize, "kb")
		}

		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		new := filehandler.File{
			Id_file:     f.GetNewID(),
			Module_name: x.Source.ModuleName,
			Field_name:  x.DescriptiveName,
			Object_id:   object_id,
			File_area:   area_file,
			Extension:   getExtensionOnly(fileHeader),
			Description: fileHeader.Filename,
		}

		var found_extension bool
		for _, ext := range x.GetAllowedExtensions() {
			if ext == new.Extension {
				found_extension = true
			}
		}

		if !found_extension {
			return nil, model.Error("extension archivo", new.Extension, "no valida como", x.DescriptiveName, "solo se admiten:", x.GetAllowedExtensions())
		}

		// obtengo extension original con punto
		new.Extension = filepath.Ext(fileHeader.Filename)

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}

		err = fileStoreInHDD(file, upload_folder, &new)
		if err != nil {
			return nil, err
		}

		out, err := f.FileRegisterInDB(&new)
		if err != nil {
			//borrar archivo creado en disco solo si corresponde
			err2 := os.Remove(upload_folder + "/" + new.Id_file + new.Extension)
			if err2 != nil {
				return nil, model.Error("FileRegisterInDB", err, err2)
			}
			return nil, err
		}

		data_out = append(data_out, out)
	}

	return data_out, nil
}
