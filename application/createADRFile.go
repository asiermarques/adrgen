package application

import (
	"fmt"
	"strconv"

	"github.com/asiermarques/adrgen/domain"
)

// CreateADRFile is the application service for creating a new ADR file
//
func CreateADRFile(
	date string,
	title string,
	meta []string,
	supersedesTargetADRId int,
	amendsTargetADRId int,
	config domain.Config,
	repository domain.ADRRepository,
	writer domain.ADRWriter,
	templateService domain.TemplateService,
	relationsManager domain.RelationsManager,
) (string, error) {
	lastId := repository.GetLastId()
	ADRId := lastId + 1
	templateContentData := domain.TemplateData{
		Title:  strconv.Itoa(ADRId) + ". " + title,
		Status: config.DefaultStatus,
		Date:   date,
		Meta:   meta,
	}

	var content string
	if config.TemplateFilename == "" {
		content = templateService.RenderDefaultTemplateContent(templateContentData)
	} else {
		_content, err := templateService.RenderCustomTemplateContent(templateContentData)
		if err != nil {
			return "", err
		}
		content = _content
	}

	adr, _ := domain.CreateADR(
		ADRId,
		content,
		domain.CreateADRFilename(ADRId, title, config.IdDigitNumber),
	)

	var relationError error
	if supersedesTargetADRId > 0 {
		adr, _, relationError = addRelation(
			adr,
			supersedesTargetADRId,
			"supersede",
			writer,
			relationsManager,
			repository,
		)
	}
	if amendsTargetADRId > 0 {
		adr, _, relationError = addRelation(
			adr,
			amendsTargetADRId,
			"amend",
			writer,
			relationsManager,
			repository,
		)
	}
	if relationError != nil {
		return "", relationError
	}

	err := writer.Persist(adr)
	if err != nil {
		return adr.Filename().Value(), err
	}

	return adr.Filename().Value(), err
}

func addRelation(
	adr domain.ADR,
	targetADRId int,
	relation string,
	writer domain.ADRWriter,
	relationsManager domain.RelationsManager,
	repository domain.ADRRepository) (domain.ADR, domain.ADR, error) {
	targetADR, err := repository.FindById(targetADRId)
	if err != nil {
		return adr, targetADR, fmt.Errorf(
			"error finding the target ADR, the ADR file was not created. %s",
			err,
		)
	}

	adr, targetADR, err = relationsManager.AddRelation(adr, targetADR, relation)
	if err != nil {
		return adr, targetADR, err
	}

	err = writer.Persist(targetADR)
	if err != nil {
		return adr, targetADR, err
	}

	return adr, targetADR, nil
}
