package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/graceful"
	"github.com/takahiroaoki/go-env/handler"
)

func NewServerCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use: "server",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello world!")

			server := http.NewServeMux()
			server.HandleFunc("/", handler.HealthCheck)

			fmt.Println("Starting web server...")
			graceful.Run(":8080", 1*time.Second, server)
			fmt.Println("Stopping web server...")

			return nil
		},
	}
	return serverCmd
}
