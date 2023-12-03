package fileserver

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
)

// dir ej: modules/mymodule
// ext ej: .js, .html, .css
func ReadFiles(dir, ext string, buffer_out *bytes.Buffer) (err string) {
	const this = "ReadFiles error "
	er := filepath.Walk(dir, func(path string, info os.FileInfo, er error) error {
		if er != nil {
			return er
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ext {
			file, er := os.Open(path)
			if er != nil {
				return er
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				buffer_out.Write(scanner.Bytes())
				buffer_out.WriteString("\n")
			}
			if er := scanner.Err(); er != nil {
				return er
			}
		}
		return nil
	})

	if er != nil {
		return this + er.Error()
	}

	return ""
}

func FileGet(file string, buffer_out *bytes.Buffer) (err string) {
	const this = "FileGet error "

	// Leemos el contenido del archivo
	content, er := os.ReadFile(file)
	if er != nil {
		return this + er.Error()
	}
	// Escribimos el contenido en el buffer de salida
	_, er = buffer_out.Write(content)
	if er != nil {
		err = this + er.Error()
	}
	return
}

// ej: dir/files, .svg
func AddStringContendFromDirAndExtension(dir, ext string, out *string) (err string) {
	const this = "AddStringContendFromDirAndExtension error "

	files, er := os.ReadDir(dir)
	if er != nil {
		return this + er.Error()
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if ext == filepath.Ext(file.Name()) {
			file_path := filepath.Join(dir, file.Name())
			// Leemos el contenido del archivo
			content, er := os.ReadFile(file_path)
			if er != nil {
				return this + er.Error()
			}

			*out += string(content) + "\n"
		}
	}

	return ""
}

func AddStringContentFromFile(file_path string, out *string) (err string) {
	// Leemos el contenido del archivo
	content, er := os.ReadFile(file_path)
	if er == nil {
		*out += string(content) + "\n"
		return ""
	}

	return "AddStringContentFromFile error " + er.Error()
}
