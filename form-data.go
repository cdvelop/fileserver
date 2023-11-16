package fileserver

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/strings"

	"github.com/cdvelop/filehandler"
)

func multipartFormDataFile(f *filehandler.FileSetting, r *http.Request, w http.ResponseWriter) (map[string]string, error) {

	r.Body = http.MaxBytesReader(w, r.Body, f.GetMaximumFileSize()) //ej: 220 KB
	err := r.ParseMultipartForm(f.GetMaximumFileSize())
	if err != nil {
		if strings.Contains(err.Error(), "multipart") != 0 {
			return nil, fmt.Errorf("CreateFile ParseMultipartForm %v", err)
		} else {
			return nil, fmt.Errorf("error tamaño de archivo excedido máximo admitido: %v kb", f.GetMaximumFileSize())
		}
	}

	params := make(map[string]string)
	for key, values := range r.PostForm {
		if len(values) > 1 {
			params[key] = strings.Join(values, ",")
		} else {
			params[key] = values[0]
		}
	}

	return params, nil
}
