package cmd

import (
	"fmt"
	"index/suffixarray"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	returnAllIndices        = -1
	defaultTemplateSuffix   = ".tmpl"
	defaultBomSuffix        = ".json"
	defaultResultSuffix     = ".md"
	defaultTemplateFilePath = "template.tmpl"
	defaultBomFilePath      = "bom.json"
	defaultResultFilePath   = "result.md"
)

type bomdFiles struct {
	InputFilePath    string
	TemplateFilePath string
	OutputFilePath   string
}

type order string

const (
	Input    order = "input"
	Template order = "template"
	Output   order = "output"
)

// var configFile string
var (
	files        []string
	ignoreErrors bool
	logLevel     int
	fileOrder    []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bomd",
	Short: "Convert BOM to markdown",
	Long: `go-bomd can read in Software Bill Of Materials (SBOMs)
	based on the CycloneDX standard and convert relevant information
	into markdown based documents using custom templates.`,
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil && !ignoreErrors {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogger)
	//rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yml", "config file (default ./config.yml)")
	rootCmd.PersistentFlags().StringArrayVarP(&files, "file", "f", []string{}, "the file(s) on which an operation should be performed")
	rootCmd.PersistentFlags().StringArrayVar(&fileOrder, "order", []string{"input", "template", "output"}, "the order of the file(s) provided via the --file flag")
	rootCmd.PersistentFlags().BoolVar(&ignoreErrors, "ignore-errors", false, "do not error out")
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

func getDefaultFilePath(suffix string) (filePath string, err error) {
	if suffix == defaultTemplateSuffix {
		filePath = defaultTemplateFilePath
	}
	if suffix == defaultBomSuffix {
		filePath = defaultBomFilePath
	}
	if suffix == defaultResultSuffix {
		filePath = defaultResultFilePath
	}
	if len(filePath) == 0 {
		err = fmt.Errorf("unable to identify default file path for suffix: %s", suffix)
	}
	return
}

func getOrderlyFilePath(files []string, orderElem Order) (filePath string, err error) {
	if len(files) == len(order) {
		for i, elem := range order {
			if elem == string(orderElem) {
				filePath = files[i]
			}
		}
	} else {
		err = fmt.Errorf("provided files: %s do not match defined order: %s", files, order)
	}
	return
}

func getFilePath(files []string, suffix string) (filePath string, err error) {

}

func getOrderlyFiles(files []string) (orderlyFiles bomdFiles, err error) {
	suffixes := []string{defaultBomSuffix, defaultResultSuffix, defaultTemplateSuffix}
	data := []byte("\x00" + strings.Join(files, "\x00") + "\x00")
	suffixFiles := suffixarray.New(data)
	for _, suffix := range suffixes {
		fileIndices := suffixFiles.Lookup([]byte(suffix), returnAllIndices)
		if len(fileIndices) > 1 {

		} else if len(fileIndices) > 0 {
			filePath = files[fileIndices[0]]
		}
	}

	if len(filePath) == 0 {
		filePath, err = getDefaultFilePath(suffix)
	}
	return
}
