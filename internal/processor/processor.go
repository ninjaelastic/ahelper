package processor

type Processor struct {
    recursive       bool
    ignorePatterns  []string
    includePatterns []string
    showTree        bool
}

func New(recursive bool, ignorePatterns, includePatterns []string, showTree bool) *Processor {
    return &Processor{
        recursive:       recursive,
        ignorePatterns:  ignorePatterns,
        includePatterns: includePatterns,
        showTree:        showTree,
    }
}

func (p *Processor) Run(paths []string) error {
    for _, path := range paths {
        if err := p.process(path); err != nil {
            return err
        }
    }
    return nil
}

func (p *Processor) process(path string) error {
    // Implement the processing logic here
    return nil
}