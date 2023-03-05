package template

import (
	"io/fs"
	"os"
	gTemplate "text/template"
)

// DefaultTemplateProcessor holds the functions
// which need to be defined dynamically to enable
// dependency injection for easier testing.
type DefaultTemplateProcessor struct {
	stat       func(name string) (fs.FileInfo, error)
	create     func(filename string) (*os.File, error)
	parseFiles func(filenames ...string) (*gTemplate.Template, error)
}

// Validate tries to parse the provided template.
func (p DefaultTemplateProcessor) Validate(template TemplateInfo) (err error) {
	_, err = p.stat(template.InputFilePath)
	if err != nil {
		return
	}
	_, err = p.parseFiles(template.InputFilePath)
	return
}

// Generate executes the template and writes it to defined
// output file.
func (p DefaultTemplateProcessor) Generate(template TemplateInfo) (err error) {
	tmpl, err := p.parseFiles(template.InputFilePath)
	if err != nil {
		return
	}
	file, err := p.create(template.OutputFilePath)
	if err != nil {
		return
	}
	err = tmpl.Execute(file, template.BOM)
	return
}
