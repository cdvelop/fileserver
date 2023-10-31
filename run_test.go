package fileserver_test

import (
	"log"
	"net/http/httptest"
	"testing"

	"github.com/cdvelop/api"
	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/fileinput"
	"github.com/cdvelop/gotools"
	"github.com/cdvelop/model"
	"github.com/cdvelop/sqlite"
	"github.com/cdvelop/testools"
)

var (
	dataHttp = map[string]dataTest{
		"crear 2 archivos gatito 220kb y dino 36kb": {field_name: "endoscopia", files: []string{"dino.png", "gatito.jpeg"}, file_type: "imagen", max_files: 2, max_size: 262, expected: "create", Request: &testools.Request{Endpoint: "/create/", Method: "POST"}},
		"gatito 220kb ok": {field_name: "foto_mascota", files: []string{"gatito.jpeg"}, file_type: "imagen", max_files: 1, max_size: 220, expected: "create", Request: &testools.Request{Endpoint: "/create/", Method: "POST"}},
		"tamaño gatito 220kb y permitido 200 se espera error": {field_name: "foto_mascota", files: []string{"gatito.jpeg"}, file_type: "imagen", max_files: 1, max_size: 200, expected: "error", Request: &testools.Request{Endpoint: "/create/", Method: "POST"}},
	}
)

const root_test_folder = "./root_test_folder"

func Test_CrudFILE(t *testing.T) {

	// DeleteUploadTestFiles
	err := gotools.DeleteIfFolderSizeExceeds(root_test_folder, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = gotools.CreateFolderIfNotExist(root_test_folder)
	if err != nil {
		log.Fatal(err)
	}

	for prueba, data := range dataHttp {
		t.Run((prueba), func(t *testing.T) {

			db := sqlite.NewConnection(root_test_folder, "stored_files_index.db", false)

			new_object := &model.Object{
				Name:                "name_test",
				Table:               "table_test",
				NamePrincipalFields: []string{},
				Fields:              []model.Field{},
				Module: &model.Module{
					ModuleName: "medical_history",
					Title:      "Modulo Testing",
					Areas:      []byte{},
					Objects:    []*model.Object{},
					Inputs:     []*model.Input{},
				},
				BackendHandler:  model.BackendHandler{},
				FrontendHandler: model.FrontendHandler{},
			}

			newConfig := model.FileConfig{
				MaximumFilesAllowed: data.max_files,
				InputNameWithFiles:  "",
				MaximumFileSize:     0,
				MaximumKbSize:       data.max_size,
				AllowedExtensions:   "",
				RootFolder:          root_test_folder,
				FileType:            "",
				IdFieldName:         "",
				Name:                data.field_name,
				Legend:              "",
			}
			h := model.Handlers{}

			data.file, err = fileinput.New(new_object, db, newConfig, &h)
			if err != nil {
				t.Fatal(err)
			}
			data.pk_name = data.file.Object.PrimaryKeyName()

			data.Object = data.file.Object
			// fmt.Println("OBJECT:", *data.Object)

			// data.Module.Objects = append(data.Module.Objects, data.Object)

			data.Cut = cutkey.Add(data.Object)

			api_conf := api.Add([]*model.Module{new_object.Module}, nil)

			mux := api_conf.ServeMuxAndRoutes()

			data.Server = httptest.NewServer(mux)
			defer data.Server.Close()

			responses := data.create(prueba, t)

			// log.Println("CREATE RESPUESTAS:", responses)

			for _, response := range responses {

				if response.Action != "error" {

					data.updateTest(response, t)

					data.readFileTest(response, t)

					// data.readTest(response, t)

					// data.deleteTest(response, t)

				}
			}
		})
	}

}
