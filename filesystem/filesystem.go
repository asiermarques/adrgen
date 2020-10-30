package filesystem

import (
	"io/ioutil"
	"log"
	"regexp"
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