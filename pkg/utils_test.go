package pkg

import (
	"testing"
)

func TestCommandExists(t *testing.T) {
	tests := []struct {
		name       string
		command    string
		expectBool bool
	}{
		{
			name:       "Valid command",
			command:    "ls",
			expectBool: true,
		},
		{
			name:       "Invalid command",
			command:    "invalidcommand",
			expectBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if CommandExists(tt.command) != tt.expectBool {
				t.Errorf("expected %v for command %s", tt.expectBool, tt.command)
			}
		})
	}
}
func TestParseCommand(t *testing.T) {
	tests := []struct {
		name              string
		inputs            []string
		expectedGlob      string
		expectedDirectory string
		expectedDays      int
		expectError       bool
	}{
		{
			name:              "Valid input with glob only",
			inputs:            []string{"files", "named", "*"},
			expectedGlob:      "*",
			expectedDirectory: ".",
			expectedDays:      0,
			expectError:       false,
		},
		{
			name:              "Valid input with glob and directory",
			inputs:            []string{"files", "named", "*.txt", "in", "/var/log"},
			expectedGlob:      "*.txt",
			expectedDirectory: "/var/log",
			expectedDays:      0,
			expectError:       false,
		},
		{
			name:              "Valid input with glob, directory, and days",
			inputs:            []string{"files", "named", "*.log", "in", "/tmp", "older", "than", "7", "days"},
			expectedGlob:      "*.log",
			expectedDirectory: "/tmp",
			expectedDays:      7,
			expectError:       false,
		},
		{
			name:              "Valid input with unquoted glob",
			inputs:            []string{"files", "named", `file.*`},
			expectedGlob:      "file.*",
			expectedDirectory: ".",
			expectedDays:      0,
			expectError:       false,
		},
		{
			name:              "Invalid input: missing glob",
			inputs:            []string{"files", "named"},
			expectedGlob:      "",
			expectedDirectory: "",
			expectedDays:      0,
			expectError:       true,
		},
		{
			name:              "Invalid input: missing 'named' keyword",
			inputs:            []string{"files", `"*"`},
			expectedGlob:      "",
			expectedDirectory: "",
			expectedDays:      0,
			expectError:       true,
		},
		{
			name:              "Invalid input: invalid days value",
			inputs:            []string{"files", "named", `"*"`, "older", "than", "invalid", "days"},
			expectedGlob:      "",
			expectedDirectory: "",
			expectedDays:      0,
			expectError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			glob, directory, days, err := ParseFind(tt.inputs)

			// Check for expected error
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if glob != tt.expectedGlob {
				t.Errorf("glob: got %q, want %q", glob, tt.expectedGlob)
			}

			if directory != tt.expectedDirectory {
				t.Errorf("directory: got %q, want %q", directory, tt.expectedDirectory)
			}

			if days != tt.expectedDays {
				t.Errorf("days: got %d, want %d", days, tt.expectedDays)
			}
		})
	}
}
