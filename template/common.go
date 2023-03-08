package template

import (
	gTemplate "text/template"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/spf13/afero"
)

// TemplateInfo represents all information
// which is required to execute a template
type TemplateInfo struct {
	InputFilePath  string
	OutputFilePath string
	BOM            *cdx.BOM
}

// TemplateProcessor represents the interface and thereby all the
// functions a template processor must implement.
type TemplateProcessor interface {
	Validate(template TemplateInfo) (err error)
	Execute(template TemplateInfo) (err error)
	Generate(filePath string) (err error)
}

// TemplateProcessorBuilder represents the interface a builder,
// following the builder design pattern:
// https://refactoring.guru/design-patterns/builder
// must implement.
type TemplateProcessorBuilder interface {
	SetFileSystem(afero.Fs)
	SetParseFiles(func(filenames ...string) (*gTemplate.Template, error))
	GetTemplateProcessor() DefaultTemplateProcessor
}
