package application

import (
	"../adr"
	"fmt"
)

func CreateADRFile(title string, directory string, templateFile string) error {
	files, filesSearchError := adr.FindADRFilesInDir(directory)
	if filesSearchError != nil {
		return fmt.Errorf("create file: error listing directory files in %s %s ", directory, filesSearchError)
	}
	ADRId      := adr.GetLastIdFromDir(files)
	fileName   := adr.CreateFilename(ADRId, title)

	content := adr.DefaultTemplateContent(title)
	if templateFile != "" {
		_content, templateContentError := adr.GetTemplateFileContent(templateFile)
		if templateContentError != nil {
			return fmt.Errorf("create file: error reading template file %s %s ", templateFile, templateContentError)
		}
		content = _content
	}

	return adr.WriteFile(fileName, content)
}
