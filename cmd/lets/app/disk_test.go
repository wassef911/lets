package app

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

type MockDiskService struct {
	mock.Mock
}

func (m *MockDiskService) ShowDiskSpace() error {
	m.Called()
	return nil
}
func (m *MockDiskService) ShowFolderSize(path string) error {
	m.Called(path)
	return nil
}
func (m *MockDiskService) ShowFolderSizeWithLimit(dir string, minSize float64) error {
	m.Called(dir, minSize)
	return nil
}

func TestNewShowCmd(t *testing.T) {
	mockService := &MockDiskService{}
	cmd := newShowCmd(mockService)
	mockService.On("ShowDiskSpace").Once()

	cmd.Run(&cobra.Command{}, []string{"space"})

	mockService.AssertExpectations(t)
}
