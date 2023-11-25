package fileserver

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
)

// ej files, .go, `Input: input\.(\w+)\(\)`
func GetNamesFromDirectoryExtensionAndPattern(dir, ext, pattern string) (out []string, err string) {

	files, err := FileCheck(dir, pattern)
	if err != "" {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// fmt.Println("ARCHIVO: ", file.Name())

		if ext == filepath.Ext(file.Name()) {

			file_path := filepath.Join(dir, file.Name())

			new_names, err := GetNamesFromFileAndPattern(file_path, pattern)
			if err != "" {
				return nil, err
			}

			for _, new_name := range new_names {
				var found bool
				for _, reg_name := range out {
					if new_name == reg_name {
						found = true
						break
					}
				}

				if !found {
					out = append(out, new_name)
				}

			}
		}
	}

	return out, ""
}

// ej files/file.txt, `Input: input\.(\w+)\(\)`
func GetNamesFromFileAndPattern(file_path, pattern string) (out []string, err string) {

	file, er := os.Open(file_path)
	if er != nil {
		return nil, er.Error()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputPattern := regexp.MustCompile(pattern)

	register := map[string]struct{}{}

	for scanner.Scan() {
		line := scanner.Text()
		match := inputPattern.FindStringSubmatch(line)

		if match != nil {
			if _, exist := register[match[1]]; !exist {
				out = append(out, match[1])
				register[match[1]] = struct{}{}
			}
		}
	}

	return out, ""
}
