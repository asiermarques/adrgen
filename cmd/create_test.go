package cmd

import (
	"os"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewCreateCmd()
	cmd.SetArgs([]string{"ADR title"})
	cmd.Execute()

	directory, _ := os.Getwd()
	if _, err := os.Stat(directory + "/1-adr-title.md"); os.IsNotExist(err) {
		t.Fatal("failed creating adr")
	}
}
