package fileserver_test

import (
	"fmt"

	"github.com/cdvelop/testools"
)

func analysisTestCreateFileOK(r *testools.Request, resp []map[string]string, err error) {
	fmt.Println("ANALIZANDO:", r.TestName)
	fmt.Println("resp:", resp)

	if err != nil {
		r.T.Fatal(err)
		return
	}

	if len(resp) != r.Expected {
		r.CheckTest(r.Expected, len(resp), "SE ESPERA ", r.Expected, " RESPUESTA", resp)
	}

	if len(resp[0]) != 2 {
		r.CheckTest(1, len(resp[0]), "SE ESPERA 2 DATOS DESCRIPCIÓN Y ID", resp)
	}

}

func analysisTestCreateFileERROR(r *testools.Request, resp []map[string]string, err error) {
	r.CheckTest(r.Expected, resp)
}
