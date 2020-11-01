package adr

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

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

func ValidateADRFilename(name string) bool  {
	pattern := regexp.MustCompile(`(?mi)^\d+-.+\.md`)
	return pattern.MatchString(name)
}

func GetLastIdFromDir(filenames []string) int {
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

func GetTemplateFileContent(templateFile string) (string, error) {
	file, err := os.Open(templateFile)
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

func WriteFile(fileName string, data string) (string, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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