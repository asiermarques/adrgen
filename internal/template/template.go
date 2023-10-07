package template

import (
	"fmt"
	"regexp"
	"strings"
)

// TemplateData represents the data that will be render by the Service
type TemplateData struct {
	Title  string
	Status string
	Date   string
	Meta   []string
}

// DEFAULT_CONTENT the default content for the template
const DEFAULT_CONTENT = `# {title}

Date: {date}

## Status

{status}

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?`

// DEFAULT_ASCIIDOC_CONTENT the default content for the template in asciidoc format
const DEFAULT_ASCIIDOC_CONTENT = `= {title}

Date: {date}

== Status

{status}

== Context

What is the issue that we're seeing that is motivating this decision or change?

== Decision

What is the change that we're proposing and/or doing?

== Consequences

What becomes easier or more difficult to do because of this change?`

// Writer persist the content of the template
type Writer interface {
	Persist() error
}

// CustomContentReader get the content of a template from a custom configured location
type CustomContentReader interface {
	Read() (string, error)
}

// Service service that renders the template content
type Service interface {
	RenderCustomContent(data TemplateData) (string, error)
	RenderDefaultContent(data TemplateData) string
	RenderMetaContent(parameters []string) string
	RenderRelationLink(adrTitle string, adrFilename string, relationTitle string) string
}

type privateService struct {
	customTemplateContentReader CustomContentReader
}

func parseContent(data TemplateData, content string) string {
	content = strings.Replace(content, "{title}", data.Title, -1)
	content = strings.Replace(content, "{status}", data.Status, -1)
	content = strings.Replace(content, "{date}", data.Date, -1)
	return content
}

func validateFields(content string) error {
	fields := []string{"{title}", "{status}", "{date}"}
	for _, field := range fields {
		if !strings.Contains(content, field) {
			return fmt.Errorf("the configured template has not the required field %s", field)
		}
	}

	re := regexp.MustCompile(`(?mi)^(##|==) Status\n\n?(.+)$`)
	if !re.MatchString(content) {
		return fmt.Errorf(
			"the configured template must have an status following the format \n\n## Status\n\n{status}\n\n",
		)
	}

	return nil
}

func (s privateService) RenderCustomContent(data TemplateData) (string, error) {
	content, err := s.customTemplateContentReader.Read()
	if err != nil {
		return "", err
	}

	if err := validateFields(content); err != nil {
		return "", err
	}

	if len(data.Meta) > 0 {
		content = s.RenderMetaContent(data.Meta) + "\n" + content
	}

	return parseContent(data, content), nil
}

func (s privateService) RenderDefaultContent(data TemplateData) string {
	content := ""
	if len(data.Meta) > 0 {
		content = s.RenderMetaContent(data.Meta)
	}
	content = content + "\n" + DEFAULT_CONTENT

	return parseContent(data, content)
}

func (s privateService) RenderMetaContent(parameters []string) string {
	if len(parameters) > 0 {
		valueSeparator := ": \"\"\n"
		return fmt.Sprintf("---\n%s---\n", strings.Join(parameters, valueSeparator)+valueSeparator)
	}
	return ""
}

func (s privateService) RenderRelationLink(adrTitle string, adrFilename string, relationTitle string) string {
	return fmt.Sprintf("%s [%s](%s)  ", relationTitle, adrTitle, adrFilename)
}

// CreateService creates a Service instance
func CreateService(
	customTemplateContentReader CustomContentReader,
) Service {
	return privateService{customTemplateContentReader}
}
