package utils

import (
	"os"
	"path/filepath"
)

func EnsureDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

func GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}
