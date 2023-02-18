package adr

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal/config"
	"github.com/asiermarques/adrgen/internal/template"
	"testing"
)

func TestGetTitleFromContent(t *testing.T) {
	contentStub := `
# My ADR Title
Date: 09-11-2020

## Status

accepted

## Context`

	adr, _ := CreateADR(1, contentStub, CreateFilename(1, "My ADR Title", 4))

	result := adr.Title()
	expectedString := "My ADR Title"
	if result != expectedString {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectedString, result))
	}

}

func TestRelation(t *testing.T) {
	contentStub := `
# My ADR Title
Date: 09-11-2020

## Status

accepted

## Context`

	targetContentStub := `
# My ADR Title
Date: 09-11-2020

## Status

accepted

## Context`

	relations := make(map[string]relation)
	relations["supersede"] = relation{
		mainTitle:    "Supersedes",
		targetTitle:  "Superseded by",
		targetStatus: "superseded",
	}
	relations["amend"] = relation{
		mainTitle:    "Amends",
		targetTitle:  "Amended by",
		targetStatus: "amended",
	}

	relationsManager := CreateRelationsManager(
		template.CreateService(nil),
		CreateStatusManager(config.Config{}),
	)

	for relationKey, relation := range relations {

		expectedContent := fmt.Sprintf(`
# My ADR Title
Date: 09-11-2020

## Status

accepted

%s [My ADR Title](0002-my-adr-title.md)  

## Context`, relation.mainTitle)

		expectedTargetContent := fmt.Sprintf(`
# My ADR Title
Date: 09-11-2020

## Status

%s

%s [My ADR Title](0001-my-adr-title.md)  

## Context`, relation.targetStatus, relation.targetTitle)

		adr, _ := CreateADR(1, contentStub, CreateFilename(1, "My ADR Title", 4))
		targetAdr, _ := CreateADR(1, targetContentStub, CreateFilename(2, "My ADR Title", 4))
		adr, targetAdr, err := relationsManager.AddRelation(adr, targetAdr, relationKey)
		if err != nil {
			t.Fatal(fmt.Sprintf("an unexpected error was returned %s", err))
		}
		if adr.Content() != expectedContent {
			t.Fatal(fmt.Sprintf("expected: %s, returned %s", expectedContent, adr.Content()))
		}
		if targetAdr.Content() != expectedTargetContent {
			t.Fatal(
				fmt.Sprintf(
					"expected: %s, returned %s",
					expectedTargetContent,
					targetAdr.Content(),
				),
			)
		}
	}
}

func TestCreateFilename(t *testing.T) {
	testFilename := func(expectedString string, id int, digitNumber int) {
		result := CreateFilename(id, "New ADR", digitNumber)
		if result.Value() != expectedString {
			t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectedString, result.Value()))
		}
	}

	testFilename("1-new-adr.md", 1, 0)
	testFilename("1-new-adr.md", 1, 1)
	testFilename("123-new-adr.md", 123, 1)
	testFilename("0123-new-adr.md", 123, 4)
	testFilename("00003-new-adr.md", 3, 5)
}

func TestValidateCorrectFilenames(t *testing.T) {
	for _, value := range []string{
		"something.md",
		"some_thing",
		"a0001-some_thing",
		"a0001-some_thing.md",
	} {
		if ValidateFilename(value) {
			t.Fatal("failed with wrong param: " + value)
		}
	}

	for _, value := range []string{
		"0001-something.md",
		"2-some_thing.md",
		"0-some_thing.md",
	} {
		if !ValidateFilename(value) {
			t.Fatal("failed with right param: " + value)
		}
	}
}

func TestChangeStatusInADRContent(t *testing.T) {
	contentStub := `
# Title
Date: 09-11-2020

## Status

proposed

## Context

What is the issue that we're seeing that is motivating this decision or change?
`

	expected := `
# Title
Date: 09-11-2020

## Status

accepted

## Context

What is the issue that we're seeing that is motivating this decision or change?
`

	statusManager := CreateStatusManager(config.Config{})
	_adr, _ := CreateADR(1, contentStub, CreateFilename(1, "Title", 4))
	adr, err := statusManager.ChangeStatus(_adr, "accepted")
	if err != nil {
		t.Fatal(fmt.Sprintf("%s", err))
	}
	if expected != adr.Content() {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expected, adr.Content()))
	}

	contentStub = `
# Title
Date: 09-11-2020

## Status

accepted

## Context

`
	expected = `
# Title
Date: 09-11-2020

## Status

rejected

## Context

`
	_adr, _ = CreateADR(1, contentStub, CreateFilename(1, "Title", 4))
	adr, err = statusManager.ChangeStatus(_adr, "rejected")
	if err != nil || expected != adr.Content() {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s :%s", expected, adr.Content(), err))
	}
}

func TestValidateStatus(t *testing.T) {
	allowedStatuses := []string{"status", "status2"}
	statusManager := CreateStatusManager(config.Config{Statuses: allowedStatuses})
	if statusManager.ValidateStatus("ñe") != false {
		t.Fatal(fmt.Sprintf("failed validating an incorrect status"))
	}

	statusManager = CreateStatusManager(config.Config{Statuses: []string{}})
	if statusManager.ValidateStatus("ñe") != true {
		t.Fatal(fmt.Sprintf("failed validating status when there is not any configured status"))
	}

	if statusManager.ValidateStatus("status") != true {
		t.Fatal(fmt.Sprintf("failed not validating a correct status"))
	}
}
