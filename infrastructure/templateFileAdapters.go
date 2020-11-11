package infrastructure

import (
	"github.com/asiermarques/adrgen/domain"
	"os"
	"path/filepath"
)

type privateCustomTemplateContentFileReader struct {
	configuration domain.Config
}

func CreateCustomTemplateContentFileReader(config domain.Config) domain.CustomTemplateContentReader {
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
	configuration domain.Config
}

func CreateTemplateFileWriter(config domain.Config) domain.TemplateWriter {
	return privateTemplateFileWriter{configuration: config}
}

func (w privateTemplateFileWriter) Persist() error {
	return WriteFile(w.configuration.TemplateFilename, domain.DEFAULT_TEMPLATE_CONTENT)
}
