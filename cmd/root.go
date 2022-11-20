package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configFile string
var file string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bomd",
	Short: "Convert BOM to markdown",
	Long: `go-bomd can read in Software Bill Of Materials (SBOMs)
	based on the CycloneDX standard and convert relevant information
	into markdown based documents using custom templates.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yml", "config file (default ./config.yml)")
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "the file on which an operation should be performed")
}
