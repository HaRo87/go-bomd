package mimic

import (
	"io/fs"
	"os"
	"text/template"
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

// MockFile represents a mock for
// all file operations.
type MockFile struct {
	mock.Mock
}

// Stat mocks the os.Stat function.
func (m *MockFile) Stat(name string) (fs.FileInfo, error) {
	arguments := m.Called(name)
	return arguments.Get(0).(fs.FileInfo), arguments.Error(1)
}

// ReadFile mocks the ioutil.ReadFile function.
func (m *MockFile) ReadFile(filename string) ([]byte, error) {
	arguments := m.Called(filename)
	return arguments.Get(0).([]byte), arguments.Error(1)
}

// Create mocks the os.Create function.
func (m *MockFile) Create(name string) (*os.File, error) {
	arguments := m.Called(name)
	return arguments.Get(0).(*os.File), arguments.Error(1)
}

// MockTemplate represents a mock for
// all template operations.
type MockTemplate struct {
	mock.Mock
}

// ParseFiles mocks the template.ParseFiles function.
func (m *MockTemplate) ParseFiles(filenames ...string) (*template.Template, error) {
	arguments := m.Called(filenames)
	return arguments.Get(0).(*template.Template), arguments.Error(1)
}
