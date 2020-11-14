package domain

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

	templateService := CreateTemplateService(nil)
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
