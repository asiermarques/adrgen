package application

import "fmt"

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
