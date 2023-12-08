package project

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal/config"
	"github.com/asiermarques/adrgen/internal/template"
)

// InitProject is the application service for initialize the workdir and configuration
//
func InitProject(
	configData config.Config,
	configManager config.Manager,
	templateWriter template.Writer,
) error {
	err := templateWriter.Persist()
	if err != nil {
		return err
	}

	err = configManager.Persist(configData)
	if err != nil {
		return fmt.Errorf("error creating config file %s %s", config.FILENAME + "." + config.FORMAT, err)
	}

	return nil
}
