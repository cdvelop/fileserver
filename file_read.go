package fileserver

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
)

// dir ej: modules/mymodule
// ext ej: .js, .html, .css
func ReadFiles(dir, ext string, buffer_out *bytes.Buffer) (err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ext {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				buffer_out.Write(scanner.Bytes())
				buffer_out.WriteString("\n")
			}
			if err := scanner.Err(); err != nil {
				return err
			}
		}
		return nil
	})
	return
}

func FileGet(file string, buffer_out *bytes.Buffer) error {
	// Leemos el contenido del archivo
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	// Escribimos el contenido en el buffer de salida
	_, err = buffer_out.Write(content)
	if err != nil {
		return err
	}
	return nil
}

// ej: dir/files, .svg
func AddStringContendFromDirAndExtension(dir, ext string, out *string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if ext == filepath.Ext(file.Name()) {
			file_path := filepath.Join(dir, file.Name())
			// Leemos el contenido del archivo
			content, err := os.ReadFile(file_path)
			if err != nil {
				return err
			}

			*out += string(content) + "\n"
		}
	}

	return nil
}

func AddStringContentFromFile(file_path string, out *string) error {
	// Leemos el contenido del archivo
	content, err := os.ReadFile(file_path)
	if err == nil {
		*out += string(content) + "\n"
		return nil
	}

	return err
}
