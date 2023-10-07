package adr

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal/config"
	"github.com/asiermarques/adrgen/internal/template"
	"path"
	"strconv"
)

// CreateFile is the application service for creating a new ADR file
func CreateFile(
	date string,
	title string,
	meta []string,
	supersedesTargetADRId int,
	amendsTargetADRId int,
	config config.Config,
	repository Repository,
	writer Writer,
	templateService template.Service,
	relationsManager RelationsManager,
) (string, error) {
	lastId := repository.GetLastId()
	ADRId := lastId + 1
	templateContentData := template.TemplateData{
		Title:  strconv.Itoa(ADRId) + ". " + title,
		Status: config.DefaultStatus,
		Date:   date,
		Meta:   meta,
	}

	var content string
	if config.TemplateFilename == "" {
		content = templateService.RenderDefaultContent(templateContentData)
	} else {
		_content, err := templateService.RenderCustomContent(templateContentData)
		if err != nil {
			return "", err
		}
		content = _content
	}

	extension := path.Ext(config.TemplateFilename)
	if extension == "" {
		extension = ".md"
	}

	adr, _ := CreateADR(
		ADRId,
		content,
		CreateFilename(ADRId, title, config.IdDigitNumber, extension),
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
	adr ADR,
	targetADRId int,
	relation string,
	writer Writer,
	relationsManager RelationsManager,
	repository Repository) (ADR, ADR, error) {
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
