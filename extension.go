package fileserver

import "mime/multipart"

func getExtension(fileHeader *multipart.FileHeader) string {
	buff := make([]byte, 512)
	file, err := fileHeader.Open()
	if err != nil {
		return ""
	}
	defer file.Close()

	_, err = file.Read(buff)
	if err != nil {
		return ""
	}

	return ArchiveType(&buff)
}
