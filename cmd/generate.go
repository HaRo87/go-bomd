package cmd

import (
	"fmt"

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

// generateMarkdownCmd represents the generate markdown command
var generateMarkdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Generate a specified markdown report",
	Long:  `Generate (bomd generate markdown) will create the specified markdown report.`,
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
		logrus.Debugf("Trying to generate default template: %s", file)
		err := processor.Generate(file)
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
	generateCmd.AddCommand(generateMarkdownCmd)
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
