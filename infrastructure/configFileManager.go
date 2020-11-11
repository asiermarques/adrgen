package infrastructure

import (
	"fmt"
	"github.com/asiermarques/adrgen/domain"
	"github.com/spf13/viper"
	"path/filepath"
)

type privateConfigFileManager struct {
	directory string
}

func CreateConfigFileManager(directory string) domain.ConfigManager {
	return privateConfigFileManager{directory}
}

func (manager privateConfigFileManager) Persist(config domain.Config) error {
	viper.SetConfigName(domain.CONFIG_FILENAME)
	viper.SetConfigType(domain.CONFIG_FORMAT)
	viper.Set("directory", config.TargetDirectory)
	viper.Set("template_file", filepath.Join(config.TargetDirectory, config.TemplateFilename))
	viper.Set("default_meta", config.MetaParams)
	viper.Set("supported_statuses", config.Statuses)
	viper.Set("default_status", config.DefaultStatus)
	viper.Set("id_digit_number", config.IdDigitNumber)
	return viper.WriteConfigAs(domain.CONFIG_FILENAME + ".yml")
}

func (manager privateConfigFileManager) Read() (domain.Config, error) {
	viper.SetConfigName(domain.CONFIG_FILENAME)
	viper.SetConfigType(domain.CONFIG_FORMAT)
	viper.AddConfigPath(manager.directory)
	err := viper.ReadInConfig()
	if err != nil {
		return domain.Config{}, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return domain.Config{
		TargetDirectory:  viper.GetString("directory"),
		TemplateFilename: viper.GetString("template_file"),
		MetaParams:       viper.GetStringSlice("default_meta"),
		Statuses:         viper.GetStringSlice("supported_statuses"),
		DefaultStatus:    viper.GetString("default_status"),
		IdDigitNumber:    viper.GetInt("id_digit_number"),
	}, nil
}

func (manager privateConfigFileManager) GetDefault() domain.Config {
	return domain.Config{
		TemplateFilename: "",
		TargetDirectory:  ".",
		Statuses: []string{
			"proposed",
			"accepted",
			"rejected",
			"superseded",
			"amended",
			"deprecated",
		},
		DefaultStatus: "proposed",
		MetaParams:    []string{},
		IdDigitNumber: 4,
	}
}
