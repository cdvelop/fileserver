package fileserver_test

import "github.com/cdvelop/testools"

func crudTestAnalysis(r *testools.Request, resp_create []map[string]string, err error) {
	// create
	var response any
	if err != nil {
		response = err.Error()
	} else {
		response = resp_create
	}

	if !r.CheckTest(r.Expected, response) {
		r.Fatal()
		return
	}

	// update
	endpoint := r.Server.URL + "/update"

	// cambio el dato descripción a oso
	for _, v := range resp_create {
		v["description"] = "oso"
	}

	r.SendOneRequest("POST", endpoint, r.Object, resp_create, func(resp_update []map[string]string, err error) {
		if err != nil {
			r.Fatal("no se esperaba error al actualizar", err)
			return
		}
		// fmt.Println("RESPUESTA UPDATE:", resp_update)
	})

	// read
	endpoint = r.Server.URL + "/read"
	r.SendOneRequest("POST", endpoint, r.Object, resp_create, func(resp_read []map[string]string, err error) {
		if err != nil {
			r.Fatal("no se esperaba error al leer", err)
			return
		}
		// fmt.Println("RESPUESTA READ:", resp_read)

		for _, resp := range resp_read {
			if !r.CheckTest("oso", resp["description"], "EN READ") {
				r.Fatal()
				return
			}
		}
	})

	// delete
	endpoint = r.Server.URL + "/delete"
	r.SendOneRequest("POST", endpoint, r.Object, resp_create, func(resp_delete []map[string]string, err error) {
		if err != nil {
			r.Fatal("no se esperaba error al eliminar", err)
			return
		}

		if len(resp_delete) != len(resp_create) {
			r.Fatal("se esperaba que la cantidad eliminada:", len(resp_delete), " fuera igual a la solicitada:", len(resp_create))
			return
		}
		// fmt.Println("RESPUESTA ELIMINACIÓN:", resp_delete)
	})
}
