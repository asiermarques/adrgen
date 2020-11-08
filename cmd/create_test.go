package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
	"os"
	"path/filepath"
	"testing"
)

var getFileContent = adr.GetFileContent
var getDefaultTemplateFileContent = adr.DefaultTemplateContent

func Test_ExecuteCreateCommand(t *testing.T) {
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

func Test_ExecuteCreateCommandWithConfig(t *testing.T) {
	directory, _ := os.Getwd()
	configuredDirectory := "tests/adr"
	configuredDirectoryAbs := filepath.Join(directory, configuredDirectory)
	expectedFile := configuredDirectoryAbs + "/1-adr-title-c.md"
	expectedFile2 := configuredDirectoryAbs + "/2-adr-title2-c.md"
	expectedFile3WithMeta := configuredDirectoryAbs + "/3-adr-title-with-meta-c.md"
	testFiles := []string{
		expectedFile,
		expectedFile2,
		expectedFile3WithMeta,
		filepath.Join(configuredDirectoryAbs, "adr_template.md"),
		filepath.Join(directory, "tests"),
		filepath.Join(directory, "adrgen.config.yml"),
	}

	cleanTestFiles(testFiles)
	defer cleanTestFiles(testFiles)

	createConfigAndDirs(configuredDirectoryAbs, configuredDirectory)

	cmd := NewCreateCmd()
	cmd.SetArgs([]string{"ADR title C"})
	cmd.Execute()


	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatal("failed creating adr " + expectedFile)
	}


	cmd = NewCreateCmd()
	cmd.SetArgs([]string{"ADR title2 C"})
	cmd.Execute()

	if _, err := os.Stat(expectedFile2); os.IsNotExist(err) {
		t.Fatal("failed creating adr C")
	}

	cmd = NewCreateCmd()
	cmd.SetArgs([]string{"ADR title With Meta C"})
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
` + getDefaultTemplateFileContent("ADR title With Meta C")
	if content != expectdContent {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectdContent, content))
	}
}

func createConfigAndDirs(directoryToCreate string, directoryName string) {

	os.MkdirAll(directoryToCreate, os.ModePerm)
	adr.WriteFile(filepath.Join(directoryToCreate, "adr_template.md"), adr.DefaultTemplateContent("{title}"))
	adr.WriteFile("adrgen.config.yml", fmt.Sprintf(`default_meta: []
default_status: proposed
directory: %s
supported_statuses:
- proposed
- accepted
- rejected
- superseeded
- amended
- deprecated
template_file: %s/adr_template.md

`, directoryName, directoryName))

}

func cleanTestFiles(files []string) {
	for _, file := range files {
		os.RemoveAll(file)
	}
}
