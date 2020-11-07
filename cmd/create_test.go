package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
	"os"
	"testing"
)

var getFileContent = adr.GetTemplateFileContent
var getDefaultTemplateFileContent = adr.DefaultTemplateContent

func Test_ExecuteCommand(t *testing.T) {
	directory, _ := os.Getwd()

	cmd := NewCreateCmd()
	cmd.SetArgs([]string{"ADR title"})
	cmd.Execute()
	expectedFile := directory + "/1-adr-title.md"

	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatal("failed creating adr " + expectedFile)
	}


	cmd = NewCreateCmd()
	cmd.SetArgs([]string{"ADR title2"})
	cmd.Execute()
	expectedFile2 := directory + "/2-adr-title2.md"

	if _, err := os.Stat(expectedFile2); os.IsNotExist(err) {
		t.Fatal("failed creating adr")
	}

	cmd = NewCreateCmd()
	cmd.SetArgs([]string{"ADR title With Meta"})
	cmd.Execute()
	expectedFile3WithMeta := directory + "/3-adr-title-with-meta.md"

	if _, err := os.Stat(expectedFile3WithMeta); os.IsNotExist(err) {
		t.Fatal("failed creating adr")
	}

	content, _ := getFileContent(expectedFile3WithMeta)
	expectdContent :=  `---
param1: ""
param2: ""
---
` + getDefaultTemplateFileContent("ADR title With Meta")
	if content != expectdContent {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectdContent, content))
	}

	os.RemoveAll(expectedFile)
	os.RemoveAll(expectedFile2)
	os.RemoveAll(expectedFile3WithMeta)
}
