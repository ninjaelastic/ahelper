package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetAbsolutePath(t *testing.T) {
	// Test with a relative path
	relPath := "testfile.txt"
	absPath, err := GetAbsolutePath(relPath)
	if err != nil {
		t.Fatalf("GetAbsolutePath failed: %v", err)
	}
	if !filepath.IsAbs(absPath) {
		t.Errorf("Expected absolute path, got %s", absPath)
	}

	// Test with an absolute path
	currentDir, _ := os.Getwd()
	absInputPath := filepath.Join(currentDir, "testfile.txt")
	absPath, err = GetAbsolutePath(absInputPath)
	if err != nil {
		t.Fatalf("GetAbsolutePath failed: %v", err)
	}
	if absPath != absInputPath {
		t.Errorf("Expected %s, got %s", absInputPath, absPath)
	}

	// Test with an empty path
	_, err = GetAbsolutePath("")
	if err == nil {
		t.Error("GetAbsolutePath should fail with empty path")
	}
}

func TestIsDirectory(t *testing.T) {
	// Test with a directory
	tempDir, err := os.MkdirTemp("", "utils_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	isDir, err := IsDirectory(tempDir)
	if err != nil {
		t.Fatalf("IsDirectory failed: %v", err)
	}
	if !isDir {
		t.Errorf("Expected %s to be a directory", tempDir)
	}

	// Test with a file
	tempFile, err := os.CreateTemp("", "utils_test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	isDir, err = IsDirectory(tempFile.Name())
	if err != nil {
		t.Fatalf("IsDirectory failed: %v", err)
	}
	if isDir {
		t.Errorf("Expected %s to not be a directory", tempFile.Name())
	}

	// Test with a non-existent path
	isDir, err = IsDirectory("/path/that/does/not/exist")
	if err == nil {
		t.Error("IsDirectory should fail with non-existent path")
	}
}

func TestSplitFileName(t *testing.T) {
	testCases := []struct {
		input         string
		expectedName  string
		expectedExt   string
	}{
		{"file.txt", "file", ".txt"},
		{"file", "file", ""},
		{"file.tar.gz", "file.tar", ".gz"},
		{".hidden", ".hidden", ""},
		{"path/to/file.txt", "file", ".txt"},
	}

	for _, tc := range testCases {
		name, ext := SplitFileName(tc.input)
		if name != tc.expectedName || ext != tc.expectedExt {
			t.Errorf("SplitFileName(%q) = (%q, %q), expected (%q, %q)", tc.input, name, ext, tc.expectedName, tc.expectedExt)
		}
	}
}