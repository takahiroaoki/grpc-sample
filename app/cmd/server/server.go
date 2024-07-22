package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/graceful"
	"github.com/takahiroaoki/go-env/handler"
)

func NewCmdServer() *cobra.Command {
	var profile string

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Boot command of web server with a profile.",
		Long:  "Boot command of web server with a profile. Usage example: 'go run main.go server -p prod'",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(profile)

			server := http.NewServeMux()
			server.HandleFunc("/", handler.HealthCheck)

			fmt.Println("Starting web server...")
			graceful.Run(":8080", 1*time.Second, server)
			fmt.Println("Stopping web server...")

			return nil
		},
	}

	serverCmd.Flags().StringVarP(&profile, "profile", "p", "local", "Running profile: 'local', 'prod'")
	return serverCmd
}
