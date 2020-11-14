package domain

import (
	"fmt"
	"strings"
)

type TemplateData struct {
	Title  string
	Status string
	Date   string
	Meta   []string
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
	RenderCustomTemplateContent(data TemplateData) (string, error)
	RenderDefaultTemplateContent(data TemplateData) string
	RenderMetaContent(parameters []string) string
	RenderRelationLink(adr ADR, relationTitle string) string
}

type privateTemplateService struct {
	customTemplateContentReader CustomTemplateContentReader
}

func parseTemplateContent(data TemplateData, content string) string {
	content = strings.Replace(content, "{title}", data.Title, -1)
	content = strings.Replace(content, "{status}", data.Status, -1)
	content = strings.Replace(content, "{date}", data.Date, -1)
	return content
}

func (s privateTemplateService) RenderCustomTemplateContent(data TemplateData) (string, error) {
	content, err := s.customTemplateContentReader.Read()
	if err != nil {
		return "", err
	}

	if len(data.Meta) > 0 {
		content = s.RenderMetaContent(data.Meta) + "\n" + content
	}

	return parseTemplateContent(data, content), nil
}

func (s privateTemplateService) RenderDefaultTemplateContent(data TemplateData) string {
	content := ""
	if len(data.Meta) > 0 {
		content = s.RenderMetaContent(data.Meta)
	}
	content = content + "\n" + DEFAULT_TEMPLATE_CONTENT

	return parseTemplateContent(data, content)
}

func (s privateTemplateService) RenderMetaContent(parameters []string) string {
	if len(parameters) > 0 {
		valueSeparator := ": \"\"  \n"
		return fmt.Sprintf("---\n%s---\n", strings.Join(parameters, valueSeparator)+valueSeparator)
	}
	return ""
}

func (s privateTemplateService) RenderRelationLink(adr ADR, relationTitle string) string {
	adrTitle, err := adr.getTitleFromContent()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s [%s](%s)  ", relationTitle, adrTitle, adr.Filename().Value())
}



func CreateTemplateService(customTemplateContentReader CustomTemplateContentReader) TemplateService {
	return privateTemplateService{customTemplateContentReader}
}
