package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewServerCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use: "server",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello world!")
			return nil
		},
	}
	return serverCmd
}
