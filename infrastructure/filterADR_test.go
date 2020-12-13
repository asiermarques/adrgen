package infrastructure

import (
	"fmt"
	"github.com/asiermarques/adrgen/domain"
	"testing"
)

func TestFilterADR(t *testing.T) {
	filterParams := map[string][]string{
		"status": {"accepted", "proposed"},
	}
	adr1, _ := domain.CreateADR(1, "## Status\naccepted", domain.CreateADRFilename(1, "test", 2))
	adr2, _ := domain.CreateADR(2, "## Status\nproposed", domain.CreateADRFilename(2, "test", 2))
	adr3, _ := domain.CreateADR(3, "## Status\ndeprecated", domain.CreateADRFilename(3, "test", 2))

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