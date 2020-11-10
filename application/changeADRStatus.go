package application

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/asiermarques/adrgen/adr"
)

// ChangeADRStatus is the application service for changing the status in an ADR File
// It validates the status if there is a list of allowed statuses configured by user
//
func ChangeADRStatus(adrId int, status string, config adr.Config) (string, error) {
	if config.Statuses != nil && len(config.Statuses) > 0 {
		if !adr.ValidateStatus(status, config.Statuses) {
			return "", fmt.Errorf(
				"status %s not allowed, please use one of these %s",
				status,
				strings.Join(config.Statuses, ", "),
			)
		}
	}

	files, err := adr.FindADRFilesInDir(config.TargetDirectory)
	if err != nil {
		return "", fmt.Errorf("no ADR files found in dir %s", config.TargetDirectory)
	}

	file, err := adr.FindADRFileById(adrId, files)
	if err != nil {
		return "", fmt.Errorf("ADR with ID %d not found in dir %s", adrId, config.TargetDirectory)
	}

	content, err := adr.GetFileContent(filepath.Join(config.TargetDirectory, file))
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %s", file, err)
	}

	newContent, err := adr.ChangeStatusInADRContent(status, content)
	if err != nil {
		return "", fmt.Errorf("error changing status in file %s: %s", file, err)
	}

	_, err = adr.WriteFile(filepath.Join(config.TargetDirectory, file), newContent)
	return file, err
}
