package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	gbom "gitlab.com/HaRo87go-bomd/bom"
)

var licenseCheck bool

func validateItem(config string) {
	fmt.Println("Validating ...")
}

func validateBOM(bomFile string, validateLicenses bool) (err error) {
	builder := gbom.NewDefaultBOMProcessorBuilder()
	processor := builder.GetBOMProcessor()
	logrus.Debugf("Trying to read BOM: %s", bomFile)
	bom, err := processor.GetBOM(bomFile)
	if err != nil {
		return
	}
	logrus.Debugf("Trying to validate BOM")
	err = processor.ValidateBOM(&bom)
	if err != nil {
		return
	}
	if validateLicenses {
		logrus.Debugf("Trying to validate BOM component license information")
		err = processor.ValidateComponentLicenses(&bom)
		if err != nil {
			return
		}
	}
	return
}

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a specified item",
	Long: `Validate (bomd validate) will support with checking the integrity
	of the specified item.`,
}

// validateConfigCmd represents the validate config command
var validateConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Validate a specified config",
	Long: `Validate (bomd validate config) will support with checking the integrity
	of the specified config.`,
	Run: func(cmd *cobra.Command, args []string) {
		validateItem(args[0])
	},
}

// validateBomCmd represents the validate BOM command
var validateBomCmd = &cobra.Command{
	Use:   "bom",
	Short: "Validate a specified BOM",
	Long: `Validate (bomd validate bom) will support with checking the integrity
	of the specified BOM.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Infof("Validating BOM: %s", file)
		err := validateBOM(file, licenseCheck)
		if err != nil {
			logrus.Error("😱 something went wrong")
		} else {
			logrus.Info("😎 everything seems to be fine")
		}
		return err
	},
}

// validateTemplateCmd represents the validate template command
var validateTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Validate a specified template",
	Long: `Validate (bomd validate template) will support with checking the integrity
	of the specified template.`,
	Run: func(cmd *cobra.Command, args []string) {
		validateItem(args[0])
	},
}

func init() {
	validateBomCmd.Flags().BoolVarP(&licenseCheck, "license-check", "l", false, "check if license info is present (default false)")
	validateCmd.AddCommand(validateBomCmd)
	validateCmd.AddCommand(validateConfigCmd)
	validateCmd.AddCommand(validateTemplateCmd)
	rootCmd.AddCommand(validateCmd)
}
