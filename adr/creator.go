package adr

import (
	"fmt"
	"github.com/gosimple/slug"
)

func CreateFilename(id int, title string) string  {
	return fmt.Sprintf("%d-%s.md", id, slug.Make(title))
}
