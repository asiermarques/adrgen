package application

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
)

func CreateADRFile(title string, directory string, templateFile string) error {
	files, filesSearchError := adr.FindADRFilesInDir(directory)
	if filesSearchError != nil {
		return fmt.Errorf("create file: error listing directory files in %s %s ", directory, filesSearchError)
	}
	ADRId      := adr.GetLastIdFromDir(files)
	fileName   := adr.CreateFilename(ADRId + 1, title)

	var content string
	if templateFile != "" {
		templateContent, templateContentError := adr.GetTemplateFileContent(templateFile)
		if templateContentError != nil {
			return fmt.Errorf("create file: error reading template file %s %s ", templateFile, templateContentError)
		}
		content = templateContent
	}else{
		content = adr.DefaultTemplateContent(title)
	}


	return adr.WriteFile(fileName, content)
}
