package domain

import (
	"fmt"
	"strings"
)

type TemplateData struct {
	Title string
	Status string
	Date string
	Meta []string
}

const DEFAULT_TEMPLATE_CONTENT = `# {title}

Date: {date}

## Status

Status: {status}

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?`

type TemplateWriter interface {
	Persist() error
}

type CustomTemplateContentReader interface {
	Read() (string, error)
}

type TemplateService interface {
	ParseCustomTemplateContent(data TemplateData) (string, error)
	ParseDefaultTemplateContent(data TemplateData) string
	CreateMetaContent(parameters []string) string
}

type privateTemplateService struct {
	customTemplateContentReader CustomTemplateContentReader
}

func parseTemplateContent(data TemplateData, content string) string  {
	content = strings.Replace(content, "{title}", data.Title, -1)
	content = strings.Replace(content, "{status}", data.Status, -1)
	content = strings.Replace(content, "{date}", data.Date, -1)
	return content
}

func (s privateTemplateService) ParseCustomTemplateContent(data TemplateData) (string, error) {
	content, err := s.customTemplateContentReader.Read()
	if err != nil {
		return "", err
	}

	if len(data.Meta) > 0 {
		content = s.CreateMetaContent(data.Meta) + "\n" + content
	}

	return parseTemplateContent(data, content), nil
}

func (s privateTemplateService) ParseDefaultTemplateContent(data TemplateData) string {
	content := ""
	if len(data.Meta)>0 {
		content = s.CreateMetaContent(data.Meta)
	}
	content = content + "\n" + DEFAULT_TEMPLATE_CONTENT

	return parseTemplateContent(data, content)
}

func (s privateTemplateService) CreateMetaContent(parameters []string) string {
	if len(parameters) > 0 {
		valueSeparator := ": \"\"  \n"
		return fmt.Sprintf("---\n%s---\n", strings.Join(parameters, valueSeparator)+valueSeparator)
	}
	return ""
}

func CreateTemplateService(customTemplateContentReader CustomTemplateContentReader) TemplateService {
	return privateTemplateService{customTemplateContentReader}
}