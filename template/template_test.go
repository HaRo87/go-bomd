package template

import (
	"fmt"
	"testing"
	gtemplate "text/template"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	gmock "gitlab.com/HaRo87go-bomd/mock"
)

func getDefaultTemplateProcessor() DefaultTemplateProcessor {
	builder := NewDefaultTemplateProcessorBuilder()
	return builder.GetTemplateProcessor()
}

func TestValidateNonExistingInputFileReturnsError(t *testing.T) {
	proc := getDefaultTemplateProcessor()
	err := proc.Validate(TemplateInfo{InputFilePath: "some.tmpl"})
	assert.Error(t, err)
	assert.Equal(t, "stat some.tmpl: no such file or directory", err.Error())
}

func TestValidateParseFilesIssueReturnsError(t *testing.T) {
	builder := NewDefaultTemplateProcessorBuilder()
	fileMock := afero.NewMemMapFs()
	templateMock := new(gmock.MockTemplate)
	builder.SetFileSystem(fileMock)
	builder.SetParseFiles(templateMock.ParseFiles)
	proc := builder.GetTemplateProcessor()
	fileMock.Create("some.tmpl")
	templateMock.On("ParseFiles", []string{"some.tmpl"}).Return(new(gtemplate.Template), fmt.Errorf("Some error"))
	err := proc.Validate(TemplateInfo{InputFilePath: "some.tmpl"})
	assert.Error(t, err)
	assert.Equal(t, "Some error", err.Error())
}

func TestExecuteParseFilesIssueReturnsError(t *testing.T) {
	builder := NewDefaultTemplateProcessorBuilder()
	templateMock := new(gmock.MockTemplate)
	builder.SetParseFiles(templateMock.ParseFiles)
	proc := builder.GetTemplateProcessor()
	templateMock.On("ParseFiles", []string{"some.tmpl"}).Return(new(gtemplate.Template), fmt.Errorf("Some error"))
	err := proc.Execute(TemplateInfo{InputFilePath: "some.tmpl"})
	assert.Error(t, err)
	assert.Equal(t, "Some error", err.Error())
}

func TestExecuteFsCreateIssueReturnsError(t *testing.T) {
	builder := NewDefaultTemplateProcessorBuilder()
	fileMock := afero.NewReadOnlyFs(afero.NewMemMapFs())
	templateMock := new(gmock.MockTemplate)
	builder.SetFileSystem(fileMock)
	builder.SetParseFiles(templateMock.ParseFiles)
	proc := builder.GetTemplateProcessor()
	templateMock.On("ParseFiles", []string{"some.tmpl"}).Return(new(gtemplate.Template), nil)
	err := proc.Execute(TemplateInfo{InputFilePath: "some.tmpl", OutputFilePath: "some.md"})
	assert.Error(t, err)
	assert.Equal(t, "operation not permitted", err.Error())
}

func TestExecuteInvalidTemplateReturnsError(t *testing.T) {
	builder := NewDefaultTemplateProcessorBuilder()
	fileMock := afero.NewMemMapFs()
	templateMock := new(gmock.MockTemplate)
	builder.SetFileSystem(fileMock)
	builder.SetParseFiles(templateMock.ParseFiles)
	proc := builder.GetTemplateProcessor()
	templateMock.On("ParseFiles", []string{"some.tmpl"}).Return(new(gtemplate.Template), nil)
	err := proc.Execute(TemplateInfo{InputFilePath: "some.tmpl", OutputFilePath: "some.md"})
	assert.Error(t, err)
	assert.Equal(t, "template: : \"\" is an incomplete or empty template", err.Error())
}

func TestExecuteSuccess(t *testing.T) {
	builder := NewDefaultTemplateProcessorBuilder()
	fileMock := afero.NewOsFs()
	builder.SetFileSystem(fileMock)
	proc := builder.GetTemplateProcessor()
	dirPath, err := afero.TempDir(fileMock, "", "")
	assert.NoError(t, err)
	file, err := fileMock.Create(dirPath + "/some.tmpl")
	assert.NoError(t, err)
	data := "This bom contains {{ range .Components }} {{ .Name }} {{ end }}  "
	file.Write([]byte(data))
	file.Close()
	bom := cdx.NewBOM()
	components := []cdx.Component{
		{
			BOMRef:     "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
			Type:       cdx.ComponentTypeLibrary,
			Author:     "CycloneDX",
			Name:       "cyclonedx-go",
			Version:    "v0.3.0",
			PackageURL: "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
			Licenses:   &cdx.Licenses{cdx.LicenseChoice{License: &cdx.License{ID: "MIT", Name: "MIT License"}}},
		},
	}
	bom.Components = &components
	err = proc.Execute(TemplateInfo{InputFilePath: dirPath + "/some.tmpl", OutputFilePath: dirPath + "/some.md", BOM: bom})
	assert.NoError(t, err)
	file, err = fileMock.Open(dirPath + "/some.md")
	assert.NoError(t, err)
	result, err := afero.ReadAll(file)
	assert.NoError(t, err)
	assert.Contains(t, string(result), "cyclonedx-go")
}
