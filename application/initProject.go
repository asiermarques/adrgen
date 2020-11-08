package application

import (
	"fmt"
	"path/filepath"
)

func InitProject(targetDirectory string, templateFilename string, metaParams []string) error  {

	_, err := writeFile(filepath.Join(targetDirectory, templateFilename), defaultTemplateContent("{title}", "{status"))
	if err!=nil {
		return fmt.Errorf("error creating template file %s ", filepath.Join(targetDirectory, templateFilename))
	}

	err = createConfigFile(targetDirectory, templateFilename, metaParams)
	if err != nil {
		return fmt.Errorf("error creating config file %s %s", filepath.Join(targetDirectory, configFilename), err)
	}

	return nil
}
