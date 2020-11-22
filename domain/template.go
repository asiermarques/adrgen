package domain

import (
	"fmt"
	"strings"
)

// TemplateData represents the data that will be render by the TemplateService
//
type TemplateData struct {
	Title  string
	Status string
	Date   string
	Meta   []string
}

// DEFAULT_TEMPLATE_CONTENT the default content for the template
//
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

// TemplateWriter persist the content of the template
//
type TemplateWriter interface {
	Persist() error
}

// CustomTemplateContentReader get the content of a template from a custom configured location
//
type CustomTemplateContentReader interface {
	Read() (string, error)
}

// TemplateService service that renders the template content
//
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

func validateTemplateFields(content string) error {
	fields := []string{"{title}", "{status}", "{date}"}
	for _, field := range fields {
		if !strings.Contains(content, field) {
			return fmt.Errorf("the configured template has not the required field %s", field)
		}
	}

	return nil
}

func (s privateTemplateService) RenderCustomTemplateContent(data TemplateData) (string, error) {
	content, err := s.customTemplateContentReader.Read()
	if err != nil {
		return "", err
	}

	if err := validateTemplateFields(content); err != nil {
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
		valueSeparator := ": \"\"\n"
		return fmt.Sprintf("---\n%s---\n", strings.Join(parameters, valueSeparator)+valueSeparator)
	}
	return ""
}

func (s privateTemplateService) RenderRelationLink(adr ADR, relationTitle string) string {
	return fmt.Sprintf("%s [%s](%s)  ", relationTitle, adr.Title(), adr.Filename().Value())
}

// CreateTemplateService creates a TemplateService instance
//
func CreateTemplateService(
	customTemplateContentReader CustomTemplateContentReader,
) TemplateService {
	return privateTemplateService{customTemplateContentReader}
}
