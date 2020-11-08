package application

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
	"path/filepath"
)

func CreateADRFile(title string, config adr.Config) (string, error) {
	files, filesSearchError := findFilesInDir(config.TargetDirectory)
	if filesSearchError != nil {
		return "", fmt.Errorf("create file: error listing directory files in %s %s ", config.TargetDirectory, filesSearchError)
	}
	ADRId := getLastIdFromFilenames(files)
	NextId := ADRId + 1
	fileName := createFilename(NextId, title)

	content := defaultTemplateContent(title, config.DefaultStatus)
	if config.MetaParams != nil && len(config.MetaParams) > 0 {
		content = createMetaContent(config.MetaParams) + "\n" + content
	}

	return writeFile(filepath.Join(config.TargetDirectory, fileName), content)
}
