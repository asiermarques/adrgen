package filesystem

import "testing"

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
