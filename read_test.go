package fileserver_test

import (
	"github.com/cdvelop/testools"
)

func readFileTest(r *testools.Request, all_data ...map[string]string) bool {

	for _, data := range all_data {
		// fmt.Println("DATA A ENVIAR PARA LECTURA ARCHIVO: ", data)

		endpoint := r.Server.URL + "/file?id=" + data["id_file"]

		r.SendOneRequest("GET", endpoint, "", nil, func(resp_read []map[string]string, err error) {

			if err != nil {
				if !r.CheckTest("", err.Error(), "no se esperaba error") {
					r.T.Fatal()
					return
				}
			}

			if !r.CheckTest(1, len(resp_read), "se esperaba un elemento leído") {
				r.T.Fatal()
				return
			}

		})
	}

	return true
}

func readFileJsonDataTest(r *testools.Request, all_data ...map[string]string) {
	const test_name = "readFileJsonDataTest"
	// var folders_ids []map[string]string
	var json_data []map[string]string
	for _, data := range all_data {
		// fmt.Println("DATA A ENVIAR PARA LECTURA JSON: ", data)

		endpoint := r.Server.URL + "/read/" + r.Object

		// fmt.Println("ENDPOINT:", endpoint)
		id := data["id_file"]
		// fmt.Println("ID:", id)

		// enviar solo id

		r.SendOneRequest("POST", endpoint, r.Object, map[string]string{"id_file": id}, func(resp_read []map[string]string, err error) {

			if err != nil {
				if !r.CheckTest("", err.Error(), "no se esperaba error", test_name) {
					r.T.Fatal()
					return
				}
			}

			if len(resp_read) != 1 {
				if !r.CheckTest(1, len(resp_read), "se esperaba solo un resultado", test_name, resp_read) {
					r.T.Fatal()
					return
				}

			}
			json_data = append(json_data, resp_read[0])

			// fmt.Println("--***--")

		})

	}
	// fmt.Println("*** DATA RECUPERADA:", json_data)

	if len(json_data) != len(all_data) {
		if !r.CheckTest(len(all_data), len(json_data), "se esperaba igual tamaño de resultado en información json recuperada", test_name, json_data) {
			r.T.Fatal()
			return
		}

	}

	// 	new_data := map[string]string{d.pk_name: data[d.pk_name]}

	// 	responses, code, err := d.Get(new_data)

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	// fmt.Println("RESPUESTA LECTURA JSON: ", responses)

	// 	for i, resp := range responses {

	// 		if id, folder_id := resp.Data[i][d.file.Object_id]; !folder_id {
	// 			log.Fatalln("se esperaba recuperar folder id")
	// 		} else {
	// 			folders_ids = append(folders_ids, map[string]string{d.file.Object_id: id})
	// 		}

	// 		testools.CheckTest("read data json", 200, code, resp)
	// 	}

	// t.Run("READ ALL BY FOLDER ID DATA JSON:", func(t *testing.T) {

	// fmt.Println("DATA FOLDER ID PARA LECTURA JSON: ", folders_ids)

	// for _, new_data := range folders_ids {
	// 	d.Endpoint = "/read/" + d.file.Object.Name

	// 	responses, code, err := d.Get(new_data)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if len(responses) != 1 {
	// 		log.Fatal("error se esperaba 1 respuesta se obtuvo: ", len(responses))
	// 	}

	// 	// fmt.Println("RESPUESTAS LECTURA FOLDER ID JSON: ", len(responses))

	// 	for _, resp := range responses {

	// 		// if _, folder_id := resp.Data[i]["folder_id"]; folder_id {
	// 		// 	log.Fatalln("error no se espera recibir nuevamente el dato folder_id")
	// 		// }

	// 		testools.CheckTest("read data json", 200, code, resp)
	// 	}
	// }
	// })

}
