package template

import (
	"fmt"
	"testing"
)

func TestCreateMetaContent(t *testing.T) {
	expectedString := `---
param1: ""
param2: ""
---
`

	templateService := CreateService(nil)
	result := templateService.RenderMetaContent([]string{"param1", "param2"})
	if expectedString != result {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectedString, result))
	}

	expectedString = ``
	result = templateService.RenderMetaContent([]string{})
	if expectedString != result {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectedString, result))
	}
}

func TestValidateTemplateFields(t *testing.T) {
	content := `# {title}
## Status
{status}

## Date
{date}

sddssd
`

	error := validateFields(content)
	if error != nil {
		t.Fatal(fmt.Sprintf("failed validating correct template: %s", error))
	}

	content2 := `# {title}
## Status
{status}

## Date
[date]

sddssd
`
	error = validateFields(content2)
	if error == nil {
		t.Fatal(fmt.Sprintf("failed validating incorrect template"))
	}

	errorExpected := "the configured template has not the required field {date}"
	if error.Error() != errorExpected {
		t.Fatal(fmt.Sprintf("failed, error expected '%s', returned '%s'", errorExpected, error))
	}

	content3 := `# {title}
Date: {date}
Status: {status}
`
	error = validateFields(content3)
	if error == nil {
		t.Fatal(fmt.Sprintf("failed validating incorrect template"))
	}

	errorExpected = "the configured template must have an status following the format \n\n## Status\n\n{status}\n\n"
	if error.Error() != errorExpected {
		t.Fatal(fmt.Sprintf("failed, error expected '%s', returned '%s'", errorExpected, error))
	}
}
