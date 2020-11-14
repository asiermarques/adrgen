package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/asiermarques/adrgen/domain"
	"github.com/asiermarques/adrgen/infrastructure"
)

func Test_ExecuteStatusCommand(t *testing.T) {
	directory, _ := os.Getwd()
	targetFile := filepath.Join(directory, "1-adr_file.md")
	testFiles := []string{targetFile}

	cleanTestFiles(testFiles)
	defer cleanTestFiles(testFiles)

	templateService := domain.CreateTemplateService(nil)

	infrastructure.WriteFile(
		targetFile,
		templateService.RenderDefaultTemplateContent(domain.TemplateData{
			Status: "proposed",
			Date:   "20-10-2020",
			Title:  "Test",
		}),
	)

	cmd := NewStatusChangeCmd()
	cmd.SetArgs([]string{"1", "accepted"})
	err := cmd.Execute()

	content, _ := infrastructure.GetFileContent(targetFile)
	if !strings.Contains(content, "Status: accepted") {
		t.Fatal(fmt.Sprintf("failed, status %s not found in file content %s", "accepted", content))
	}

	if err != nil {
		t.Fatal(fmt.Sprintf("failed, an error occurred: %s", err))
	}
}
