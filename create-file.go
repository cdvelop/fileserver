package fileserver

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cdvelop/model"
)

func CreateFileInServer(r *http.Request, o *model.Object, user_area_file string, form_data map[string]string) ([]map[string]string, error) {

	f := o.ConfigFile()

	// fmt.Println("OBJETO: ", o)
	files := r.MultipartForm.File[f.InputNameWithFiles]
	if len(files) == 0 {
		return nil, model.Error("error no hay archivos detectados en el formulario")
	}

	if len(files) > int(f.MaximumFilesAllowed) {
		return nil, model.Error("error se pretende subir", len(files), "archivos, pero el máximo permitido es:", f.MaximumFilesAllowed)
	}

	data_out := []map[string]string{}

	upload_folder := o.UploadFolderPath(form_data)

	for _, fileHeader := range files {
		if fileHeader.Size > f.MaximumFileSize {
			return nil, model.Error("error archivo(s) excede(n) el tamaño admitido de:", f.MaximumKbSize, "kb")
		}

		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		new := model.FileNewToStore{
			DescriptionInputName: fileHeader.Filename,
			UploadFolder:         upload_folder,
			FileNameOnDisk:       o.GenerateFileNameOnDisk(),
			FileArea:             user_area_file,
			Extension:            getExtensionOnly(fileHeader),
		}

		if !strings.Contains(f.AllowedExtensions, new.Extension) {
			return nil, model.Error("formato de archivo no valido solo se admiten:", f.AllowedExtensions)
		}

		// obtengo extension original con punto
		new.Extension = filepath.Ext(fileHeader.Filename)

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}

		err = fileStore(file, &new)
		if err != nil {
			return nil, err
		}

		out, err := o.RegisterNewFile(&new, form_data)
		if err != nil {

			//borrar archivo creado en disco
			err2 := os.Remove(upload_folder + "/" + new.FileNameOnDisk)
			if err2 != nil {
				return nil, model.Error(err, err2)
			}

			return nil, err
		}

		data_out = append(data_out, out)
	}

	return data_out, nil
}
