package cmd

import (
	"os"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	directory, _ := os.Getwd()

	cmd := NewCreateCmd()
	cmd.SetArgs([]string{"ADR title"})
	cmd.Execute()
	expectedFile := directory + "/1-adr-title.md"

	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatal("failed creating adr")
	}

	cmd.SetArgs([]string{"ADR title2"})
	cmd.Execute()
	expectedFile2 := directory + "/2-adr-title2.md"

	if _, err := os.Stat(expectedFile2); os.IsNotExist(err) {
		t.Fatal("failed creating adr")
	}

	os.RemoveAll(expectedFile)
	os.RemoveAll(expectedFile2)
}
