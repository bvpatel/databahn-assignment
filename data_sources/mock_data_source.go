package data_sources

import (
	"github.com/stretchr/testify/mock"
)

// MockDataSource is a mock data source for testing.
type MockDataSource struct {
	mock.Mock
}

func (m *MockDataSource) PushData(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

// MockUtils is a mock utility for testing.
type MockUtils struct {
	mock.Mock
}
