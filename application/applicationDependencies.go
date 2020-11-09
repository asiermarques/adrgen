package application

import "github.com/asiermarques/adrgen/adr"

var createFilename func(id int, t string) string =
	adr.CreateFilename
var defaultTemplateContent func(d string, t string, s string) string =
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
	func(directory string, templateFilename string, meta []string) error {
		config := adr.DefaultConfig()
		config.TargetDirectory = directory
		config.TemplateFilename = templateFilename
		config.MetaParams = meta
		return adr.CreateConfigFile(config)
	}

var configFilename string =
	adr.CONFIG_FILENAME