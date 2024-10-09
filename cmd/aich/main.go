package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ninjaelastic/ahelper/internal/processor"
	"github.com/ninjaelastic/ahelpers/internal/filter"
)

func main() {
	// Parse command-line flags
	recursive := flag.Bool("r", true, "Include subfolders recursively")
	ignorePatterns := flag.String("i", "*.tmp,*.log,node_modules,.git", "Comma-separated list of ignore patterns")
	includePatterns := flag.String("I", "", "Comma-separated list of include patterns")
	showTree := flag.Bool("t", false, "Display directory tree structure")
	flag.Parse()

	// Get paths from remaining arguments
	paths := flag.Args()
	if len(paths) == 0 {
		fmt.Println("No paths provided.")
		flag.Usage()
		os.Exit(1)
	}

	// Create processor with parsed options
	proc := processor.New(
		*recursive,
		filter.SplitPatterns(*ignorePatterns),
		filter.SplitPatterns(*includePatterns),
		*showTree,
	)

	// Process paths
	if err := proc.Run(paths); err != nil {
		log.Fatal(err)
	}
}