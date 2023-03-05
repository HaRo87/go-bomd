package template

import (
	"io/fs"
	"os"
	gTemplate "text/template"
)

// DefaultTemplateProcessorBuilder holds the functions
// which need to be defined dynamically to enable
// dependency injection for easier testing.
type DefaultTemplateProcessorBuilder struct {
	stat       func(name string) (fs.FileInfo, error)
	create     func(name string) (*os.File, error)
	parseFiles func(filenames ...string) (*gTemplate.Template, error)
}

// NewDefaultTemplateProcessorBuilder returns a new instance
// of the default builder for a template processor.
func NewDefaultTemplateProcessorBuilder() *DefaultTemplateProcessorBuilder {
	return &DefaultTemplateProcessorBuilder{}
}

// SetStat allows to define the stat function which will
// be used by the template processor for file operations.
func (b *DefaultTemplateProcessorBuilder) SetStat(stat func(name string) (fs.FileInfo, error)) {
	b.stat = stat
}

// SetCreate allows to define the Create function which will
// be used by the template processor for file operations.
func (b *DefaultTemplateProcessorBuilder) SetCreate(create func(name string) (*os.File, error)) {
	b.create = create
}

// SetParseFiles allows to define the ParseFiles function which will
// be used by the template processor for file operations.
func (b *DefaultTemplateProcessorBuilder) SetParseFiles(parseFiles func(filenames ...string) (*gTemplate.Template, error)) {
	b.parseFiles = parseFiles
}

// GetTemplateProcessor "builds" and returns the actual template
// processor.
func (b *DefaultTemplateProcessorBuilder) GetTemplateProcessor() DefaultTemplateProcessor {
	if b.stat == nil {
		b.stat = os.Stat
	}
	if b.create == nil {
		b.create = os.Create
	}
	if b.parseFiles == nil {
		b.parseFiles = gTemplate.ParseFiles
	}
	return DefaultTemplateProcessor{stat: b.stat, create: b.create, parseFiles: b.parseFiles}
}
