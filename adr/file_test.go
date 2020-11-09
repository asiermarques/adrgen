package adr

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCreateFilename(t *testing.T) {
	testFilename := func(expectedString string, id int, digitNumber int) {
		result := CreateFilename(id, "New ADR", digitNumber)
		if result != expectedString {
			t.Fatal(fmt.Sprintf("failed: expected %s, returned %s", expectedString, result))
		}
	}

	testFilename("1-new-adr.md", 1, 0)
	testFilename("1-new-adr.md", 1, 1)
	testFilename("123-new-adr.md", 123, 1)
	testFilename("0123-new-adr.md", 123, 4)
	testFilename("00003-new-adr.md", 3, 5)
}

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
		"0-some_thing.md",
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
	var id = GetLastIdFromFilenames(emptyFileListStub)
	if id != 0 {
		t.Fatal("failed: expected 0, returned " + strconv.FormatInt(int64(id), 10))
	}

	id = GetLastIdFromFilenames(fileListStub)
	if id != 2 {
		t.Fatal("failed: expected 2, returned " + strconv.FormatInt(int64(id), 10))
	}
}

func TestFindADRFileById(t *testing.T) {
	file, err := FindADRFileById(1, fileListStub)
	if err != nil || file != "001-test.md" {
		t.Fatal(fmt.Sprintf("failed: expected 001-test.md, returned %s :%s", file, err))
	}
}
