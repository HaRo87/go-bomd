package bom

import (
	"bytes"
	"fmt"
	"strings"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/spf13/afero"
)

// DefaultBOMProcessor holds the elements
// which need to be defined dynamically to enable
// dependency injection for easier testing.
type DefaultBOMProcessor struct {
	fileSystem afero.Fs
}

// bomComponentCheck represents a handler implementing
// the bomCheck interface and is used for evaluating
// BOM components.
type bomComponentCheck struct {
	next bomCheck
}

// execute checks whether the provided BOM contains any
// components or not.
func (c *bomComponentCheck) execute(bom *cdx.BOM) (err error) {
	if bom.Components == nil {
		err = fmt.Errorf("BOM does not contain any components")
		return
	}
	if len(*bom.Components) < 1 {
		err = fmt.Errorf("No components in BOM")
		return
	}
	if c.next != nil {
		err = c.next.execute(bom)
	}
	return
}

// setNext can be used to define the next handler.
func (c *bomComponentCheck) setNext(next bomCheck) {
	c.next = next
}

// bomMetaCheck represents a handler implementing
// the bomCheck interface and is used for evaluating
// BOM meta data.
type bomMetaCheck struct {
	next bomCheck
}

// execute checks whether the provided BOM contains any
// meta data or not.
func (c *bomMetaCheck) execute(bom *cdx.BOM) (err error) {
	if bom.Metadata == nil {
		err = fmt.Errorf("BOM does not contain any meta data")
		return
	}
	if c.next != nil {
		err = c.next.execute(bom)
	}
	return
}

// setNext can be used to define the next handler.
func (c *bomMetaCheck) setNext(next bomCheck) {
	c.next = next
}

// bomLicenseCheck represents a handler implementing
// the bomCheck interface and is used for evaluating
// BOM component license information.
type bomLicenseCheck struct {
	next bomCheck
}

// execute checks whether the provided BOM contains any
// components without proper license information or not.
func (c *bomLicenseCheck) execute(bom *cdx.BOM) (err error) {
	_, err = getComponentsWithEmptyLicenseIDs(bom)
	if err != nil {
		return
	}
	if c.next != nil {
		err = c.next.execute(bom)
	}
	return
}

// setNext can be used to define the next handler.
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
	_, err = p.fileSystem.Stat(filePath)
	if err != nil {
		return
	}
	file, err := p.fileSystem.Open(filePath)
	if err != nil {
		return
	}
	content, err := afero.ReadAll(file)
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

// GetComponentsWithEmptyLicenseIDs checks all components in
// the provided BOM contain license information and returns an
// error if that is not the case and a list of all components
// which lack proper license information.
func (p DefaultBOMProcessor) GetComponentsWithEmptyLicenseIDs(bom *cdx.BOM) (comps []string, err error) {
	return getComponentsWithEmptyLicenseIDs(bom)
}

func getComponentsWithEmptyLicenseIDs(bom *cdx.BOM) (comps []string, err error) {
	if bom.Components == nil {
		err = fmt.Errorf("No components in BOM")
		return
	}
	for _, comp := range *bom.Components {
		if comp.Licenses == nil {
			err = fmt.Errorf("Component(s) without licenses detected")
			comps = append(comps, comp.Name)
		} else {
			if len(*comp.Licenses) == 0 {
				err = fmt.Errorf("Component(s) without licenses detected")
				comps = append(comps, comp.Name)
			} else {
				for _, license := range *comp.Licenses {
					if len(license.License.ID) == 0 {
						err = fmt.Errorf("Component(s) without licenses detected")
						comps = append(comps, comp.Name)
					}
				}
			}
		}
	}
	return
}
