package app

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"

	"github.com/wassef911/lets/pkg"
)

var logger = pkg.NewLogger()

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

type MockDiskService struct {
	mock.Mock
}

func (m *MockDiskService) ShowDiskSpace() ([]string, error) {
	fmt.Println("MockDiskService - ------------------------------------------")
	m.Called()
	return []string{}, nil
}
func (m *MockDiskService) ShowFolderSize(path string) (string, error) {
	m.Called(path)
	return "", nil
}
func (m *MockDiskService) ShowFolderSizeWithLimit(dir string, minSize float64) ([]string, error) {
	m.Called(dir, minSize)
	return []string{}, nil
}

type mockProcService struct {
	mock.Mock
}

func (m *mockProcService) Processes() error {
	m.Called()
	return nil
}

func (m *mockProcService) Resources() error {
	m.Called()
	return nil
}

func (m *mockProcService) KillProcessByName(processName string) error {
	m.Called(processName)
	return nil
}

type mockSearchService struct {
	mock.Mock
}

func (m *mockSearchService) SearchFiles(pattern, directory string) error {
	m.Called(pattern, directory)
	return nil
}
func (m *mockSearchService) CountMatches(pattern, filename string) error {
	m.Called(pattern, filename)
	return nil
}

func (m *mockSearchService) FindFiles(glob, directory string, days int) error {
	m.Called(glob, directory, days)
	return nil
}

func TestNewGetCmd(t *testing.T) {
	mockService := &MockInputOutputService{}
	cmd := NewGetCmd(logger, mockService)
	mockService.On("GetColumn", "test.csv", 2).Once().Return(nil)

	cmd.Run(&cobra.Command{}, []string{"column", "2", "from", "test.csv"})

	mockService.AssertExpectations(t)
}

func TestReplaceCmd(t *testing.T) {
	mockService := &MockInputOutputService{}
	cmd := NewReplaceCmd(logger, mockService)
	mockService.On("ReplaceText", "test.txt", "foo", "bar").Once().Return(nil)

	cmd.Run(&cobra.Command{}, []string{"foo", "with", "bar", "in", "test.txt"})

	mockService.AssertExpectations(t)
}

func TestRootCmd(t *testing.T) {
	mockDiskService := &MockDiskService{}
	mockProcService := &mockProcService{}
	mockSearchService := &mockSearchService{}
	mockInputOutputService := &MockInputOutputService{}

	mockDiskService.On("ShowDiskSpace").Once().Return(nil)

	mockProcService.On("Processes").Once().Return(nil)
	mockProcService.On("Resources").Once().Return(nil)

	cmd := NewRootCmd(logger, mockDiskService, mockInputOutputService, mockProcService, mockSearchService)
	cmd.Run(&cobra.Command{}, []string{})
	cmd.SetArgs([]string{"show", "disk", "space"})
	cmd.Execute()

	cmd.SetArgs([]string{"show", "proc", "processes"})
	cmd.Execute()

	cmd.SetArgs([]string{"show", "proc", "resources"})
	cmd.Execute()

	mockDiskService.AssertExpectations(t)

}
