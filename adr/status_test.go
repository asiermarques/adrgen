package adr

import (
	"fmt"
	"testing"
)

func TestChangeStatusInADRContent(t *testing.T) {
	contentStub := `
# Title
Date: 09-11-2020

## Status

Status: proposed

## Context

What is the issue that we're seeing that is motivating this decision or change?
`

	expected := `
# Title
Date: 09-11-2020

## Status

Status: accepted

## Context

What is the issue that we're seeing that is motivating this decision or change?
`

	content, err := ChangeStatusInADRContent("accepted", contentStub)
	if err != nil || expected != content {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s :%s", expected, content, err))
	}

	contentStub = `
# Title
Date: 09-11-2020

## Status

Status:accepted

## Context

`
	expected = `
# Title
Date: 09-11-2020

## Status

Status: rejected

## Context

`

	content, err = ChangeStatusInADRContent("rejected", contentStub)
	if err != nil || expected != content {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s :%s", expected, content, err))
	}
}

func TestValidateStatus(t *testing.T) {

	allowedStatuses := []string{"status", "status2"}
	if ValidateStatus("ñe", allowedStatuses) != false {
		t.Fatal(fmt.Sprintf("failed validating an incorrect status"))
	}
	if ValidateStatus("status", allowedStatuses) != true {
		t.Fatal(fmt.Sprintf("failed not validating a correct status"))
	}
}
