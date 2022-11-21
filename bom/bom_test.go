package bom

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getDefaultBOMProcessor() DefaultBOMProcessor {
	builder := NewDefaultBOMProcessorBuilder()
	return builder.GetBOMProcessor()
}

func TestHasWrongSuffixReturnsError(t *testing.T) {
	proc := getDefaultBOMProcessor()
	_, err := proc.GetBom("bom.xml")
	assert.Error(t, err)
	assert.Equal(t, "Only JSON file format supported", err.Error())
}

func TestFileDoesNotExistReturnsError(t *testing.T) {
	builder := NewDefaultBOMProcessorBuilder()
	bomFileMock := new(MockBOMFile)
	builder.SetStat(bomFileMock.Stat)
	proc := builder.GetBOMProcessor()
	bomFileMock.On("Stat", "bom.json").Return(new(MockFileInfo), fmt.Errorf("File does not exist"))
	_, err := proc.GetBom("bom.json")
	assert.Error(t, err)
	assert.Equal(t, "File does not exist", err.Error())
}

func TestCannotReadFileReturnsError(t *testing.T) {
	builder := NewDefaultBOMProcessorBuilder()
	bomFileMock := new(MockBOMFile)
	builder.SetStat(bomFileMock.Stat)
	builder.SetReadFile(bomFileMock.ReadFile)
	proc := builder.GetBOMProcessor()
	bomFileMock.On("Stat", "bom.json").Return(new(MockFileInfo), nil)
	bomFileMock.On("ReadFile", "bom.json").Return([]byte{}, fmt.Errorf("Content could not be read"))
	_, err := proc.GetBom("bom.json")
	assert.Error(t, err)
	assert.Equal(t, "Content could not be read", err.Error())
}
