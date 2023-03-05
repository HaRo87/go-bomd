package bom

import (
	"io/fs"
	"os"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockFileInfo represents a mocked
// FileInfo object.
type MockFileInfo struct {
	FileName    string
	IsDirectory bool
}

func (m MockFileInfo) Name() string       { return m.FileName }
func (m MockFileInfo) Size() int64        { return int64(0) }
func (m MockFileInfo) Mode() os.FileMode  { return os.ModePerm }
func (m MockFileInfo) ModTime() time.Time { return time.Now() }
func (m MockFileInfo) IsDir() bool        { return m.IsDirectory }
func (m MockFileInfo) Sys() interface{}   { return nil }

// MockBOMFile represents a mock for
// all BOM specific file operations.
type MockBOMFile struct {
	mock.Mock
}

// Stat mocks the os.Stat function.
func (m *MockBOMFile) Stat(name string) (fs.FileInfo, error) {
	arguments := m.Called(name)
	return arguments.Get(0).(fs.FileInfo), arguments.Error(1)
}

// ReadFile mocks the ioutil.ReadFile function.
func (m *MockBOMFile) ReadFile(filename string) ([]byte, error) {
	arguments := m.Called(filename)
	return arguments.Get(0).([]byte), arguments.Error(1)
}
