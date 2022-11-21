package bom

import (
	"bytes"
	"fmt"
	"io/fs"
	"strings"

	cdx "github.com/CycloneDX/cyclonedx-go"
)

type DefaultBOMProcessor struct {
	stat     func(name string) (fs.FileInfo, error)
	readFile func(filename string) ([]byte, error)
}

// GetBom reads a CycloneDX BOM in json format from
// the provided filePath and returns a BOM object.
func (p DefaultBOMProcessor) GetBom(filePath string) (bom cdx.BOM, err error) {
	if !strings.HasSuffix(filePath, ".json") {
		err = fmt.Errorf("Only JSON file format supported")
		return
	}
	_, err = p.stat(filePath)
	if err != nil {
		return
	}
	content, err := p.readFile(filePath)
	if err != nil {
		return
	}
	decoder := cdx.NewBOMDecoder(bytes.NewReader(content), cdx.BOMFileFormatJSON)
	err = decoder.Decode(&bom)
	if err != nil {
		return
	}
	return
}
