package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	rootCmd.AddCommand(newCmdServer())
	return rootCmd
}
