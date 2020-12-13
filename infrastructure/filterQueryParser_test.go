package infrastructure

import (
	"fmt"
	"testing"
)

func TestParseFilterParamsWithOneParam(t *testing.T) {
	query := "status=active"
	params, err := ParseFilterParams(query)
	if err != nil {
		t.Fatal(fmt.Sprintf("error parsing query filter %s", err))
	}

	if len(params["status"]) < 1 {
		t.Fatal(fmt.Sprintf("expected status param with value 'active' but status is not present"))
	}

	if params["status"][0] != "active" {
		t.Fatal(fmt.Sprintf("expected status param with value 'active', %s returned", params["status"] ))
	}

}

func TestParseFilterParamsWithEmptyQuery(t *testing.T) {
	query := ""
	params, err := ParseFilterParams(query)
	if err != nil {
		t.Fatal(fmt.Sprintf("error parsing query filter %s", err))
	}

	if len(params) >0 {
		t.Fatal(fmt.Sprintf("expected empty param list"))
	}
}

func TestParseFilterParamsWithMoreThanOne(t *testing.T) {
	query := "status=active&status=proposed&meta_variable=test"
	params, err := ParseFilterParams(query)
	if err != nil {
		t.Fatal(fmt.Sprintf("error parsing query filter %s", err))
	}

	if len(params["status"]) < 1 {
		t.Fatal(fmt.Sprintf("expected status param with value 'active' but status is not present"))
	}

	expected := map[string][]string{
		"status": {"active", "proposed"},
		"meta_variable": {"test"},
	}
	for mapKey, values  := range params {
		for key, value  := range values {
			if expected[mapKey][key] != value {
				t.Fatal(fmt.Sprintf("expected %s, returned %s", expected[mapKey][key], value))
			}
		}
	}
}