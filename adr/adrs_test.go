package adr

import (
	"strconv"
	"testing"
)

func getEmptyFileListStub(dirname string) []string {
	return []string{}
}

func getFileListStub(dirname string) []string {
	return []string{
		"000-test.md",
		"001-test.md",
		"002-test.md",
	}
}

func TestGetLastId(t *testing.T) {
	var id = GetLastIdFromDir(".", getEmptyFileListStub)
	if id != 0 {
		t.Fatal("failed: expected 0, returned " + strconv.FormatInt(int64(id), 10))
	}

	id = GetLastIdFromDir(".", getFileListStub)
	if id != 2 {
		t.Fatal("failed: expected 2, returned " + strconv.FormatInt(int64(id), 10))
	}
}
