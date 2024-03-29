package config

// FILENAME the name for the config file
//
const FILENAME = "adrgen.config"

// FORMAT the configuration format
//
const FORMAT = "yaml"

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

// Manager is the interface for the service responsible for reading and persisting the configuration
//
type Manager interface {
	Persist(config Config) error
	Read() (Config, error)
	GetDefault() Config
}
