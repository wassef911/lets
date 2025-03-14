package pkg

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func findIndex(inputs []string, keyword string) int {
	for i, part := range inputs {
		if part == keyword {
			return i
		}
	}
	return -1
}

// commandExists checks if a command is available on machine
func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func ParseFind(inputs []string) (glob string, directory string, days int, err error) {
	// extract the glob (required)
	namedIndex := findIndex(inputs, "named")
	if namedIndex == -1 {
		err = errors.New("invalid input format: 'named' keyword not found " + strings.Join(inputs, " "))
		return
	}
	if namedIndex+1 >= len(inputs) {
		err = errors.New("invalid input format: 'named' keyword not found " + strings.Join(inputs, " "))
		return
	}
	glob = inputs[namedIndex+1]

	// extract the directory (optional)
	directory = "."
	inIndex := findIndex(inputs, "in")
	if inIndex != -1 && inIndex+1 < len(inputs) {
		directory = inputs[inIndex+1]
	}

	// find the index of the "older" keyword (optional)
	days = 0
	olderIndex := findIndex(inputs, "older")
	if olderIndex != -1 && inputs[olderIndex+1] == "than" { // default to days for now
		daysStr := inputs[olderIndex+2]
		days, err = strconv.Atoi(daysStr)
		if err != nil {
			err = errors.New("invalid days value: " + err.Error())
			return
		}
	}

	return
}
