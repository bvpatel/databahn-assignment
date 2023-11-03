package util

import "github.com/stretchr/testify/mock"

// MockUtils is a mock utility for testing.
type MockUtils struct {
	mock.Mock
}

func (m *MockUtils) ReadTemplateFile(filePath string) (string, error) {
	args := m.Called(filePath)
	return args.String(0), args.Error(1)
}

func (m *MockUtils) RenderTemplate(template string, data interface{}) (string, error) {
	args := m.Called(template, data)
	return args.String(0), args.Error(1)
}
