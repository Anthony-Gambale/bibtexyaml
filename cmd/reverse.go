package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nickng/bibtex"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(reverseTemplateCmd)
}

var reverseTemplateCmd = &cobra.Command{
	Use:   "reverse [filename]",
	Short: "Generate a YAML file from a BibTeX file",
	Args:  cobra.ExactArgs(1),
	RunE:  runReverse,
}

func runReverse(cmd *cobra.Command, args []string) error {
	fileName := args[0]

	// Verify file name format
	if !isBibtexFile(fileName) {
		return fmt.Errorf("file name must have .bib extension")
	}

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file %w", err)
	}
	defer file.Close()

	// Read file
	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file %w", err)
	}

	// Parse the BibTeX
	bib, err := bibtex.Parse(strings.NewReader(string(content)))
	if err != nil {
		fmt.Printf("Error parsing BibTeX: %v\n", err)
		return nil
	}

	// Iterate over the entries and print them
	for _, entry := range bib.Entries {
		fmt.Printf("Entry Type: %s\n", entry.Type)
		fmt.Printf("Citation Key: %s\n", entry.CiteName)
		fmt.Println("Fields:")
		for key, value := range entry.Fields {
			fmt.Printf("  %s = %s\n", key, value)
		}
		fmt.Println()
	}

	return nil
}

func isBibtexFile(name string) bool {
	if len(name) < 5 {
		return false
	}
	return name[len(name)-4:] == ".bib"
}
