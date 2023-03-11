package cmd

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/HaRo87go-bomd/replicator"
)

func generateItem(what string) {
	fmt.Println("Generating ...")
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a specified item",
	Long:  `Generate (bomd generate) will support with creating a new item.`,
}

// generateConfigCmd represents the generate config command
var generateConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate a specified config",
	Long:  `Generate (bomd generate config) will support with creating the specified config.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateItem(args[0])
	},
}

// generateResultCmd represents the generate result command
var generateResultCmd = &cobra.Command{
	Use:   "result",
	Short: "Generate a specified result",
	Long:  `Generate (bomd generate result) will create the specified result based on the provided BOM and template.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateItem(args[0])
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
		filePath := ""
		for _, file := range files {
			if strings.HasSuffix(file, ".tmpl") {
				filePath = file
			}
		}
		if len(filePath) == 0 {
			logrus.Error("😱 something went wrong")
			return fmt.Errorf("you must provide a valid file path for your template ending with .tmpl")
		}
		logrus.Debugf("Trying to generate default template: %s", filePath)
		err := processor.Generate(filePath)
		if err != nil {
			logrus.Error("😱 something went wrong")
			return err
		}
		logrus.Info("😎 everything seems to be fine")
		return nil
	},
}

func init() {
	generateCmd.AddCommand(generateConfigCmd)
	generateCmd.AddCommand(generateResultCmd)
	generateCmd.AddCommand(generateTemplateCmd)
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
