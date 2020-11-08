package adr

import (
	"fmt"
	"strings"
)

func DefaultTemplateContent(title string, status string) string {
	return fmt.Sprintf(`# %s

## Status

Status: %s

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?`, title, status)
}

func CreateMetaContent(parameters [] string) string {
	if len(parameters)>0 {
		valueSeparator := ": \"\"  \n"
		return fmt.Sprintf("---\n%s---", strings.Join(parameters, valueSeparator) + valueSeparator)
	}
	return ""
}

