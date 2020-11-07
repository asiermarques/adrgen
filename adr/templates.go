package adr

import (
	"fmt"
	"strings"
)

func DefaultTemplateContent(title string) string {
	return "# " + title + `

## Status

What is the status, such as proposed, accepted, rejected, deprecated, superseded, etc.?

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?`
}

func CreateMetaContent(parameters [] string) string {
	valueSeparator := ": \"\"  \n"
	return fmt.Sprintf("---\n%s---", strings.Join(parameters, valueSeparator) + valueSeparator)
}

