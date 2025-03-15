package pkg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestShowDiskSpace(t *testing.T) {
	diskService := NewDisk()
	err := diskService.ShowDiskSpace()
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
	err = diskService.ShowFolderSize(tmpDir)
	if err != nil {
		t.Errorf("ShowFolderSize failed: %v", err)
	}
}

func TestShowFolderSizeWithLimit(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a large file in the temporary directory
	largeFilePath := filepath.Join(tmpDir, "largefile.txt")
	err := os.WriteFile(largeFilePath, make([]byte, 150*1024*1024), 0644) // 150MB
	if err != nil {
		t.Fatalf("Failed to create large test file: %v", err)
	}

	// Create a small file in the temporary directory
	smallFilePath := filepath.Join(tmpDir, "smallfile.txt")
	err = os.WriteFile(smallFilePath, []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Failed to create small test file: %v", err)
	}

	diskService := NewDisk()
	err = diskService.ShowFolderSizeWithLimit(tmpDir, 100*1024*1024) // 100MB
	if err != nil {
		t.Errorf("ShowFolderSizeWithLimit failed: %v", err)
	}
}
