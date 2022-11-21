package bom

import (
	"io/fs"
	"io/ioutil"
	"os"
)

type DefaultBOMProcessorBuilder struct {
	stat     func(name string) (fs.FileInfo, error)
	readFile func(filename string) ([]byte, error)
}

func NewDefaultBOMProcessorBuilder() *DefaultBOMProcessorBuilder {
	return &DefaultBOMProcessorBuilder{}
}

func (b *DefaultBOMProcessorBuilder) SetStat(stat func(name string) (fs.FileInfo, error)) {
	b.stat = stat
}

func (b *DefaultBOMProcessorBuilder) SetReadFile(readFile func(filename string) ([]byte, error)) {
	b.readFile = readFile
}

func (b *DefaultBOMProcessorBuilder) GetBOMProcessor() DefaultBOMProcessor {
	if b.stat == nil {
		b.stat = os.Stat
	}
	if b.readFile == nil {
		b.readFile = ioutil.ReadFile
	}
	return DefaultBOMProcessor{stat: b.stat, readFile: b.readFile}
}
