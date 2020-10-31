package main

import "fmt"

func CreateFilename(id int, title string) string  {
	return fmt.Sprintf("%d-%s.md", id, title)
}
