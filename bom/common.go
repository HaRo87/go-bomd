package bom

import (
	"io/fs"

	cdx "github.com/CycloneDX/cyclonedx-go"
)

type BOMProcessor interface {
	GetBOM(filePath string) (bom cdx.BOM, err error)
	ValidateBOM(bom *cdx.BOM) (err error)
	ValidateComponentLicenses(bom *cdx.BOM) (err error)
}

type BOMProcessorBuilder interface {
	SetStat(func(name string) (fs.FileInfo, error))
	SetReadFile(func(filename string) ([]byte, error))
	GetBOMProcessor() DefaultBOMProcessor
}

type bomCheck interface {
	execute(bom *cdx.BOM) (err error)
	setNext(next bomCheck)
}
