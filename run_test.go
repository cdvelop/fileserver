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

type dataTest struct {
	testools.Request
	filehandler.FileSetting
}

const root_test_folder = "./root_test_folder"

func Test_CrudFILE(t *testing.T) {
	object := &model.Object{
		ObjectName: "patient",
		Table:      "patient",
		Module:     &model.Module{ModuleName: "medical_test", Title: "Modulo Testing", Areas: []byte{'s'}},
	}

	var (
		testData = []dataTest{
			{
				Request: testools.Request{
					TestName: "2 archivos de diferente clase se envían con id cada uno se espera ok",
					Method:   "POST",
					Endpoint: "upload",
					Object:   "foto_mascota",
					Data:     map[string]string{"dino.png": "1111.0_dino chico", "gatito.jpeg": "2222_gato grande"}, //con ids propuesto
					Expected: []map[string]string{{"description": "dino chico", "id_file": "1111.0"}, {"description": "gato grande", "id_file": "2222"}},
					Analysis: crudTestAnalysis,
				},
				FileSetting: filehandler.FileSetting{AllowedExtensions: []string{".jpeg", ".jpg", ".png"}, MaximumFilesAllowed: 2, MaximumKbSize: 270, DescriptiveName: "foto_mascota"},
			},
			{
				Request: testools.Request{
					TestName: "gatito 220kb solo un solo archivo nombre de fichero formato texto de debe crear id nuevo se espera ok",
					Method:   "POST",
					Endpoint: "upload",
					Object:   "foto_mascota",
					Data:     map[string]string{"gatito.jpeg": ""}, //sin id propuesto
					Expected: 1,
					Analysis: analysisCreateFileNameNoIdType,
				},
				FileSetting: filehandler.FileSetting{AllowedExtensions: []string{".jpeg", ".jpg"}, MaximumFilesAllowed: 1, MaximumKbSize: 220, DescriptiveName: "foto_mascota"},
			},
			{

				Request: testools.Request{
					TestName: "crear 2 archivos gatito 220kb y dino 36kb",
					Method:   "POST",
					Endpoint: "upload",
					Object:   "endoscopia",
					Data:     map[string]string{"dino.png": "", "gatito.jpeg": ""}, //sin id propuesto
					Expected: 2,
					Analysis: analysisCreateFileNameNoIdType,
				},
				FileSetting: filehandler.FileSetting{AllowedExtensions: []string{".jpeg", ".jpg", ".png"}, MaximumFilesAllowed: 2, MaximumKbSize: 262, DescriptiveName: "endoscopia"},
			},
			{

				Request: testools.Request{
					TestName: "tamaño gatito 220kb y permitido 200 se espera error",
					Method:   "POST",
					Endpoint: "upload",
					Object:   "gato_malo",
					Data:     map[string]string{"dino.png": "", "gatito.jpeg": ""}, //sin id propuesto
					Expected: "error tamaño de archivo excedido máximo admitido: 215040 kb",
					Analysis: analysisCreateFileNameNoIdType,
				},
				FileSetting: filehandler.FileSetting{AllowedExtensions: []string{".jpeg", ".jpg", ".png"}, MaximumFilesAllowed: 1, MaximumKbSize: 200, DescriptiveName: "gato_malo"},
			},
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

			app, err := testools.NewApiTestDefault(t, h)
			if err != nil {
				t.Fatal(err)
				return
			}
			defer app.Server.Close()
			r.ApiTest = app

			// AGREGAR API FILE INPUT
			_, err = fileinput.NewUploadFileApi(h, object, r.FileSetting)
			if err != nil {
				t.Fatal(err)
				return
			}

			//*** CREAR FORMULARIO PARA ENVIÓ
			new := filehandler.File{
				Module_name: object.ModuleName,
				Field_name:  r.DescriptiveName,
				Object_id:   testools.RandomNumber(),
				File_area:   string(object.Areas[0]),
			}

			form, err := maps.BuildFormString(&new)
			if err != nil {
				t.Fatal(err)
				return
			}

			body, boundary, err := fileserver.MultiPartFileForm(path_files, r.Data, form)
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

		})
	}

}
