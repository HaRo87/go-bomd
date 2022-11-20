package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func validateItem(cmd *cobra.Command, args []string) {
	fmt.Println("Validating ...")
}

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a specified item",
	Long: `Validate (bomd validate) will support with checking the integrity
	of the specified item.`,
	Run: func(cmd *cobra.Command, args []string) {
		validateItem(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
