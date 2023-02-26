package bom

import (
	"fmt"
	"testing"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/stretchr/testify/assert"
)

func getDefaultBOMProcessor() DefaultBOMProcessor {
	builder := NewDefaultBOMProcessorBuilder()
	return builder.GetBOMProcessor()
}

func TestGetBOMHasWrongSuffixReturnsError(t *testing.T) {
	proc := getDefaultBOMProcessor()
	_, err := proc.GetBOM("bom.xml")
	assert.Error(t, err)
	assert.Equal(t, "Only JSON file format supported", err.Error())
}

func TestGetBOMFileDoesNotExistReturnsError(t *testing.T) {
	builder := NewDefaultBOMProcessorBuilder()
	bomFileMock := new(MockBOMFile)
	builder.SetStat(bomFileMock.Stat)
	proc := builder.GetBOMProcessor()
	bomFileMock.On("Stat", "bom.json").Return(new(MockFileInfo), fmt.Errorf("File does not exist"))
	_, err := proc.GetBOM("bom.json")
	assert.Error(t, err)
	assert.Equal(t, "File does not exist", err.Error())
}

func TestGetBOMCannotReadFileReturnsError(t *testing.T) {
	builder := NewDefaultBOMProcessorBuilder()
	bomFileMock := new(MockBOMFile)
	builder.SetStat(bomFileMock.Stat)
	builder.SetReadFile(bomFileMock.ReadFile)
	proc := builder.GetBOMProcessor()
	bomFileMock.On("Stat", "bom.json").Return(new(MockFileInfo), nil)
	bomFileMock.On("ReadFile", "bom.json").Return([]byte{}, fmt.Errorf("Content could not be read"))
	_, err := proc.GetBOM("bom.json")
	assert.Error(t, err)
	assert.Equal(t, "Content could not be read", err.Error())
}

func TestGetBOMReadExaqmpleFileReturnsBOM(t *testing.T) {
	builder := NewDefaultBOMProcessorBuilder()
	proc := builder.GetBOMProcessor()
	bom, err := proc.GetBOM("../examples/boms/go-bomd-bom.json")
	assert.NoError(t, err)
	assert.Equal(t, "gitlab.com/HaRo87go-bomd", bom.Metadata.Component.Name)
}

func TestValidateBOMNoMetaDataReturnsError(t *testing.T) {
	bom := cdx.NewBOM()
	proc := getDefaultBOMProcessor()
	err := proc.ValidateBOM(bom)
	assert.Error(t, err)
	assert.Equal(t, "BOM does not contain any meta data", err.Error())
}

func TestValidateBOMNilComponentsReturnsError(t *testing.T) {
	bom := cdx.NewBOM()
	meta := cdx.Metadata{}
	bom.Metadata = &meta
	bom.Components = nil
	proc := getDefaultBOMProcessor()
	err := proc.ValidateBOM(bom)
	assert.Error(t, err)
	assert.Equal(t, "BOM does not contain any components", err.Error())
}

func TestValidateBOMNoComponentsReturnsError(t *testing.T) {
	bom := cdx.NewBOM()
	meta := cdx.Metadata{}
	bom.Metadata = &meta
	components := []cdx.Component{}
	bom.Components = &components
	proc := getDefaultBOMProcessor()
	err := proc.ValidateBOM(bom)
	assert.Error(t, err)
	assert.Equal(t, "No components in BOM", err.Error())
}

func TestValidateComponentLicensesNoLicenseReturnsError(t *testing.T) {
	bom := cdx.NewBOM()
	components := []cdx.Component{
		{
			BOMRef:     "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
			Type:       cdx.ComponentTypeLibrary,
			Author:     "CycloneDX",
			Name:       "cyclonedx-go",
			Version:    "v0.3.0",
			PackageURL: "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
		},
	}
	bom.Components = &components
	proc := getDefaultBOMProcessor()
	err := proc.ValidateComponentLicenses(bom)
	assert.Error(t, err)
	assert.Equal(t, "Component(s) without licenses detected", err.Error())
}

