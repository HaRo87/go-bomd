package template

import (
	"fmt"
	"testing"
	gtemplate "text/template"

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
	fileMock := new(gmock.MockFile)
	templateMock := new(gmock.MockTemplate)
	builder.SetStat(fileMock.Stat)
	builder.SetCreate(fileMock.Create)
	builder.SetParseFiles(templateMock.ParseFiles)
	proc := builder.GetTemplateProcessor()
	fileMock.On("Stat", "some.tmpl").Return(new(gmock.MockFileInfo), nil)
	templateMock.On("ParseFiles", []string{"some.tmpl"}).Return(new(gtemplate.Template), fmt.Errorf("Some error"))
	err := proc.Validate(TemplateInfo{InputFilePath: "some.tmpl"})
	assert.Error(t, err)
	assert.Equal(t, "Some error", err.Error())
}
