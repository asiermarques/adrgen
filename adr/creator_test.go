package adr

import (
	"fmt"
	"testing"
)

func TestCreateFilename(t *testing.T) {
	var expectedString = "1-new-adr.md"
	var result = CreateFilename(1, "New ADR")
	if result != expectedString {
		t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectedString, result))
	}
}