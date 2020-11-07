package application

import (
	"../adr"
)

var createFilename func(id int, t string) string =
	adr.CreateFilename
var defaultTemplateContent func(t string) string =
	adr.DefaultTemplateContent
var createMetaContent func(m []string) string =
	adr.CreateMetaContent

var getLastIdFromFilenames func(s []string) int =
	adr.GetLastIdFromFilenames
var findFilesInDir func(d string) ([]string, error) =
	adr.FindADRFilesInDir
var writeFile func(f string, d string) (string, error) =
	adr.WriteFile

var createConfigFile func(dir string, templateFilename string, meta []string) error =
	adr.CreateConfigFile
var configFilename string =
	adr.CONFIG_FILENAME