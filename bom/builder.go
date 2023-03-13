package bom

import "github.com/spf13/afero"

// DefaultBOMProcessorBuilder holds the functions
// which need to be defined dynamically to enable
// dependency injection for easier testing.
type DefaultBOMProcessorBuilder struct {
	fileSystem afero.Fs
}

// NewDefaultBOMProcessorBuilder returns a new instance
// of the default builder for a BOM processor.
func NewDefaultBOMProcessorBuilder() *DefaultBOMProcessorBuilder {
	return &DefaultBOMProcessorBuilder{}
}

// SetFileSystem allows to define the file system which will
// be used by the bom processor for file operations.
func (b *DefaultBOMProcessorBuilder) SetFileSystem(fileSystem afero.Fs) {
	b.fileSystem = fileSystem
}

// GetBOMProcessor "builds" and returns the actual BOM
// processor.
func (b *DefaultBOMProcessorBuilder) GetBOMProcessor() DefaultBOMProcessor {
	if b.fileSystem == nil {
		b.fileSystem = afero.NewOsFs()
	}
	return DefaultBOMProcessor{fileSystem: b.fileSystem}
}
