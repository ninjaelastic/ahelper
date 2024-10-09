package processor

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/ninjaelastic/ahelpers/internal/filter"
)

func (p *Processor) processDirectory(path string) error {
	if filter.IsIgnored(path, p.ignorePatterns) {
		return nil
	}

	if p.recursive {
		return filepath.WalkDir(path, p.walkDirFunc)
	}

	entries, err := fs.ReadDir(os.DirFS(path), ".")
	if err != nil {
		return fmt.Errorf("error reading directory %s: %v", path, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if err := p.processFile(filepath.Join(path, entry.Name())); err != nil {
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
			return fs.SkipDir
		}
		return nil
	}

	if !d.IsDir() {
		return p.processFile(path)
	}

	return nil
}