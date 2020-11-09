package cmd

import "os"

func cleanTestFiles(files []string) {
	for _, file := range files {
		os.RemoveAll(file)
	}
}