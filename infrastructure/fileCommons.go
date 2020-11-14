package infrastructure

import (
	"os"
)

// GetFileContent get the content of a file
//
func GetFileContent(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	buffer := make([]byte, fileinfo.Size())

	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

// WriteFile writes a file in disk, overwriting the previous content
//
func WriteFile(filename string, content string) error {
	file, err := os.OpenFile(
		filename,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}
