package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configFile string
var file string
var igrnoreErrors bool
var logLevel int

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
	if err != nil && !igrnoreErrors {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogger)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yml", "config file (default ./config.yml)")
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "the file on which an operation should be performed")
	rootCmd.PersistentFlags().BoolVar(&igrnoreErrors, "ignore-errors", false, "do not error out")
	rootCmd.PersistentFlags().CountVarP(&logLevel, "verbose", "v", "logger verbosity")
}

func initLogger() {
	switch logLevel {
	case 0:
		logrus.SetLevel(logrus.ErrorLevel)
	case 1:
		logrus.SetLevel(logrus.WarnLevel)
	case 2:
		logrus.SetLevel(logrus.InfoLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
}
