package cmd

import (
	gbom "github.com/HaRo87/go-bomd/bom"
	"github.com/HaRo87/go-bomd/replicator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a specified item",
	Long:  `Generate (bomd generate) will support with creating a new item.`,
}

// generateConfigCmd represents the generate config command
// var generateConfigCmd = &cobra.Command{
// 	Use:   "config",
// 	Short: "Generate a specified config",
// 	Long:  `Generate (bomd generate config) will support with creating the specified config.`,
// 	Run: func(cmd *cobra.Command, args []string) {

// 	},
// }

// generateResultCmd represents the generate result command
var generateResultCmd = &cobra.Command{
	Use:   "result",
	Short: "Generate a specified result",
	Long:  `Generate (bomd generate result) will create the specified result based on the provided BOM and template.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bomFilePath, err := getFilePath(files, ".json")
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		templateFilePath, err := getFilePath(files, ".tmpl")
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		templateBuilder := replicator.NewDefaultTemplateProcessorBuilder()
		templateProcessor := templateBuilder.GetTemplateProcessor()
		bomBuilder := gbom.NewDefaultBOMProcessorBuilder()
		bomProcessor := bomBuilder.GetBOMProcessor()
		bom, err := bomProcessor.GetBOM(bomFilePath)
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		filePath := ""
		for _, file := range files {
			if file != bomFilePath && file != templateFilePath {
				filePath = file
			}
		}
		logrus.Debugf("Trying to generate result: %s", filePath)
		err = templateProcessor.Execute(replicator.TemplateInfo{
			InputFilePath:  templateFilePath,
			OutputFilePath: filePath,
			BOM:            &bom},
		)
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		logrus.Info("😎 everything seems to be fine")
		return nil
	},
}

// generateTemplateCmd represents the generate template command
var generateTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate a specified template",
	Long:  `Generate (bomd generate template) will support with creating the specified template.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		builder := replicator.NewDefaultTemplateProcessorBuilder()
		processor := builder.GetTemplateProcessor()
		filePath, err := getFilePath(files, ".tmpl")
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		logrus.Debugf("Trying to generate default template: %s", filePath)
		err = processor.Generate(filePath)
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		logrus.Info("😎 everything seems to be fine")
		return nil
	},
}

func init() {
	//generateCmd.AddCommand(generateConfigCmd)
	generateCmd.AddCommand(generateResultCmd)
	generateCmd.AddCommand(generateTemplateCmd)
	rootCmd.AddCommand(generateCmd)
}
