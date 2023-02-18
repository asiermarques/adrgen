package _infrastructure

import (
	"github.com/asiermarques/adrgen/internal/config"
	"github.com/asiermarques/adrgen/internal/template"
	"os"
	"path/filepath"
)

type privateCustomTemplateContentFileReader struct {
	configuration config.Config
}

// CreateCustomTemplateContentFileReader creates an domain.CustomContentReader that reads a template content from a file
//
func CreateCustomTemplateContentFileReader(
	config config.Config,
) template.CustomContentReader {
	return privateCustomTemplateContentFileReader{configuration: config}
}

func (r privateCustomTemplateContentFileReader) Read() (string, error) {
	rootDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return GetFileContent(filepath.Join(rootDirectory, r.configuration.TemplateFilename))
}

type privateTemplateFileWriter struct {
	configuration config.Config
}

// CreateTemplateFileWriter creates an domain.Writer that writes the template content in a file
//
func CreateTemplateFileWriter(config config.Config) template.Writer {
	return privateTemplateFileWriter{configuration: config}
}

func (w privateTemplateFileWriter) Persist() error {
	return WriteFile(w.configuration.TemplateFilename, template.DEFAULT_CONTENT)
}
