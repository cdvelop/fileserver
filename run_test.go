package fileserver_test

import (
	"log"
	"testing"

	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/fileinput"
	"github.com/cdvelop/fileserver"
	"github.com/cdvelop/gotools"
	"github.com/cdvelop/maps"
	"github.com/cdvelop/model"
	"github.com/cdvelop/sqlite"
	"github.com/cdvelop/testools"
)

const root_test_folder = "./root_test_folder"

func Test_CrudFILE(t *testing.T) {
	object := &model.Object{
		Name:   "patient",
		Table:  "patient",
		Module: &model.Module{ModuleName: "medical_test", Title: "Modulo Testing", Areas: []byte{'s'}},
	}

	var (
		testData = []dataTest{
			{
				Request: testools.Request{
					TestName: "gatito 220kb  solo un solo archivo ok",
					Method:   "POST",
					Endpoint: "upload",
					Object:   "foto_mascota",
					Expected: 1,
					Analysis: analysisTestCreateFileOK,
				},
				files:       []string{"gatito.jpeg"},
				FileSetting: filehandler.FileSetting{MaximumFilesAllowed: 1, MaximumKbSize: 220, FileType: "imagen", DescriptiveName: "foto_mascota", Source: object},
			},
			{
				Request: testools.Request{
					TestName: "crear 2 archivos gatito 220kb y dino 36kb",
					Method:   "POST",
					Endpoint: "upload",
					Object:   "endoscopia",
					Expected: 2,
					Analysis: analysisTestCreateFileOK,
				},
				files:       []string{"dino.png", "gatito.jpeg"},
				FileSetting: filehandler.FileSetting{MaximumFilesAllowed: 2, MaximumKbSize: 262, FileType: "imagen", DescriptiveName: "endoscopia", Source: object},
			},
			// {
			// 	Request: testools.Request{
			// 		TestName: "tamaño gatito 220kb y permitido 200 se espera error",
			// 		Method:   "POST",
			// 		Endpoint: "upload",
			// 		Object:   "gato_malo",
			// 		Expected: []map[string]string{{"error": "error tamaño de archivo excedido máximo admitido: 215040 kb"}},
			// 		Analysis: analysisTestCreateFileERROR,
			// 	},
			// 	files:       []string{"dino.png", "gatito.jpeg"},
			// 	FileSetting: filehandler.FileSetting{MaximumFilesAllowed: 1, MaximumKbSize: 200, FileType: "imagen", DescriptiveName: "gato_malo", Source: object},
			// },
		}
	)

	// DeleteUploadTestFiles
	err := gotools.DeleteIfFolderSizeExceeds(root_test_folder, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = gotools.CreateFolderIfNotExist(root_test_folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range testData {
		t.Run((r.TestName), func(t *testing.T) {

			h := &model.Handlers{
				FileRootFolder:  root_test_folder,
				DataBaseAdapter: sqlite.NewConnection(root_test_folder, "stored_files_index.db", false),
			}

			fs, err := fileserver.AddFileApi(h, h, h, root_test_folder)
			if err != nil {
				t.Fatal(err)
				return
			}
			h.FileApi = fs

			app, err := testools.NewApiTestDefault(t, h)
			if err != nil {
				t.Fatal(err)
				return
			}
			defer app.Server.Close()
			r.ApiTest = app

			// AGREGAR API FILE INPUT
			_, err = fileinput.NewUploadFileApi(h, r.Source, r.FileSetting)
			if err != nil {
				t.Fatal(err)
				return
			}

			//*** CREAR FORMULARIO PARA ENVIÓ
			new := filehandler.File{
				Module_name: r.Source.ModuleName,
				Field_name:  r.DescriptiveName,
				Object_id:   testools.RandomNumber(),
				File_area:   string(r.Source.Areas[0]),
				// Extension:   r.Extension,
				// Description: "",
			}

			form, err := maps.BuildFormString(&new)
			if err != nil {
				t.Fatal(err)
				return
			}

			body, boundary, err := fileserver.MultiPartFileForm(path_files, r.files, form)
			if err != nil {
				t.Fatal(err)
				return
			}

			send := map[string][]byte{
				boundary: body,
			}

			r.SendOneRequest(r.Method, app.BuildEndPoint(r.Request), r.Object, send, func(response []map[string]string, err error) {
				r.Analysis(&r.Request, response, err)
			})

			// fmt.Println("METHOD: ", body)

			// var code int
			// d.Endpoint += d.file.Object.Name

			// fmt.Println("ENDPOINT CREATE: ", d.Endpoint)

			// responses, _, err = d.SendRequest(body.Bytes())
			// if err != nil {
			// 	t.Fatal(err)
			// }

			// for _, resp := range responses {
			// 	testools.CheckTest(prueba, d.expected, resp.Action, resp)
			// }

			// log.Println("CREATE RESPUESTAS:")

			// for _, response := range responses {

			// if response.Action != "error" {

			// data.updateTest(response, t)

			// data.readFileTest(response, t)

			// data.readTest(response, t)

			// data.deleteTest(response, t)

			// }
			// }
		})
	}

}
