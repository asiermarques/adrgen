package application

import (
	"github.com/asiermarques/adrgen/adr"
)

var createFilename = adr.CreateFilename
var defaultTemplateContent = adr.DefaultTemplateContent
var createMetaContent = adr.CreateMetaContent

var getLastIdFromDir = adr.GetLastIdFromDir
var findFilesInDir = adr.FindADRFilesInDir
var writeFile = adr.WriteFile

var createConfigFile = adr.CreateConfigFile