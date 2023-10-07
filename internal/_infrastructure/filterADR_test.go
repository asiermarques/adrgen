package _infrastructure

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal/adr"
	"testing"
)

func TestFilterADR(t *testing.T) {
	filterParams := map[string][]string{
		"status": {"accepted", "proposed"},
	}
	adr1, _ := adr.CreateADR(1, "## Status\naccepted", adr.CreateFilename(1, "test", 2, ".md"))
	adr2, _ := adr.CreateADR(2, "## Status\nproposed", adr.CreateFilename(2, "test", 2, ".md"))
	adr3, _ := adr.CreateADR(3, "## Status\ndeprecated", adr.CreateFilename(3, "test", 2, ".md"))

	if FilterADR(adr1, filterParams) != true {
		t.Fatal(fmt.Sprintf("adr1 should be not filtered"))
	}
	if FilterADR(adr2, filterParams) != true {
		t.Fatal(fmt.Sprintf("adr2 should be not filtered"))
	}
	if FilterADR(adr3, filterParams) != false {
		t.Fatal(fmt.Sprintf("adr3 should be filtered"))
	}
	if FilterADR(adr1, map[string][]string{
		"status": {"proposed"},
	}) != false {
		t.Fatal(fmt.Sprintf("adr1 should be filtered"))
	}
}
