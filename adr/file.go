package adr

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/gosimple/slug"
)

// CreateFilename creates the ADR File filename
//
func CreateFilename(id int, title string, idDigits int) string {
	if idDigits < 1 {
		return fmt.Sprintf("%d-%s.md", id, slug.Make(title))
	}
	return fmt.Sprintf("%0"+strconv.Itoa(idDigits)+"d-%s.md", id, slug.Make(title))
}

// FindADRFilesInDir looks for ADR files in a give directory
//
func FindADRFilesInDir(dirname string) ([]string, error) {
	var result []string

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && ValidateADRFilename(file.Name()) {
			result = append(result, file.Name())
		}
	}
	return result, nil
}

// FindADRFileById looks for a ADR file by its id in a file list
//
func FindADRFileById(adrId int, files []string) (string, error) {
	re := regexp.MustCompile(`(?mi)^(\d+)-.+\.md`)
	for _, file := range files {
		matches := re.FindStringSubmatch(file)

		if len(matches) > 1 {
			idMatch, err := strconv.Atoi(matches[1])

			if err != nil {
				return "", err
			}

			if idMatch == adrId {
				return file, nil
			}
		}
	}
	return "", fmt.Errorf("file not found for ADR Id %d", adrId)
}

// ValidateADRFilename validates an ADR filename
//
func ValidateADRFilename(name string) bool {
	pattern := regexp.MustCompile(`(?mi)^\d+-.+\.md`)
	return pattern.MatchString(name)
}

// GetLastIdFromFilenames find the last ID from a file list
//
func GetLastIdFromFilenames(filenames []string) int {
	number := 0
	if len(filenames) > 0 {
		re := regexp.MustCompile("[0-9]+")
		for _, name := range filenames {
			current := re.FindAllString(name, 1)
			if len(current) > 0 {
				rawNumber := current[0]
				currentNumber, _ := strconv.Atoi(rawNumber)
				if currentNumber > number {
					number = currentNumber
				}
			}
		}
		return number
	}
	return number
}

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

// WriteFile writes a file and overwrites its content if it exists
//
func WriteFile(fileName string, data string) (string, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write([]byte(data))
	if err != nil {
		return "", err
	}

	return fileName, nil
}
