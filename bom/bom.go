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

type bomComponentCheck struct {
	next bomCheck
}

func (c *bomComponentCheck) execute(bom *cdx.BOM) (err error) {
	if *bom.Components == nil {
		err = fmt.Errorf("BOM does not contain any components")
		return
	}
	if len(*bom.Components) < 1 {
		err = fmt.Errorf("No components in BOM")
		return
	}
	err = c.next.execute(bom)
	return
}

func (c *bomComponentCheck) setNext(next bomCheck) {
	c.next = next
}

type bomMetaCheck struct {
	next bomCheck
}

func (c *bomMetaCheck) execute(bom *cdx.BOM) (err error) {
	if bom.Metadata == nil {
		err = fmt.Errorf("BOM does not contain any meta data")
		return
	}
	err = c.next.execute(bom)
	return
}

func (c *bomMetaCheck) setNext(next bomCheck) {
	c.next = next
}

type bomLicenseCheck struct {
	next bomCheck
}

func (c *bomLicenseCheck) execute(bom *cdx.BOM) (err error) {
	for _, comp := range *bom.Components {
		if comp.Licenses == nil {
			err = fmt.Errorf("Component: %s without licenses detected", comp.Name)
			return
		} else {
			if len(*comp.Licenses) == 0 {
				err = fmt.Errorf("Component: %s without licenses detected", comp.Name)
				return
			}
		}
	}
	err = c.next.execute(bom)
	return
}

func (c *bomLicenseCheck) setNext(next bomCheck) {
	c.next = next
}

// GetBOM reads a CycloneDX BOM in json format from
// the provided filePath and returns a BOM object.
func (p DefaultBOMProcessor) GetBOM(filePath string) (bom cdx.BOM, err error) {
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

// ValidateBOM checks the provided BOM for required information
// which needs to be present for bomd to work.
func (p DefaultBOMProcessor) ValidateBOM(bom *cdx.BOM) (err error) {
	comp := &bomComponentCheck{}
	meta := &bomMetaCheck{}
	meta.setNext(comp)
	err = meta.execute(bom)
	return
}

// ValidateComponentLicenses checks all components in
// the provided BOM contain license information.
func (p DefaultBOMProcessor) ValidateComponentLicenses(bom *cdx.BOM) (err error) {
	lic := &bomLicenseCheck{}
	comp := &bomComponentCheck{}
	comp.setNext(lic)
	err = comp.execute(bom)
	return
}
