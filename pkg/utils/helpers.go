package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func EnsureDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func SplitFileName(filename string) (name, ext string) {
	name = filepath.Base(filename)
	if strings.HasPrefix(name, ".") && !strings.Contains(name[1:], ".") {
		return name, ""
	}
	ext = filepath.Ext(name)
	name = name[:len(name)-len(ext)]
	return
}

func IsDirectory(path string) (bool, error) {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

func GetAbsolutePath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("error getting absolute path: %w", err)
	}

	return absPath, nil
}
