package processor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ninjaelastic/ahelper/internal/filter"
)

func (p *Processor) printTree(root string) error {
	fmt.Println("Directory tree structure:")
	return p.printTreeHelper(root, "", true)
}

func (p *Processor) printTreeHelper(path string, prefix string, isLast bool) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	fmt.Print(prefix)
	if isLast {
		fmt.Print("└── ")
		prefix += "    "
	} else {
		fmt.Print("├── ")
		prefix += "│   "
	}
	fmt.Println(fileInfo.Name())

	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		var filteredEntries []os.DirEntry
		for _, entry := range entries {
			if !filter.IsIgnored(filepath.Join(path, entry.Name()), p.ignorePatterns) {
				filteredEntries = append(filteredEntries, entry)
			}
		}

		for i, entry := range filteredEntries {
			err := p.printTreeHelper(filepath.Join(path, entry.Name()), prefix, i == len(filteredEntries)-1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
