package replicator

import (
	"text/template"

	"github.com/spf13/afero"
)

// DefaultTemplateProcessorBuilder holds the functions
// which need to be defined dynamically to enable
// dependency injection for easier testing.
type DefaultTemplateProcessorBuilder struct {
	fileSystem afero.Fs
	parseFiles func(filenames ...string) (*template.Template, error)
}

// NewDefaultTemplateProcessorBuilder returns a new instance
// of the default builder for a template processor.
func NewDefaultTemplateProcessorBuilder() *DefaultTemplateProcessorBuilder {
	return &DefaultTemplateProcessorBuilder{}
}

// SetFileSystem allows to define the file system which will
// be used by the template processor for file operations.
func (b *DefaultTemplateProcessorBuilder) SetFileSystem(fileSystem afero.Fs) {
	b.fileSystem = fileSystem
}

// SetParseFiles allows to define the ParseFiles function which will
// be used by the template processor for file operations.
func (b *DefaultTemplateProcessorBuilder) SetParseFiles(parseFiles func(filenames ...string) (*template.Template, error)) {
	b.parseFiles = parseFiles
}

// GetTemplateProcessor "builds" and returns the actual template
// processor.
func (b *DefaultTemplateProcessorBuilder) GetTemplateProcessor() DefaultTemplateProcessor {
	if b.fileSystem == nil {
		b.fileSystem = afero.NewOsFs()
	}
	if b.parseFiles == nil {
		b.parseFiles = template.ParseFiles
	}
	return DefaultTemplateProcessor{fileSystem: b.fileSystem, parseFiles: b.parseFiles}
}
