package pkg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestShowDiskSpace(t *testing.T) {
	diskService := NewDisk()
	_, err := diskService.ShowDiskSpace()
	if err != nil {
		t.Errorf("ShowDiskSpace failed: %v", err)
	}
}

func TestShowFolderSize(t *testing.T) {
	tmpDir := t.TempDir()

	filePath := filepath.Join(tmpDir, "testfile.txt")
	err := os.WriteFile(filePath, []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	diskService := NewDisk()
	_, err = diskService.ShowFolderSize(tmpDir)
	if err != nil {
		t.Errorf("ShowFolderSize failed: %v", err)
	}
}

func TestShowFolderSizeWithLimit(t *testing.T) {
	tmpDir := t.TempDir()

	// large file in the temporary directory
	largeFilePath := filepath.Join(tmpDir, "largefile.txt")
	err := os.WriteFile(largeFilePath, make([]byte, 150*1024*1024), 0644) // 150MB
	if err != nil {
		t.Fatalf("Failed to create large test file: %v", err)
	}

	// small file in the temporary directory
	smallFilePath := filepath.Join(tmpDir, "smallfile.txt")
	err = os.WriteFile(smallFilePath, []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Failed to create small test file: %v", err)
	}

	diskService := NewDisk()
	content, err := diskService.ShowFolderSizeWithLimit(tmpDir, 100*1024*1024) // 100MB
	if err != nil {
		t.Errorf("ShowFolderSizeWithLimit failed: %v", err)
	}

	if len(content) != 1 {
		t.Errorf("Expected 1 file, got %d", len(content))
	}
}

func TestReplaceText(t *testing.T) {
	// Create a temporary file for testing
	content := "Hello, World!\nHello, Go!"
	tmpfile, err := os.CreateTemp("", "test.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpfile.Close()

	service := NewInputOutput()

	// Test case 1: Valid replacement
	err = service.ReplaceText(tmpfile.Name(), "Hello", "Hi")
	if err != nil {
		t.Errorf("ReplaceText returned an error: %v", err)
	}

	// Read the file to verify the replacement
	updatedContent, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	expected := "Hi, World!\nHi, Go!"
	if string(updatedContent) != expected {
		t.Errorf("ReplaceText resulted in %q, expected %q", string(updatedContent), expected)
	}

	// Test case 2: No occurrences of oldText
	err = service.ReplaceText(tmpfile.Name(), "Nonexistent", "Replacement")
	if err == nil {
		t.Error("Expected an error for no occurrences of oldText, but got none")
	}

	// Test case 3: Non-existent file
	err = service.ReplaceText("nonexistent.txt", "Hello", "Hi")
	if err == nil {
		t.Error("Expected an error for non-existent file, but got none")
	}
}
