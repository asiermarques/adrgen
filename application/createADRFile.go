package application

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/asiermarques/adrgen/adr"
)

func CreateADRFile(date string, title string, config adr.Config) (string, error) {
	files, filesSearchError := adr.FindADRFilesInDir(config.TargetDirectory)
	if filesSearchError != nil {
		return "", fmt.Errorf(
			"create file: error listing directory files in %s %s ",
			config.TargetDirectory,
			filesSearchError,
		)
	}
	ADRId := adr.GetLastIdFromFilenames(files)
	NextId := ADRId + 1
	fileName := adr.CreateFilename(NextId, title, config.IdDigitNumber)

	var content string
	if config.TemplateFilename == "" {
		content = adr.DefaultTemplateContent(date, title, config.DefaultStatus)
	} else {
		_content, err := createContentBodyFromTemplate(date, title, config.DefaultStatus, config.TemplateFilename)
		if err != nil {
			return "", fmt.Errorf("error creating ADR from template %s file: %s ", config.TemplateFilename, err)
		}
		content = _content
	}

	if config.MetaParams != nil && len(config.MetaParams) > 0 {
		content = adr.CreateMetaContent(config.MetaParams) + "\n" + content
	}

	return adr.WriteFile(filepath.Join(config.TargetDirectory, fileName), content)
}

func createContentBodyFromTemplate(
	date string,
	title string,
	status string,
	templateFile string,
) (string, error) {
	var content, err = adr.GetFileContent(templateFile)
	if err != nil {
		return "", err
	}
	content = strings.Replace(content, "{title}", title, -1)
	content = strings.Replace(content, "{status}", status, -1)
	content = strings.Replace(content, "{date}", date, -1)
	return content, nil
}
