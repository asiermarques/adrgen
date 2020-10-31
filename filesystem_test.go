package main

import (
	"strconv"
	"testing"
)

func TestValidateCorrectFilenames(t *testing.T) {

	for _, value := range []string{
		"something.md",
		"some_thing",
		"a0001-some_thing",
		"a0001-some_thing.md",
	} {
		if ValidateADRFilename(value) {
			t.Fatal("failed with wrong param: " + value)
		}
	}

	for _, value := range []string{
		"0001-something.md",
		"2-some_thing.md",
	} {
		if !ValidateADRFilename(value) {
			t.Fatal("failed with right param: " + value)
		}
	}

}

var emptyFileListStub []string
var fileListStub = []string{
"000-test.md",
"001-test.md",
"002-test.md",
}

func TestGetLastId(t *testing.T) {
	var id = GetLastIdFromDir(emptyFileListStub)
	if id != 0 {
		t.Fatal("failed: expected 0, returned " + strconv.FormatInt(int64(id), 10))
	}

	id = GetLastIdFromDir(fileListStub)
	if id != 2 {
		t.Fatal("failed: expected 2, returned " + strconv.FormatInt(int64(id), 10))
	}
}