func TestValidateComponentLicensesEmptyLicenseReturnsError(t *testing.T) {
	bom := cdx.NewBOM()
	components := []cdx.Component{
		{
			BOMRef:     "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
			Type:       cdx.ComponentTypeLibrary,
			Author:     "CycloneDX",
			Name:       "cyclonedx-go",
			Version:    "v0.3.0",
			PackageURL: "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
			Licenses:   &cdx.Licenses{},
		},
	}
	bom.Components = &components
	proc := getDefaultBOMProcessor()
	err := proc.ValidateComponentLicenses(bom)
	assert.Error(t, err)
	assert.Equal(t, "Component(s) without licenses detected", err.Error())
}

func TestValidateComponentLicensesEmptyLicenseIDReturnsError(t *testing.T) {
	bom := cdx.NewBOM()
	components := []cdx.Component{
		{
			BOMRef:     "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
			Type:       cdx.ComponentTypeLibrary,
			Author:     "CycloneDX",
			Name:       "cyclonedx-go",
			Version:    "v0.3.0",
			PackageURL: "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
			Licenses:   &cdx.Licenses{cdx.LicenseChoice{License: &cdx.License{ID: ""}}},
		},
	}
	bom.Components = &components
	proc := getDefaultBOMProcessor()
	err := proc.ValidateComponentLicenses(bom)
	assert.Error(t, err)
	assert.Equal(t, "Component(s) without licenses detected", err.Error())
}

func TestValidateComponentLicensesSuccess(t *testing.T) {
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
	proc := getDefaultBOMProcessor()
	err := proc.ValidateComponentLicenses(bom)
	assert.NoError(t, err)
}

func TestGetComponentsWithEmptyLicenseIDsNoComponentsReturnsError(t *testing.T) {
	bom := cdx.NewBOM()
	meta := cdx.Metadata{}
	bom.Metadata = &meta
	bom.Components = nil
	proc := getDefaultBOMProcessor()
	_, err := proc.GetComponentsWithEmptyLicenseIDs(bom)
	assert.Error(t, err)
	assert.Equal(t, "No components in BOM", err.Error())
}

func TestGetComponentsWithEmptyLicenseIDsMultipleEmptyLicenseIDsReturnsError(t *testing.T) {
	bom := cdx.NewBOM()
	components := []cdx.Component{
		{
			BOMRef:     "pkg:golang/github.com/org/package-one@v0.1.0",
			Type:       cdx.ComponentTypeLibrary,
			Author:     "org",
			Name:       "package-one",
			Version:    "v0.1.0",
			PackageURL: "pkg:golang/github.com/org/package-one@v0.1.0",
			Licenses:   &cdx.Licenses{cdx.LicenseChoice{License: &cdx.License{ID: ""}}},
		},
		{
			BOMRef:     "pkg:golang/github.com/org/package-two@v0.1.1",
			Type:       cdx.ComponentTypeLibrary,
			Author:     "org",
			Name:       "package-two",
			Version:    "v0.1.1",
			PackageURL: "pkg:golang/github.com/org/package-two@v0.1.0",
		},
		{
			BOMRef:     "pkg:golang/github.com/org/package-three@v0.1.0",
			Type:       cdx.ComponentTypeLibrary,
			Author:     "org",
			Name:       "package-three",
			Version:    "v0.1.0",
			PackageURL: "pkg:golang/github.com/org/package-three@v0.1.0",
			Licenses:   &cdx.Licenses{},
		},
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
	proc := getDefaultBOMProcessor()
	comps, err := proc.GetComponentsWithEmptyLicenseIDs(bom)
	assert.Error(t, err)
	assert.Equal(t, "Component(s) without licenses detected", err.Error())
	assert.Equal(t, []string{"package-one", "package-two", "package-three"}, comps)
}
