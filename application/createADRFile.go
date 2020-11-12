package application

import (
	"fmt"
	"github.com/asiermarques/adrgen/domain"
)

// CreateADRFile is the application service for creating a new ADR file
//
func CreateADRFile(
	date string,
	title string,
	meta []string,
	supersedesTargetADRId int,
	config domain.Config,
	repository domain.ADRRepository,
	writer domain.ADRWriter,
	templateService domain.TemplateService,
	relationsManager domain.RelationsManager,
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

	var targetADR domain.ADR
	if supersedesTargetADRId > 0 {
		_targetADR, err := repository.FindById(supersedesTargetADRId)
		if err != nil {
			return "", fmt.Errorf("error finding the superseeded ADR, the ADR file was not created. %s", err)
		}
		targetADR = _targetADR
	}

	adr, targetADR, err := relationsManager.PersistSupersedeOperation(adr, targetADR)
	if err != nil {
		return "", err
	}

	err = writer.Persist(targetADR)
	if err != nil {
		return "", err
	}

	err = writer.Persist(adr)
	if err != nil {
		return adr.Filename.Value(), err
	}

	return adr.Filename.Value(), err
}
