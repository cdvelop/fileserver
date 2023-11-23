package fileserver_test

import (
	"log"
	"testing"

	"github.com/cdvelop/fileserver"
	"github.com/cdvelop/model"
)

func TestGetInputNamesFromGoFile(t *testing.T) {

	var testData = map[string]struct {
		file_path string
		pattern   string
		expected  []string
	}{
		"búsqueda de inputs_1 se esperan 3": {"test/inputs_1.txt", model.INPUT_PATTERN, []string{"Phone", "RadioGenero", "TextArea"}},
		"inputs-rand se esperan n":          {"test/inputs-rand.txt2", model.INPUT_PATTERN, []string{"CheckBox", "Rut", "TextOnly", "DateAge", "TextArea", "RadioGenero"}},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp, err := fileserver.GetNamesFromFileAndPattern(data.file_path, data.pattern)

			if len(resp) != len(data.expected) {
				log.Fatalln("se esperaban ", len(data.expected), "elementos y se obtuvieron: ", len(resp), " ", resp)
			}

			// Comparar elemento por elemento
			for i, valor := range resp {
				if valor != data.expected[i] {
					log.Fatalf("se esperaba %v, pero se obtuvo %v\nerr: %v", data.expected[i], valor, err)
				}
			}

		})
	}
}

func TestGetNamesFromDirectoryExtensionAndPattern(t *testing.T) {
	var testData = map[string]struct {
		file_path string
		extension string
		pattern   string
		expected  []string
	}{
		"búsqueda en 2 archivos goo, Phone repetido": {"test", ".goo", model.INPUT_PATTERN, []string{"Phone", "RadioGenero", "TextArea"}},
		"búsqueda en 2 archivos gol, ":               {"test/golfiles", ".gol", model.INPUT_PATTERN, []string{"Text", "TextArea", "Date", "Hour"}},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp, err := fileserver.GetNamesFromDirectoryExtensionAndPattern(data.file_path, data.extension, data.pattern)

			if len(resp) != len(data.expected) {
				log.Fatalf("\n-se esperaba\n%v\n-pero se obtuvo\n%v", data.expected, resp)
			}
			// Comparar elemento por elemento
			for i, valor := range resp {
				if valor != data.expected[i] {
					log.Fatalf("se esperaba %v, pero se obtuvo %v\nresp:%v\nerr: %v", data.expected[i], valor, resp, err)
				}
			}

		})
	}
}
