package cmd

import (
	"github.com/asiermarques/adrgen/domain"
	"os"
	"path/filepath"
	"testing"
)

func Test_ExecuteInitCommand(t *testing.T) {
	directory, _ := os.Getwd()
	testDirectory := filepath.Join(directory, "tests")
	expectedTemplateFile := filepath.Join(testDirectory, "adr", "adr_template.md")
	expectedConfigFile := filepath.Join(directory, domain.CONFIG_FILENAME+".yml")
	testFiles := []string{testDirectory, expectedTemplateFile, expectedConfigFile}

	cleanTestFiles(testFiles)
	defer cleanTestFiles(testFiles)

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
