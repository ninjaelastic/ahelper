package processor

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ninjaelastic/ahelper/internal/filter"
)

func (p *Processor) processDirectory(path string) error {
	// Check if the path exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", path)
	}
	if err != nil {
		return fmt.Errorf("error accessing path %s: %v", path, err)
	}

	fmt.Printf("Processing directory: %s\n", path)
	if filter.IsIgnored(path, p.ignorePatterns) {
		fmt.Printf("Skipping ignored directory: %s\n", path)
		return nil
	}

	if p.recursive {
		return filepath.WalkDir(path, p.walkDirFunc)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %v", path, err)
	}

	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			continue
		}
		if filter.IsIgnored(entryPath, p.ignorePatterns) {
			fmt.Printf("Skipping ignored file: %s\n", entryPath)
			continue
		}
		if len(p.includePatterns) > 0 && !filter.IsIncluded(entryPath, p.includePatterns) {
			fmt.Printf("Skipping file: %s (did not match include pattern)\n", entryPath)
			continue
		}
		if err := p.processFile(entryPath); err != nil {
			return err
		}
	}

	return nil
}

func (p *Processor) processDirectory2(path string) error {
	fmt.Printf("Processing directory: %s\n", path)
	if filter.IsIgnored(path, p.ignorePatterns) {
		fmt.Printf("Skipping ignored directory: %s\n", path)
		return nil
	}

	if p.recursive {
		return filepath.WalkDir(path, p.walkDirFunc)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %v", path, err)
	}

	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			continue
		}
		if filter.IsIgnored(entryPath, p.ignorePatterns) {
			fmt.Printf("Skipping ignored file: %s\n", entryPath)
			continue
		}
		if len(p.includePatterns) > 0 && !filter.IsIncluded(entryPath, p.includePatterns) {
			fmt.Printf("Skipping file: %s (did not match include pattern)\n", entryPath)
			continue
		}
		if err := p.processFile(entryPath); err != nil {
			return err
		}
	}

	return nil
}

func (p *Processor) walkDirFunc(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return fmt.Errorf("error accessing path %s: %v", path, err)
	}

	if filter.IsIgnored(path, p.ignorePatterns) {
		if d.IsDir() {
			fmt.Printf("Skipping ignored directory: %s\n", path)
			return fs.SkipDir
		}
		fmt.Printf("Skipping ignored file: %s\n", path)
		return nil
	}

	if !d.IsDir() {
		if len(p.includePatterns) > 0 && !filter.IsIncluded(path, p.includePatterns) {
			fmt.Printf("Skipping file: %s (did not match include pattern)\n", path)
			return nil
		}
		return p.processFile(path)
	}

	fmt.Printf("Entering directory: %s\n", path)
	return nil
}
