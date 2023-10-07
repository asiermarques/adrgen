package _infrastructure

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal/adr"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
)

func extractIdFromADRFilename(filename string) (int, error) {
	re := regexp.MustCompile(`(?mi)^(\d+)-.+\.(md|adoc)`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) < 2 {
		return -1, fmt.Errorf("filename not valid %s", filename)
	}

	matchId, err := strconv.Atoi(matches[1])
	if err != nil {
		return -1, fmt.Errorf("error retrieving the id from file %s: %s", filename, err)
	}
	return matchId, nil
}

type privateADRDirectoryRepository struct {
	directory string
}

func (repo privateADRDirectoryRepository) FindAll() ([]adr.ADR, error) {
	var result []adr.ADR
	files, err := ioutil.ReadDir(repo.directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && adr.ValidateFilename(file.Name()) {
			id, _ := extractIdFromADRFilename(file.Name())
			content, _ := GetFileContent(filepath.Join(repo.directory, file.Name()))
			filename, _ := adr.CreateFilenameFromFilenameString(file.Name())
			adr, err := adr.CreateADR(id, content, filename)
			if err != nil {
				return nil, fmt.Errorf("file %s has not content", file.Name())
			}

			result = append(result, adr)
		}
	}
	return result, nil
}

func (repo privateADRDirectoryRepository) Query(filterParams map[string][]string) ([]adr.ADR, error) {
	allItems, err := repo.FindAll()
	if err != nil {
		return allItems, err
	}
	filteredItems := []adr.ADR{}
	for _, adr := range allItems {
		if FilterADR(adr, filterParams) == true {
			filteredItems = append(filteredItems, adr)
		}
	}
	return filteredItems, nil
}

func (repo privateADRDirectoryRepository) FindById(id int) (adr.ADR, error) {
	ADRs, err := repo.FindAll()
	if err != nil {
		return nil, err
	}
	for _, adr := range ADRs {
		if adr.ID() == id {
			return adr, nil
		}
	}
	return nil, fmt.Errorf("file not found for ADR Id %d", id)
}

func (repo privateADRDirectoryRepository) GetLastId() int {
	number := 0

	ADRs, _ := repo.FindAll()
	if len(ADRs) > 0 {
		for _, ADR := range ADRs {
			if ADR.ID() > number {
				number = ADR.ID()
			}
		}
	}
	return number
}

// CreateADRDirectoryRepository creates an instance of domain.Repository repository that finds ADRs in a directory
func CreateADRDirectoryRepository(directory string) adr.Repository {
	return privateADRDirectoryRepository{directory}
}

type privateFileADRWriter struct {
	directory string
}

// CreateFileADRWriter creates a domain.Writer instance that persist ADR Files in a directory
func CreateFileADRWriter(directory string) adr.Writer {
	return privateFileADRWriter{directory: directory}
}

func (w privateFileADRWriter) Persist(adr adr.ADR) error {
	return WriteFile(filepath.Join(w.directory, adr.Filename().Value()), adr.Content())
}
