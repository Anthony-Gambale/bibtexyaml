package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type YAMLBibliography struct {
	Entries []BibliographyEntry `yaml:"entries"`
}

type BibliographyEntry struct {
	ID     string            `yaml:"id"`
	Type   string            `yaml:"type"`
	Fields map[string]string `yaml:"fields"`
}

func init() {
	rootCmd.AddCommand(templateCmd)
}

var templateCmd = &cobra.Command{
	Use:   "template [filename]",
	Short: "Generate BibTeX from a YAML file",
	Args:  cobra.ExactArgs(1),
	RunE:  runTemplating,
}

func runTemplating(cmd *cobra.Command, args []string) error {
	fileName := args[0]

	// Verify file name format
	if !isYamlFile(fileName) {
		return fmt.Errorf("file name must have .yaml extension")
	}

	// Open file
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

	// Parse the YAML
	var data YAMLBibliography
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return fmt.Errorf("parse error in YAML: %w", err)
	}

	// Generate BibTeX
	outputText, err := convertToBibTeX(&data)
	if err != nil {
		return fmt.Errorf("failed to generate BibTeX: %w", err)
	}

	// File name of output file
	bibFileName := fileName[:len(fileName)-5] + ".bib"

	file, err = os.Create(bibFileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", bibFileName, err)
	}
	defer file.Close()

	_, err = file.Write([]byte(outputText))

	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", bibFileName, err)
	}

	return nil
}

func isYamlFile(name string) bool {
	if len(name) < 5 {
		return false
	}
	return name[len(name)-5:] == ".yaml"
}

func convertToBibTeX(bibliography *YAMLBibliography) (string, error) {
	var result strings.Builder
	for _, entry := range bibliography.Entries {
		result.WriteString(fmt.Sprintf("@%s{%s,\n", entry.Type, entry.ID))
		for key, value := range entry.Fields {
			result.WriteString(fmt.Sprintf("  %s = {%s},\n", key, value))
		}
		result.WriteString("}\n\n")
	}
	return result.String(), nil
}
