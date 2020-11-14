package infrastructure

import (
	"fmt"
	"github.com/asiermarques/adrgen/domain"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
)

func extractIdFromADRFilename(filename string) (int, error) {
	re := regexp.MustCompile(`(?mi)^(\d+)-.+\.md`)
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

func (repo privateADRDirectoryRepository) FindAll() ([]domain.ADR, error) {
	var result []domain.ADR
	files, err := ioutil.ReadDir(repo.directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && domain.ValidateADRFilename(file.Name()) {
			id, _ := extractIdFromADRFilename(file.Name())
			content, _ := GetFileContent(filepath.Join(repo.directory, file.Name()))
			filename, _ := domain.CreateADRFilenameFromFilenameString(file.Name())
			adr, err := domain.CreateADR(id, content, filename)
			if err != nil {
				return nil, fmt.Errorf("file %s has not content", file.Name())
			}

			result = append(result, adr)
		}
	}
	return result, nil
}

// FindADRFileById looks for a ADR file by its id in a file list
//
func (repo privateADRDirectoryRepository) FindById(id int) (domain.ADR, error) {
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

// GetLastIdFromFilenames find the last ID from a file list
//
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

func CreateADRRepository(directory string) domain.ADRRepository {
	return privateADRDirectoryRepository{directory}
}

type privateFileADRWriter struct {
	directory string
}

func CreateFileADRWriter(directory string) domain.ADRWriter {
	return privateFileADRWriter{directory: directory}
}

func (w privateFileADRWriter) Persist(adr domain.ADR) error {
	return WriteFile(filepath.Join(w.directory, adr.Filename().Value()), adr.Content())
}
