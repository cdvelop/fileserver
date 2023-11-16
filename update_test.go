package fileserver_test

import (
	"testing"

	"github.com/cdvelop/model"
)

func (d *dataTest) updateTest(r model.Response, t *testing.T) {
	t.Run("UPDATE Test:", func(t *testing.T) {

		// fmt.Println("DATA A ACTUALIZAR: ", r)

		// 	for _, data := range r.Data {
		// 		data[d.file.Description] = "perro"
		// 	}

		// 	d.Endpoint = "/update/" + d.file.Object.Name

		// 	d.Data = r.Data

		// 	_, _, err := d.CutPost()
		// 	if err != nil {
		// 		t.Fatal(err)
		// 	}

		// 	d.Endpoint = "/read/" + d.file.Object.Name

		// 	responses, _, err := d.Get(r.Data...)
		// 	if err != nil {
		// 		t.Fatal(err)
		// 	}

		// 	// fmt.Println("RESPUESTA:", responses)

		// 	for i, resp := range responses {
		// 		testools.CheckTest("update", "update", resp.Action, resp)

		// 		if resp.Data[i][d.file.Description] != "perro" {
		// 			fmt.Printf("se esperaba en description [perro] se obtuvo: [%v]", resp.Data[i][d.file.Description])
		// 			t.Fatal()
		// 		}
		// 	}
	})
}
