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
	expectedFile := directory + "/1-adr-title.md"
	expectedFile2 := directory + "/2-adr-title2.md"
	expectedFile3WithMeta := directory + "/3-adr-title-with-meta.md"
	testFiles := []string{expectedFile, expectedFile2, expectedFile3WithMeta}

	cleanTestFiles(testFiles)
	defer cleanTestFiles(testFiles)

	cmd := NewCreateCmd()
	cmd.SetArgs([]string{"ADR title"})
	cmd.Execute()


	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatal("failed creating adr " + expectedFile)
	}


	cmd = NewCreateCmd()
	cmd.SetArgs([]string{"ADR title2"})
	cmd.Execute()

	if _, err := os.Stat(expectedFile2); os.IsNotExist(err) {
		t.Fatal("failed creating adr")
	}

	cmd = NewCreateCmd()
	cmd.SetArgs([]string{"ADR title With Meta"})
	cmd.LocalFlags().Set("meta", "param1, param2,param3")
	cmd.Execute()


	if _, err := os.Stat(expectedFile3WithMeta); os.IsNotExist(err) {
		t.Fatal("failed creating adr")
	}

	content, _ := getFileContent(expectedFile3WithMeta)
	expectdContent :=  `---
param1: ""
param2: ""
param3: ""
---
` + getDefaultTemplateFileContent("ADR title With Meta")
	if content != expectdContent {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectdContent, content))
	}



}

func cleanTestFiles(files []string) {
	for _, file := range files {
		os.RemoveAll(file)
	}
}
