package _infrastructure

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal/config"

	"github.com/spf13/viper"
)

type privateConfigFileManager struct {
	directory string
}

// CreateConfigFileManager creates an instance of domain.Manager that manages the configuration in a config file
//
func CreateConfigFileManager(directory string) config.Manager {
	return privateConfigFileManager{directory}
}

func (manager privateConfigFileManager) Persist(configData config.Config) error {
	viper.SetConfigName(config.FILENAME)
	viper.SetConfigType(config.FORMAT)
	viper.Set("directory", configData.TargetDirectory)
	viper.Set("template_file", configData.TemplateFilename)
	viper.Set("default_meta", configData.MetaParams)
	viper.Set("supported_statuses", configData.Statuses)
	viper.Set("default_status", configData.DefaultStatus)
	viper.Set("id_digit_number", configData.IdDigitNumber)
	return viper.WriteConfigAs(config.FILENAME + "." + config.FORMAT)
}

func (manager privateConfigFileManager) Read() (config.Config, error) {
	viper.SetConfigName(config.FILENAME)
	viper.SetConfigType(config.FORMAT)
	viper.AddConfigPath(manager.directory)
	err := viper.ReadInConfig()
	if err != nil {
		return config.Config{}, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return config.Config{
		TargetDirectory:  viper.GetString("directory"),
		TemplateFilename: viper.GetString("template_file"),
		MetaParams:       viper.GetStringSlice("default_meta"),
		Statuses:         viper.GetStringSlice("supported_statuses"),
		DefaultStatus:    viper.GetString("default_status"),
		IdDigitNumber:    viper.GetInt("id_digit_number"),
	}, nil
}

func (manager privateConfigFileManager) GetDefault() config.Config {
	return config.Config{
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
