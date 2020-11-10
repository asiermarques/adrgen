package adr

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

const CONFIG_FILENAME = "adrgen.config"
const CONFIG_FORMAT = "yaml"

type Config struct {
	TargetDirectory  string
	TemplateFilename string
	MetaParams       []string
	Statuses         []string
	DefaultStatus    string
	IdDigitNumber    int
}

// CreateConfigFile creates a config file in the current directory
//
func CreateConfigFile(config Config) error {
	viper.SetConfigName(CONFIG_FILENAME)
	viper.SetConfigType(CONFIG_FORMAT)
	viper.Set("directory", config.TargetDirectory)
	viper.Set("template_file", filepath.Join(config.TargetDirectory, config.TemplateFilename))
	viper.Set("default_meta", config.MetaParams)
	viper.Set("supported_statuses", config.Statuses)
	viper.Set("default_status", config.DefaultStatus)
	viper.Set("id_digit_number", config.IdDigitNumber)
	return viper.WriteConfigAs(CONFIG_FILENAME + ".yml")
}

// GetConfig reads the configuration and return the Config object or an error
//
func GetConfig(directory string) (Config, error) {
	viper.SetConfigName(CONFIG_FILENAME)
	viper.SetConfigType(CONFIG_FORMAT)
	viper.AddConfigPath(directory)
	err := viper.ReadInConfig()
	if err != nil {
		return DefaultConfig(), fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return Config{
		TargetDirectory:  viper.GetString("directory"),
		TemplateFilename: viper.GetString("template_file"),
		MetaParams:       viper.GetStringSlice("default_meta"),
		Statuses:         viper.GetStringSlice("supported_statuses"),
		DefaultStatus:    viper.GetString("default_status"),
		IdDigitNumber:    viper.GetInt("id_digit_number"),
	}, nil
}

// DefaultConfig return a Config object with the default parameters
//
func DefaultConfig() Config {
	return Config{
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
