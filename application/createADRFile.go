package application

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
)

var findFilesInDir = adr.FindADRFilesInDir
var getLastIdFromDir = adr.GetLastIdFromDir
var createFilename = adr.CreateFilename
var writeFile = adr.WriteFile
var defaultTemplateContent = adr.DefaultTemplateContent
var createMetaContent = adr.CreateMetaContent

func CreateADRFile(title string, directory string, templateFile string, meta []string) (string, error) {
	files, filesSearchError := findFilesInDir(directory)
	if filesSearchError != nil {
		return "", fmt.Errorf("create file: error listing directory files in %s %s ", directory, filesSearchError)
	}
	ADRId      := getLastIdFromDir(files)
	NextId     := ADRId + 1
	fileName   := createFilename(NextId, title)

	var content string
	if templateFile != "" {
		templateContent, templateContentError := adr.GetTemplateFileContent(templateFile)
		if templateContentError != nil {
			return "",fmt.Errorf("create file: error reading template file %s %s ", templateFile, templateContentError)
		}
		content = templateContent
	}else{
		content = defaultTemplateContent(title)
	}

	if meta !=nil {
		content = content + createMetaContent(meta)
	}

	return writeFile(fileName, content)
}
