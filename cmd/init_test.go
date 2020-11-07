package cmd

import (
	"github.com/asiermarques/adrgen/adr"
	"os"
	"path/filepath"
	"testing"
)


func Test_ExecuteInitCommand(t *testing.T) {
	directory, _ := os.Getwd()
	testDirectory := filepath.Join(directory, "tests")
	expectedTemplateFile := filepath.Join(testDirectory, "adr", "adr_template.md")
	expectedConfigFile := filepath.Join(directory, adr.CONFIG_FILENAME + ".yaml")
	testFiles := []string{testDirectory, expectedTemplateFile, expectedConfigFile}

	cleanInitTestFiles(testFiles)
	defer cleanInitTestFiles(testFiles)

	cmd := NewInitCmd()
	cmd.SetArgs([]string{"tests/adr"})
	cmd.Execute()


	if _, err := os.Stat(expectedTemplateFile); os.IsNotExist(err) {
		t.Fatal("failed creating template file " + expectedTemplateFile)
	}

	if _, err := os.Stat(expectedConfigFile); os.IsNotExist(err) {
		t.Fatal("failed creating config file " + expectedConfigFile)
	}
}

func cleanInitTestFiles(files []string) {
	for _, file := range files {
		os.RemoveAll(file)
	}
}
