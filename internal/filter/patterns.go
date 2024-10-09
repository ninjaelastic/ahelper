package filter

import (
	"path/filepath"
	"strings"
)

func SplitPatterns(patterns string) []string {
	return strings.Split(patterns, ",")
}

func IsIgnored(path string, ignorePatterns []string) bool {
	for _, pattern := range ignorePatterns {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
		if strings.Contains(path, pattern) {
			return true
		}
	}
	return false
}

func IsIncluded(path string, includePatterns []string) bool {
	if len(includePatterns) == 0 {
		return true
	}
	for _, pattern := range includePatterns {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
	}
	return false
}
