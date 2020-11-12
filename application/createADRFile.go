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

	if supersedesTargetADRId > 0 {
		_adr, _, err := supersedesADR(adr, supersedesTargetADRId, writer, relationsManager, repository)
		if err != nil {
			return "", err
		}
		adr = _adr
	}

	err := writer.Persist(adr)
	if err != nil {
		return adr.Filename.Value(), err
	}

	return adr.Filename.Value(), err
}

func supersedesADR(
	adr domain.ADR,
	targetADRId int,
	writer domain.ADRWriter,
	relationsManager domain.RelationsManager,
	repository domain.ADRRepository) (domain.ADR, domain.ADR, error)  {
	targetADR, err := repository.FindById(targetADRId)
	if err != nil {
		return adr, targetADR, fmt.Errorf("error finding the superseeded ADR, the ADR file was not created. %s", err)
	}

	adr, targetADR, err = relationsManager.PersistSupersedeOperation(adr, targetADR)
	if err != nil {
		return adr, targetADR, err
	}

	err = writer.Persist(targetADR)
	if err != nil {
		return adr, targetADR, err
	}

	return adr, targetADR, nil
}
