package cmd

import (
	"github.com/spf13/cobra"
)

func executeCommand(command *cobra.Command, args ...string) (err error) {
	command.SetArgs(args)
	err = command.Execute()
	return
}
