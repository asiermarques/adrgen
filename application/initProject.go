package application

import (
	"fmt"
	"path/filepath"

	"github.com/asiermarques/adrgen/adr"
)

func InitProject(targetDirectory string, templateFilename string, metaParams []string) error {
	_, err := adr.WriteFile(
		filepath.Join(targetDirectory, templateFilename),
		adr.DefaultTemplateContent("{date}", "{title}", "{status}"),
	)
	if err != nil {
		return fmt.Errorf(
			"error creating template file %s ",
			filepath.Join(targetDirectory, templateFilename),
		)
	}

	err = createConfigFile(targetDirectory, templateFilename, metaParams)
	if err != nil {
		return fmt.Errorf(
			"error creating config file %s %s",
			filepath.Join(targetDirectory, adr.CONFIG_FILENAME),
			err,
		)
	}

	return nil
}

func createConfigFile(directory string, templateFilename string, meta []string) error {
	config := adr.DefaultConfig()
	config.TargetDirectory = directory
	config.TemplateFilename = templateFilename
	config.MetaParams = meta
	return adr.CreateConfigFile(config)
}
