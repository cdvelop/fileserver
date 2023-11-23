package fileserver_test

import (
	"log"
	"testing"

	"github.com/cdvelop/fileserver"
)

func TestFindFilesWithNonZeroSize(t *testing.T) {
	dir := "."                                   // Directorio actual
	filenames := []string{"README.md", "go.mod"} // Archivos que estamos buscando

	err := fileserver.FindFilesWithNonZeroSize(dir, filenames)
	if err != nil {
		t.Errorf("Se produjo un error inesperado: %v", err)
	}
}

func TestFindFile(t *testing.T) {
	var testData = map[string]struct {
		path   string
		search string
		expect string
	}{
		"archivo se espera contenido": {"test", "file.txt", "contenido"},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp, err := fileserver.FindFile(data.path, data.search)

			if resp != data.expect {
				t.Fatalf("Para la entrada '%s', '%s' se esperaba '%v' pero se obtuvo '%v' err: %v", data.path, data.search, data.expect, resp, err)
				log.Fatal()
			}
		})
	}
}

func TestFindFileByExtension(t *testing.T) {
	var testData = map[string]struct {
		path   string
		search string
		expect string
	}{
		"extension go.mod se espera contenido": {"test", ".mod", "contenido go mod único"},
		"extension .mo no existe":              {"test", ".mo", ""},
		"extension repetida .rep ":             {"test", ".rep", "contenido file 1"},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp, err := fileserver.FindFirstFileWithExtension(data.path, data.search)

			if resp != data.expect {
				t.Fatalf("Para la entrada '%s', '%s' se esperaba '%v' pero se obtuvo '%v' err: %v", data.path, data.search, data.expect, resp, err)
				log.Fatal()
			}
		})
	}
}
