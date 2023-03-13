package mimic

import (
	"text/template"

	"github.com/stretchr/testify/mock"
)

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
