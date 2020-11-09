package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/asiermarques/adrgen/adr"
	"github.com/spf13/cobra"
)

var getFileContent = adr.GetFileContent
var getDefaultTemplateFileContent = adr.DefaultTemplateContent

func assertCreateFile(
	key int,
	expectedFile string,
	cmd *cobra.Command,
	t *testing.T,
	meta []string,
) {
	cmd.SetArgs([]string{"ADR title " + fmt.Sprint(key)})
	if meta != nil && len(meta) > 0 {
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
		filepath.Join(directory, "0001-adr-title-0.md"),
		filepath.Join(directory, "0002-adr-title-1.md"),
		filepath.Join(directory, "0003-adr-title-2.md"),
	}
	fileWithMeta := filepath.Join(directory, "0004-adr-title-3.md")
	testFilesAndDirs := append(testFiles, []string{
		fileWithMeta,
	}...)

	cleanTestFiles(testFilesAndDirs)
	defer cleanTestFiles(testFilesAndDirs)

	cleanTestFiles(testFiles)
	defer cleanTestFiles(testFiles)

	for key, file := range testFiles {
		assertCreateFile(key, file, NewCreateCmd(), t, nil)
	}

	assertCreateFile(3, fileWithMeta, NewCreateCmd(), t, []string{"param1", " param2", "param3"})

	currentTime := time.Now()
	date := currentTime.Format("02-01-2006")

	content, _ := getFileContent(fileWithMeta)
	expectdContent := `---
param1: ""  
param2: ""  
param3: ""  
---
` + getDefaultTemplateFileContent(date, "ADR title 3", "proposed")
	if content != expectdContent {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectdContent, content))
	}
}

func Test_ExecuteCreateCommandWithConfig(t *testing.T) {
	directory, _ := os.Getwd()
	configuredDirectory := "tests/adr"
	configuredDirectoryAbs := filepath.Join(directory, configuredDirectory)
	testFiles := []string{
		filepath.Join(configuredDirectoryAbs, "1-adr-title-0.md"),
		filepath.Join(configuredDirectoryAbs, "2-adr-title-1.md"),
		filepath.Join(configuredDirectoryAbs, "3-adr-title-2.md"),
	}
	fileWithMeta := filepath.Join(configuredDirectoryAbs, "4-adr-title-3.md")
	testFilesAndDirs := append(testFiles, []string{
		fileWithMeta,
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

	assertCreateFile(3, fileWithMeta, NewCreateCmd(), t, []string{"param1", " param2", "param3"})

	currentTime := time.Now()
	date := currentTime.Format("02-01-2006")

	content, _ := getFileContent(fileWithMeta)
	expectdContent := `---
param1: ""  
param2: ""  
param3: ""  
---
` + getDefaultTemplateFileContent(date, "ADR title 3", "proposed")
	if content != expectdContent {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectdContent, content))
	}
}

func createConfigAndDirs(directoryToCreate string, directoryName string) {

	os.MkdirAll(directoryToCreate, os.ModePerm)
	adr.WriteFile(
		filepath.Join(directoryToCreate, "adr_template.md"),
		adr.DefaultTemplateContent("{date}", "{title}", "{status}"),
	)
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
id_digit_number: 0

`, directoryName, directoryName))

}
