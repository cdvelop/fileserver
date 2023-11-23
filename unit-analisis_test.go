package fileserver_test

import (
	"github.com/cdvelop/testools"
)

func analysisCreateFileNameNoIdType(r *testools.Request, results []map[string]string, err string) {
	// fmt.Println("result:", results)

	if err != "" {
		if !r.CheckTest(r.Expected, err, "no se esperaba error") {
			r.T.Fatal()
			return
		}

	} else if len(results) != r.Expected {

		if !r.CheckTest(r.Expected, len(results), "SE ESPERABA", r.Expected, " SE OBTUVO", results) {
			r.T.Fatal()
			return
		}

	} else { // igual a lo esperado continuamos el análisis
		for _, data := range results {
			if len(data) != 2 {
				if !r.CheckTest(2, len(data), "SE ESPERA SOLO 2 DATOS DESCRIPCIÓN Y ID", data) {
					r.T.Fatal()
					return
				}
			}

			// aca test de lectura
			readFileTest(r, data)

			readFileJsonDataTest(r, data)
		}
	}
	// fmt.Println("paso análisis ok:")
}

func defaultTestAnalysis(r *testools.Request, resp []map[string]string, err error) {
	var response any
	if err != nil {
		response = err.Error()
	} else {
		response = resp
	}

	if !r.CheckTest(r.Expected, response) {
		r.Fatal()
		return
	}
}
