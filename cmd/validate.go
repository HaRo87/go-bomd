package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	gbom "github.com/HaRo87/go-bomd/bom"
	"github.com/HaRo87/go-bomd/replicator"

	cdx "github.com/CycloneDX/cyclonedx-go"
)

var licenseCheck bool
var listMissingLicenses bool

func validateBOM(bom *cdx.BOM, validateLicenses bool, bomProc gbom.BOMProcessor) (err error) {
	err = bomProc.ValidateBOM(bom)
	if err != nil {
		return
	}
	if validateLicenses {
		logrus.Debugf("Trying to validate BOM component license information")
		err = bomProc.ValidateComponentLicenses(bom)
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
// var validateConfigCmd = &cobra.Command{
// 	Use:   "config",
// 	Short: "Validate a specified config",
// 	Long: `Validate (bomd validate config) will support with checking the integrity
// 	of the specified config.`,
// 	Run: func(cmd *cobra.Command, args []string) {

// 	},
// }

// validateBomCmd represents the validate BOM command
var validateBomCmd = &cobra.Command{
	Use:   "bom",
	Short: "Validate a specified BOM",
	Long: `Validate (bomd validate bom) will support with checking the integrity
	of the specified BOM.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		builder := gbom.NewDefaultBOMProcessorBuilder()
		processor := builder.GetBOMProcessor()
		filePath, err := getFilePath(files, ".json")
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		logrus.Debugf("Trying to read BOM: %s", filePath)
		bom, err := processor.GetBOM(filePath)
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		logrus.Infof("Validating BOM: %s", filePath)
		err = validateBOM(&bom, licenseCheck, processor)
		if err != nil {
			logrus.Error("😱 something went wrong")
		} else {
			logrus.Info("😎 everything seems to be fine")
		}
		if listMissingLicenses {
			logrus.SetLevel(logrus.WarnLevel)
			components, _ := processor.GetComponentsWithEmptyLicenseIDs(&bom)
			for _, component := range components {
				logrus.Warnf("🤔 component: %s is missing license information", component)
			}
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
	RunE: func(cmd *cobra.Command, args []string) error {
		builder := replicator.NewDefaultTemplateProcessorBuilder()
		processor := builder.GetTemplateProcessor()
		filePath, err := getFilePath(files, ".tmpl")
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		logrus.Debugf("Trying to parse template: %s", filePath)
		err = processor.Validate(replicator.TemplateInfo{InputFilePath: filePath})
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		logrus.Info("😎 everything seems to be fine")
		return nil
	},
}

func init() {
	validateBomCmd.Flags().BoolVarP(&licenseCheck, "licenseCheck", "l", false, "check if license info is present (default false)")
	validateBomCmd.Flags().BoolVar(&listMissingLicenses, "listMissing", false, "list all components missing license info (default false)")
	validateCmd.AddCommand(validateBomCmd)
	//validateCmd.AddCommand(validateConfigCmd)
	validateCmd.AddCommand(validateTemplateCmd)
	rootCmd.AddCommand(validateCmd)
}
