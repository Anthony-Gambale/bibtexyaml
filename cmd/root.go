package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "BibTeX-YAML",
	Short: "Convert YAML into BibTeX and vice versa.",
	Long:  "Convert YAML into BibTeX and vice versa. See https://github.com/Anthony-Gambale/BibTeX-YAML for more information.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
