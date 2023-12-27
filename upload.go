package fileserver

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/strings"
)

func (f *fileServer) FileUpload(object_name, area_file string, request ...any) (data_out []map[string]string, err string) {
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
		return nil, "r *http.Request o w http.ResponseWriter no enviado en FileUpload"
	}

	x, err := f.GetFileSettings(object_name)
	if err != "" {
		return nil, err
	}

	fmt.Println("CONFIGURACIONES:", x)

	form_data, err := multipartFormDataFile(x, r, w)
	if err != "" {
		return nil, err
	}

	fmt.Println("FORMULARIO ENVIADO:", form_data)

	files := r.MultipartForm.File["files"]

	total_files := len(files)

	if total_files == 0 {
		return nil, "error no hay archivos detectados en el formulario"
	}

	if total_files > int(x.MaximumFilesAllowed) {
		return nil, "se pretende subir " + strconv.Itoa(total_files) + " archivos, pero el máximo permitido es: " + strconv.FormatInt(x.MaximumFilesAllowed, 10)
	}

	data_out = []map[string]string{}

	// verificar si el id del objeto fue enviado
	object_id, exist := form_data["object_id"]
	if !exist {
		return nil, "valor campo object_id: " + x.FieldNameWithObjectID + " no enviado en formulario archivo"
	}

	for _, fileHeader := range files {
		if fileHeader.Size > x.GetMaximumFileSize() {
			return nil, "error archivo(s) excede(n) el tamaño admitido de: " + strconv.FormatInt(x.MaximumKbSize, 10) + "kb"
		}

		file, e := fileHeader.Open()
		if e != nil {
			return nil, "fileHeader " + e.Error()
		}
		defer file.Close()

		id, description := f.BuildIDFileNameAndDescription(fileHeader.Filename)
		// fmt.Println("NOMBRE DE ARCHIVO:", fileHeader.Filename)

		new := &filehandler.File{
			Id_file:     id,
			Module_name: x.ModuleName,
			Field_name:  x.DescriptiveName,
			Object_id:   object_id,
			File_area:   area_file,
			Extension:   getExtensionOnly(fileHeader),
			Description: description,
		}

		upload_folder := f.UploadFolderPath(new)

		var found_extension bool
		for _, ext := range x.AllowedExtensions {
			if ext == new.Extension {
				found_extension = true
			}
		}

		// fmt.Println("EXTENSION OBTENIDA:", new.Extension)

		if !found_extension {
			return nil, "extension archivo " + new.Extension + " no valida como " + x.DescriptiveName + " solo se admiten: " + strings.Join(x.AllowedExtensions, ",")
		}

		_, e = file.Seek(0, io.SeekStart)
		if e != nil {
			return nil, "file.Seek " + e.Error()
		}

		err = fileStoreInHDD(file, upload_folder, new)
		if err != "" {
			return nil, err
		}

		out, err := f.FileRegisterInDB(new)
		if err != "" {
			//borrar archivo creado en disco solo si corresponde
			err2 := os.Remove(upload_folder + "/" + new.Id_file + new.Extension)
			if err2 != nil {
				return nil, "FileRegisterInDB " + err + " " + err2.Error()
			}
			return nil, err
		}

		data_out = append(data_out, out)
	}

	return data_out, ""
}
