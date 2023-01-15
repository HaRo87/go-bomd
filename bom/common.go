package bom

import (
	"io/fs"

	cdx "github.com/CycloneDX/cyclonedx-go"
)

// BOMProcessor represents the interface and thereby all the
// functions a BOM processor must implement.
type BOMProcessor interface {
	GetBOM(filePath string) (bom cdx.BOM, err error)
	ValidateBOM(bom *cdx.BOM) (err error)
	ValidateComponentLicenses(bom *cdx.BOM) (err error)
}

// BOMProcessorBuilder represents the interface a builder,
// following the builder design pattern:
// https://refactoring.guru/design-patterns/builder
// must implement.
type BOMProcessorBuilder interface {
	SetStat(func(name string) (fs.FileInfo, error))
	SetReadFile(func(filename string) ([]byte, error))
	GetBOMProcessor() DefaultBOMProcessor
}

// bomCheck represents the interface a handler,
// following the chain of responsibility pattern:
// https://refactoring.guru/design-patterns/chain-of-responsibility
// must implement.
type bomCheck interface {
	execute(bom *cdx.BOM) (err error)
	setNext(next bomCheck)
}
