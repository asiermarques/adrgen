package adr

import (
	"regexp"
	"strconv"
)

func GetLastIdFromDir(dirname string, ADRFilesFindOperation func(dirname string) []string) int {

	var filenames []string = ADRFilesFindOperation(dirname)
	var number = 0
	if len(filenames) > 0 {
		var re = regexp.MustCompile("[0-9]+")
		for _, name := range filenames {
			var current = re.FindAllString(name, 1)
			if len(current) > 0 {
				var rawNumber string = current[0]
				var currentNumber, _ = strconv.Atoi(rawNumber)
				if currentNumber > number {
					number = currentNumber
				}
			}
		}
		return number
	}
	return number
}
