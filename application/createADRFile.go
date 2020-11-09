package application

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
	"path/filepath"
	"strings"
)

func CreateADRFile(date string, title string, config adr.Config) (string, error) {
	files, filesSearchError := findFilesInDir(config.TargetDirectory)
	if filesSearchError != nil {
		return "", fmt.Errorf("create file: error listing directory files in %s %s ", config.TargetDirectory, filesSearchError)
	}
	ADRId := getLastIdFromFilenames(files)
	NextId := ADRId + 1
	fileName := createFilename(NextId, title)

	var content string
	if config.TemplateFilename == "" {
		content = defaultTemplateContent(date, title, config.DefaultStatus)
	} else {
		_content, err := createContentBodyFromTemplate(date, title, config.DefaultStatus, config.TemplateFilename)
		if err != nil {
			return "", fmt.Errorf("error creating ADR from template %s file: %s ", config.TemplateFilename, err)
		}
		content = _content
	}

	if config.MetaParams != nil && len(config.MetaParams) > 0 {
		content = createMetaContent(config.MetaParams) + "\n" + content
	}

	return writeFile(filepath.Join(config.TargetDirectory, fileName), content)
}

func createContentBodyFromTemplate(date string, title string, status string, templateFile string) (string, error) {
	var content, err = adr.GetFileContent(templateFile)
	if err != nil {
		return "", err
	}
	content = strings.Replace(content, "{title}", title, -1)
	content = strings.Replace(content, "{status}", status, -1)
	content = strings.Replace(content, "{date}", date, -1)
	return content, nil
}
