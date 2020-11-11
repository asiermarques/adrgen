package application

import (
	"github.com/asiermarques/adrgen/domain"
)

// CreateADRFile is the application service for creating a new ADR file
//
func CreateADRFile(
	date string,
	title string,
	meta []string,
	config domain.Config,
	repository domain.ADRRepository,
	writer domain.ADRWriter,
	templateService domain.TemplateService,
) (string, error) {
	lastId := repository.GetLastId()
	ADRId := lastId + 1
	templateContentData := domain.TemplateData{
		Title:  title,
		Status: config.DefaultStatus,
		Date:   date,
		Meta:   meta,
	}

	var content string
	if config.TemplateFilename == "" {
		content = templateService.ParseDefaultTemplateContent(templateContentData)
	} else {
		_content, err := templateService.ParseCustomTemplateContent(templateContentData)
		if err != nil {
			return "", err
		}
		content = _content
	}

	adr := domain.ADR{
		Filename: domain.CreateADRFilename(ADRId, title, config.IdDigitNumber),
		Content:  content,
		ID:       ADRId,
		Status:   config.DefaultStatus,
	}
	err := writer.Persist(adr)
	return adr.Filename.Value(), err
}
