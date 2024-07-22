package cmd

import (
	"github.com/spf13/cobra"

	"github.com/takahiroaoki/go-env/cmd/server"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	rootCmd.AddCommand(server.NewServerCmd())
	return rootCmd
}
