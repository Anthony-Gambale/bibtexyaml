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
	Short: "Searches for a file called `new.bib`, converts all entries to YAML, and adds them to the YAML file with the given name.",
	RunE:  runReverse,
}

func runReverse(cmd *cobra.Command, args []string) error {
	fileName := args[0]

	// Verify file name format
	if !isYamlFile(fileName) {
		return fmt.Errorf("file name must have .yaml extension")
	}

	// Check if new.bib exists
	exists, err := fileExists("new.bib")
	if err != nil {
		return fmt.Errorf("failed to check file new.bib: %w", err)
	}
	if !exists {
		return fmt.Errorf("file new.bib does not exist")
	}

	// Open new.bib
	// 0644 corresponds to permissions -rw-r--r--
	newDotBib, err := os.OpenFile("new.bib", os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %w", err)
	}
	defer newDotBib.Close()

	// Read new.bib
	content, err := io.ReadAll(newDotBib)
	if err != nil {
		return fmt.Errorf("failed to read file %w", err)
	}

	// Parse the BibTeX
	bib, err := bibtex.Parse(strings.NewReader(string(content)))
	if err != nil {
		fmt.Printf("Error parsing BibTeX: %v\n", err)
		return nil
	}

	// Convert to YAML and get a string
	newYaml := convertToYAML(bib)

	// Open the YAML file (create if doesn't exist)
	// 0644 corresponds to permissions -rw-r--r--
	yamlFile, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	defer yamlFile.Close()

	// Read YAML file contents
	content, _ = io.ReadAll(yamlFile)

	// If YAML file is empty, add extra line
	if len(content) == 0 {
		yamlFile.WriteString("entries:\n")
	}

	// Append the new items
	yamlFile.Seek(0, 2) // Move cursor to end of file
	yamlFile.WriteString(newYaml)

	// Empty all contents from new.bib (but don't delete the file)
	err = newDotBib.Truncate(0)
	if err != nil {
		return fmt.Errorf("failed to clear new.bib: %w", err)
	}

	return nil
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func convertToYAML(bib *bibtex.BibTex) string {
	var result strings.Builder
	for _, entry := range bib.Entries {
		result.WriteString(fmt.Sprintf("  - id: %s\n", entry.CiteName))
		result.WriteString(fmt.Sprintf("    type: %s\n", entry.Type))
		result.WriteString("    fields:\n")
		for key, value := range entry.Fields {
			result.WriteString(fmt.Sprintf("      %s: %s\n", key, value))
		}
		result.WriteString("\n")
	}
	return result.String()
}
