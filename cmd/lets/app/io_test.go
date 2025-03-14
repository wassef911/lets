package app

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

type MockInputOutputService struct {
	mock.Mock
}

func (m *MockInputOutputService) GetColumn(filename string, columnIndex int) error {
	m.Called(filename, columnIndex)
	return nil
}
func (m *MockInputOutputService) ReplaceText(filename, oldText, newText string) error {
	m.Called(filename, oldText, newText)
	return nil
}

func TestNewGetCmd(t *testing.T) {
	mockService := &MockInputOutputService{}
	cmd := newGetCmd(mockService)
	mockService.On("GetColumn", "test.csv", 2).Once().Return(nil)

	cmd.Run(&cobra.Command{}, []string{"column", "2", "from", "test.csv"})

	mockService.AssertExpectations(t)
}

func TestReplaceCmd(t *testing.T) {
	mockService := &MockInputOutputService{}
	cmd := newReplaceCmd(mockService)
	mockService.On("ReplaceText", "test.txt", "foo", "bar").Once().Return(nil)

	cmd.Run(&cobra.Command{}, []string{"foo", "with", "bar", "in", "test.txt"})

	mockService.AssertExpectations(t)
}
