package domain

// CONFIG_FILENAME the name for the config file
//
const CONFIG_FILENAME = "adrgen.config"

// CONFIG_FORMAT the configuration format
//
const CONFIG_FORMAT = "yaml"

// Config the Configuration type with all the supported values
//
type Config struct {
	TargetDirectory  string
	TemplateFilename string
	MetaParams       []string
	Statuses         []string
	DefaultStatus    string
	IdDigitNumber    int
}

type ConfigManager interface {
	Persist(config Config) error
	Read() (Config, error)
	GetDefault() Config
}