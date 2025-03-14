package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type InputOutputInterface interface {
	GetColumn(filename string, columnIndex int) error
	ReplaceText(filename, oldText, newText string) error
}

type InputOutputService struct{}

func NewInputOutput() *InputOutputService {
	return &InputOutputService{}
}

// extracts a specific column from a CSV file (equivalent to `awk '{print $N}'`).
func (inou *InputOutputService) GetColumn(filename string, columnIndex int) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var columnData []string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV: %v", err)
		}

		if columnIndex < 0 || columnIndex >= len(record) {
			return fmt.Errorf("column index %d is out of bounds", columnIndex)
		}

		columnData = append(columnData, record[columnIndex])
	}

	fmt.Println(strings.Join(columnData, "\n"))
	return nil
}

// performs in-place text replacement in a file (equivalent to `sed -i 's/old/new/g'`).
func (inou *InputOutputService) ReplaceText(filename, oldText, newText string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	content := string(file)
	updatedContent := strings.ReplaceAll(content, oldText, newText)

	if content == updatedContent {
		return fmt.Errorf("no occurrences of '%s' found in file", oldText)
	}

	err = os.WriteFile(filename, []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
