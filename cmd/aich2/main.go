package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/ninjaelastic/ahelper/internal/filter"
	"github.com/ninjaelastic/ahelper/internal/processor"
)

const (
	programName = "aich"
	version     = "1.0.0"
)

var (
	styleHeading    = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).Padding(1, 0, 1, 0)
	styleSubHeading = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Italic(true)
	styleOption     = lipgloss.NewStyle().Foreground(lipgloss.Color("220"))
	styleError      = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	styleSuccess    = lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Bold(true)
)

func showLogo() {
	logo := `
88
88
88,dPPYba,
88P'    "8a
88       88
88       88
88       88
   AI Context Helper
`
	fmt.Println(styleHeading.Render(logo))
	fmt.Println(styleSubHeading.Render("Version " + version))
	fmt.Println()
}

func showHelp() {
	showLogo()

	styleHeading := lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	styleSubHeading := lipgloss.NewStyle().Foreground(lipgloss.Color("207")).Bold(true)
	styleOption := lipgloss.NewStyle().Foreground(lipgloss.Color("218"))
	styleDescription := lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	styleExample := lipgloss.NewStyle().Foreground(lipgloss.Color("123"))
	styleNote := lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Italic(true)

	helpText := fmt.Sprintf(`
%s
  AICH (AI Context Helper) is a powerful tool for managing and processing
  file content in AI-assisted programming tasks. It offers recursive searching,
  pattern filtering, and tree structure visualization to streamline your workflow.

%s
  aich [options] <path> [<path> ...]

%s
  %s-r%s    Include subfolders recursively (default: true)
        Use -r=false to disable recursive search

  %s-i%s    Comma-separated list of additional ignore patterns
        Files matching these patterns will be excluded
        Default ignore patterns are already included for common project files
        Example: -i "*.tmp,*.log,custom_folder"

  %s-I%s    Comma-separated list of include patterns
        Only files matching these patterns will be processed
        If empty, all files (except ignored ones) are included
        Example: -I "*.go,*.js,*.py"

  %s-t%s    Display directory tree structure before processing files

%s
  %s1. Process all files in the current directory and subdirectories:
     aich .%s

  %s2. Process specific file types in a project, with additional ignore patterns:
     aich -r -i "*.tmp,custom_folder" -I "*.go,*.js" -t /path/to/project%s

  %s3. Process only specific files without recursion:
     aich -r=false /path/to/file1.go /path/to/file2.js%s

  %s4. Display directory tree without processing files:
     aich -t -r=false /path/to/directory%s

%s
  Default ignore patterns are included for common version control, build outputs,
  and temporary files. Use the -i flag to add your custom ignore patterns.
`,
		styleHeading.Render("DESCRIPTION:"),
		styleHeading.Render("USAGE:"),
		styleHeading.Render("OPTIONS:"),
		styleOption.Render(""), styleDescription.Render(""),
		styleOption.Render(""), styleDescription.Render(""),
		styleOption.Render(""), styleDescription.Render(""),
		styleOption.Render(""), styleDescription.Render(""),
		styleHeading.Render("EXAMPLES:"),
		styleSubHeading.Render(""), styleExample.Render(""),
		styleSubHeading.Render(""), styleExample.Render(""),
		styleSubHeading.Render(""), styleExample.Render(""),
		styleSubHeading.Render(""), styleExample.Render(""),
		styleNote.Render("NOTE:"),
	)

	fmt.Println(helpText)
	fmt.Println(styleOption.Render("For more information and updates, visit: https://github.com/ninjaelastic/aich"))
}

func Run(recursive bool, ignorePatterns, includePatterns string, showTree, help bool, args []string) error {
	if help {
		showHelp()
		return nil
	}

	if len(args) == 0 {
		showHelp()
		fmt.Println(styleError.Render("Error: No paths provided."))
		fmt.Println("Run 'aich -h' for usage information.")
		return fmt.Errorf("no paths provided")
	}

	additionalIgnorePatterns := filter.SplitPatterns(ignorePatterns)
	proc := processor.New(
		recursive,
		additionalIgnorePatterns,
		filter.SplitPatterns(includePatterns),
		showTree,
	)

	if err := proc.Run(args); err != nil {
		fmt.Println(styleError.Render(fmt.Sprintf("Error: %v", err)))
		return err
	}

	fmt.Println(styleSuccess.Render("Processing complete!"))
	return nil
}

func main() {
	recursive := flag.Bool("r", true, "Include subfolders recursively")
	ignorePatterns := flag.String("i", "", "Comma-separated list of additional ignore patterns")
	includePatterns := flag.String("I", "", "Comma-separated list of include patterns")
	showTree := flag.Bool("t", false, "Display directory tree structure")
	showIgnorePatterns := flag.Bool("show-ignore", false, "Display default ignore patterns")
	help := flag.Bool("h", false, "Show help message")

	flag.Parse()

	if *help || len(flag.Args()) == 0 {
		showHelp()
		return
	}

	showLogo()

	fmt.Printf("Processing with options:\n")
	fmt.Printf("  Recursive: %v\n", *recursive)
	fmt.Printf("  Additional ignore patterns: %s\n", *ignorePatterns)
	fmt.Printf("  Include patterns: %s\n", *includePatterns)
	fmt.Printf("  Show tree: %v\n", *showTree)
	fmt.Printf("  Paths: %v\n", flag.Args())

	if *showIgnorePatterns {
		fmt.Printf("\nDefault ignore patterns:\n")
		for _, pattern := range processor.DefaultIgnorePatterns {
			fmt.Printf("  %s\n", pattern)
		}
		fmt.Printf("\nTo override default ignore patterns, use the -i flag with an empty string: -i \"\"\n\n")
	}
	if err := Run(*recursive, *ignorePatterns, *includePatterns, *showTree, *help, flag.Args()); err != nil {
		fmt.Printf("%s\n", styleError.Render(fmt.Sprintf("Error: %v", err)))
		os.Exit(1)
	}
}
