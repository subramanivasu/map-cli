package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "map",
	Short: "A CLI tool for map operations using Mappls and Nominatim APIs",
	Long: `map-cli is a command-line tool for performing various map-related operations
such as calculating aerial distance between locations, exploring nearby places,
and finding routes.

Example usage:
  map arieldist 'Yelehanka' 'Koramangala'
  map arieldist --from '13.115;77.607' --to '12.935; 77.624'
  map arieldist 'Yelahanka' --to '12.935; 77.624'`,
	Version: "0.1.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("{{.Name}} {{.Version}}\n")
}
