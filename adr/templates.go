package adr

import (
	"fmt"
	"strings"
)

// DefaultTemplateContent returns the content for the default ADR template
//
func DefaultTemplateContent(date string, title string, status string) string {
	return fmt.Sprintf(`# %s

Date: %s

## Status

Status: %s

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?`, title, date, status)
}

// CreateMetaContent creates the string that will be used as metadata in the ADR Markdown file
//
func CreateMetaContent(parameters []string) string {
	if len(parameters) > 0 {
		valueSeparator := ": \"\"  \n"
		return fmt.Sprintf("---\n%s---", strings.Join(parameters, valueSeparator)+valueSeparator)
	}
	return ""
}
