package replicator

import (
	"fmt"
	"text/template"

	"github.com/spf13/afero"
)

// DefaultTemplateProcessor holds the functions
// which need to be defined dynamically to enable
// dependency injection for easier testing.
type DefaultTemplateProcessor struct {
	fileSystem afero.Fs
	parseFiles func(filenames ...string) (*template.Template, error)
}

// Validate tries to parse the provided template.
func (p DefaultTemplateProcessor) Validate(template TemplateInfo) (err error) {
	_, err = p.fileSystem.Stat(template.InputFilePath)
	if err != nil {
		return
	}
	_, err = p.parseFiles(template.InputFilePath)
	return
}

// Execute executes the template and writes it to defined
// output file.
func (p DefaultTemplateProcessor) Execute(template TemplateInfo) (err error) {
	tmpl, err := p.parseFiles(template.InputFilePath)
	if err != nil {
		return
	}
	file, err := p.fileSystem.Create(template.OutputFilePath)
	if err != nil {
		return
	}
	err = tmpl.Execute(file, template.BOM)
	return
}

// Generate generates a default markdown template and stores it
// at the location defined by filePath.
func (p DefaultTemplateProcessor) Generate(filePath string) (err error) {
	file, err := p.fileSystem.Create(filePath)
	if err != nil {
		return
	}
	data := "# SBOM for {{ .Metadata.Component.Name }}" +
		"| Name | Version | Type |" +
		"| ---- | ------- | ---- |" +
		"{{ range .Components }}" +
		"| {{ .Name }} | {{ .Version }} | {{ .Type }} |" +
		"{{ end }}"
	_, err = file.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("%s; %w", err.Error(), file.Close())
	}
	return file.Close()
}
