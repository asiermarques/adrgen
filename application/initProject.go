package application

import (
	"fmt"

	"github.com/asiermarques/adrgen/domain"
)

// InitProject is the application service for initialize the workdir and configuration
//
func InitProject(
	config domain.Config,
	configManager domain.ConfigManager,
	templateWriter domain.TemplateWriter,
) error {
	err := templateWriter.Persist()
	if err != nil {
		return err
	}

	err = configManager.Persist(config)
	if err != nil {
		return fmt.Errorf("error creating config file %s %s", domain.CONFIG_FILENAME+".yml", err)
	}

	return nil
}
