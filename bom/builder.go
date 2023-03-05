package bom

import (
	"io/fs"
	"os"
)

// DefaultBOMProcessorBuilder holds the functions
// which need to be defined dynamically to enable
// dependency injection for easier testing.
type DefaultBOMProcessorBuilder struct {
	stat     func(name string) (fs.FileInfo, error)
	readFile func(filename string) ([]byte, error)
}

// NewDefaultBOMProcessorBuilder returns a new instance
// of the default builder for a BOM processor.
func NewDefaultBOMProcessorBuilder() *DefaultBOMProcessorBuilder {
	return &DefaultBOMProcessorBuilder{}
}

// SetStat allows to define the stat function which will
// be used by the BOM processor for file operations.
func (b *DefaultBOMProcessorBuilder) SetStat(stat func(name string) (fs.FileInfo, error)) {
	b.stat = stat
}

// SetReadFile allows to define the readFile function which will
// be used by the BOM processor for file operations.
func (b *DefaultBOMProcessorBuilder) SetReadFile(readFile func(filename string) ([]byte, error)) {
	b.readFile = readFile
}

// GetBOMProcessor "builds" and returns the actual BOM
// processor.
func (b *DefaultBOMProcessorBuilder) GetBOMProcessor() DefaultBOMProcessor {
	if b.stat == nil {
		b.stat = os.Stat
	}
	if b.readFile == nil {
		b.readFile = os.ReadFile
	}
	return DefaultBOMProcessor{stat: b.stat, readFile: b.readFile}
}
