package fileserver_test

import (
	"log"
	"testing"

	"github.com/cdvelop/fileserver"
)

func TestDeleteEmptyFolder(t *testing.T) {
	var testData = map[string]struct {
		dir    string
		expect string
	}{
		"eliminación normal directorio sin data":                {"test/folder01", ""},
		"intento de eliminación de una carpeta con archivo":     {"test/folder", fileserver.DeleteEmptyFolderResponse},
		"intento de eliminación de una carpeta con sub carpeta": {"test/folder0", fileserver.DeleteEmptyFolderResponse},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			err := fileserver.CreateFolderIfNotExist(data.dir)
			if err != "" {
				t.Fatal(err)
				return
			}

			err = fileserver.DeleteEmptyFolder(data.dir)
			if err != "" {
				if data.expect != err {
					log.Fatalf("\nPara la entrada '%s'\n-se esperaba [%s]\n-se obtuvo [%s]", data.dir, data.expect, err)
				}
			} else {
				if data.expect != "" {
					log.Fatalf("\nPara la entrada '%s'\n-se esperaba [%s]\n-no se obtuvo ningún error", data.dir, data.expect)
				}
			}

		})
	}
}
