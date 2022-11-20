package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func validateItem(config string) {
	fmt.Println("Validating ...")
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
	Run: func(cmd *cobra.Command, args []string) {
		validateItem(args[0])
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
	validateCmd.AddCommand(validateBomCmd)
	validateCmd.AddCommand(validateConfigCmd)
	validateCmd.AddCommand(validateTemplateCmd)
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
