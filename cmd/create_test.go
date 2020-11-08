package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var getFileContent = adr.GetFileContent
var getDefaultTemplateFileContent = adr.DefaultTemplateContent

func assertCreateFile(key int, expectedFile string, cmd *cobra.Command, t *testing.T, meta []string) {
	cmd.SetArgs([]string{"ADR title " + string(key)})
	if meta != nil && len(meta)>0 {
		cmd.LocalFlags().Set("meta", strings.Join(meta, ","))
	}
	cmd.Execute()

	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatal("failed creating adr " + expectedFile)
	}
}

func Test_ExecuteCreateCommand(t *testing.T) {
	directory, _ := os.Getwd()
	testFiles := []string{
		filepath.Join(directory, "1-adr-title-0.md"),
		filepath.Join(directory, "1-adr-title-1.md"),
		filepath.Join(directory, "1-adr-title-2.md"),
	}

	cleanTestFiles(testFiles)
	defer cleanTestFiles(testFiles)

	for key, file := range testFiles {
		assertCreateFile(key, file, NewCreateCmd(), t, nil)
	}

	fileWithMeta := filepath.Join(directory, "1-adr-title-3.md")
	assertCreateFile(3, fileWithMeta, NewCreateCmd(), t, []string{"param1"," param2","param3"})

	content, _ := getFileContent(fileWithMeta)
	expectdContent :=  `---
param1: ""  
param2: ""  
param3: ""  
---
` + getDefaultTemplateFileContent("ADR title 3")
	if content != expectdContent {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectdContent, content))
	}
}

func Test_ExecuteCreateCommandWithConfig(t *testing.T) {
	directory, _ := os.Getwd()
	configuredDirectory := "tests/adr"
	configuredDirectoryAbs := filepath.Join(directory, configuredDirectory)
	testFiles := []string{
		filepath.Join(directory, "1-adr-title-0.md"),
		filepath.Join(directory, "1-adr-title-1.md"),
		filepath.Join(directory, "1-adr-title-2.md"),
	}
	testFilesAndDirs := append(testFiles,[]string{
		filepath.Join(configuredDirectoryAbs, "adr_template.md"),
		filepath.Join(directory, "tests"),
		filepath.Join(directory, "adrgen.config.yml"),
	}...)

	cleanTestFiles(testFilesAndDirs)
	defer cleanTestFiles(testFilesAndDirs)

	createConfigAndDirs(configuredDirectoryAbs, configuredDirectory)

	for key, file := range testFiles {
		assertCreateFile(key, file, NewCreateCmd(), t, nil)
	}

	fileWithMeta := filepath.Join(directory, "1-adr-title-3.md")
	assertCreateFile(3, fileWithMeta, NewCreateCmd(), t, []string{"param1"," param2","param3"})

	content, _ := getFileContent(fileWithMeta)
	expectdContent :=  `---
param1: ""  
param2: ""  
param3: ""  
---
` + getDefaultTemplateFileContent("ADR title 3")
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
