package cmd

import (
	"github.com/spf13/cobra"
)

func NewBundle() *cobra.Command {
	bundle := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	bundle.AddCommand(newServerCmd())
	return bundle
}
