package cmd

import (
	"os"
	"testing"
)


func Test_ExecuteInitCommand(t *testing.T) {
	directory, _ := os.Getwd()
	expectedFile := directory + "/1-adr-title.md"
	testFiles := []string{expectedFile}

	cleanInitTestFiles(testFiles)
	defer cleanInitTestFiles(testFiles)

	cmd := NewInitCmd()
	cmd.SetArgs([]string{"."})
	cmd.Execute()


	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatal("failed creating config file " + expectedFile)
	}
}

func cleanInitTestFiles(files []string) {
	for _, file := range files {
		os.RemoveAll(file)
	}
}
