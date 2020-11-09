package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
	"os"
	"path/filepath"
	"strings"
	"testing"
)


func Test_ExecuteStatusCommand(t *testing.T) {
	directory, _ := os.Getwd()
	targetFile := filepath.Join(directory, "1-adr_file.md")
	testFiles := []string{targetFile}

	cleanTestFiles(testFiles)
	defer cleanTestFiles(testFiles)

	adr.WriteFile(targetFile, adr.DefaultTemplateContent("11-09-2020","Test", "proposed"))

	cmd := NewStatusChangeCmd()
	cmd.SetArgs([]string{"1", "accepted"})
	err := cmd.Execute()

	content, _ := adr.GetFileContent(targetFile)
	if !strings.Contains(content, "Status: accepted") {
		t.Fatal(fmt.Sprintf("failed, status %s not found in file content %s", "accepted", content))
	}

	if err != nil {
		t.Fatal(fmt.Sprintf("failed, an error occurred: %s", err))
	}

}


