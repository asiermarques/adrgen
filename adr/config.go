package adr

import (
	"github.com/spf13/viper"
	"path/filepath"
)

const CONFIG_FILENAME = "adrgen.config"
const CONFIG_FORMAT = "yaml"

func CreateConfigFile(targetDirectory string, templateFilename string, metaParams []string) error {
	viper.SetConfigName(CONFIG_FILENAME)
	viper.SetConfigType(CONFIG_FORMAT)
	viper.Set("directory", targetDirectory)
	viper.Set("template_file", filepath.Join(targetDirectory, templateFilename))
	viper.Set("defaultMeta", metaParams)
	return viper.WriteConfig()
}