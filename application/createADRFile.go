package application

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
)

var createFilename = adr.CreateFilename
var defaultTemplateContent = adr.DefaultTemplateContent
var createMetaContent = adr.CreateMetaContent

var getLastIdFromDir = adr.GetLastIdFromDir
var findFilesInDir = adr.FindADRFilesInDir
var readTemplateFileContent = adr.GetTemplateFileContent
var writeFile = adr.WriteFile

func CreateADRFile(title string, directory string, templateFile string, meta []string) (string, error) {
	files, filesSearchError := findFilesInDir(directory)
	if filesSearchError != nil {
		return "", fmt.Errorf("create file: error listing directory files in %s %s ", directory, filesSearchError)
	}
	ADRId      := getLastIdFromDir(files)
	NextId     := ADRId + 1
	fileName   := createFilename(NextId, title)

	var content string
	if meta != nil && len(meta) > 0 {
		content = createMetaContent(meta) + "\n" + defaultTemplateContent(title)
	}

	return writeFile(fileName, content)
}
