package adr

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func FindADRFilesInDir(dirname string) []string {

	var result = []string{}

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() && ValidateADRFilename(file.Name()) {
			result = append(result, file.Name())
		}
	}
	return result
}

func ValidateADRFilename(name string) bool  {
	var pattern = regexp.MustCompile(`(?mi)^\d+-.+\.md`)
	return pattern.MatchString(name)
}

func GetLastIdFromDir(filenames []string) int {
	var number = 0
	if len(filenames) > 0 {
		var re = regexp.MustCompile("[0-9]+")
		for _, name := range filenames {
			var current = re.FindAllString(name, 1)
			if len(current) > 0 {
				var rawNumber string = current[0]
				var currentNumber, _ = strconv.Atoi(rawNumber)
				if currentNumber > number {
					number = currentNumber
				}
			}
		}
		return number
	}
	return number
}