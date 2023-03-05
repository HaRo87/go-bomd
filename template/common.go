package template

import (
	"io/fs"
	"os"

	cdx "github.com/CycloneDX/cyclonedx-go"
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
	Generate(template TemplateInfo) (err error)
}

// TemplateProcessorBuilder represents the interface a builder,
// following the builder design pattern:
// https://refactoring.guru/design-patterns/builder
// must implement.
type TemplateProcessorBuilder interface {
	SetStat(func(name string) (fs.FileInfo, error))
	SetCreate(func(name string) (*os.File, error))
	GetTemplateProcessor() DefaultTemplateProcessor
}
