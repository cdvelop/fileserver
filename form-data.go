package fileserver

import (
	"net/http"
	"strconv"

	"github.com/cdvelop/strings"

	"github.com/cdvelop/filehandler"
)

func multipartFormDataFile(f *filehandler.FileSetting, r *http.Request, w http.ResponseWriter) (params map[string]string, err string) {

	r.Body = http.MaxBytesReader(w, r.Body, f.GetMaximumFileSize()) //ej: 220 KB
	e := r.ParseMultipartForm(f.GetMaximumFileSize())
	if e != nil {
		if strings.Contains(e.Error(), "multipart") != 0 {
			return nil, "CreateFile ParseMultipartForm " + e.Error()
		} else {
			return nil, "tamaño de archivo excedido máximo admitido: " + strconv.FormatInt(f.GetMaximumFileSize(), 10) + " kb"
		}
	}

	params = make(map[string]string)
	for key, values := range r.PostForm {
		if len(values) > 1 {
			params[key] = strings.Join(values, ",")
		} else {
			params[key] = values[0]
		}
	}

	return params, ""
}
