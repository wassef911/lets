package pkg

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type InputOutputServiceInterface interface {
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
		return errors.New("failed to open file: " + err.Error())
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
			return errors.New("failed to read CSV: " + err.Error())
		}

		if columnIndex < 0 || columnIndex >= len(record) {
			return errors.New("column index " + strconv.Itoa(columnIndex) + "is out of bounds")
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
		return errors.New("failed to read file: " + err.Error())
	}

	content := string(file)
	updatedContent := strings.ReplaceAll(content, oldText, newText)

	if content == updatedContent {
		return errors.New("no occurrences of " + oldText + "found in file")
	}

	err = os.WriteFile(filename, []byte(updatedContent), 0644)
	if err != nil {
		return errors.New("failed to write file: " + err.Error())
	}

	return nil
}
