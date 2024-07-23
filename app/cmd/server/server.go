package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/graceful"
	"github.com/takahiroaoki/go-env/handler"
	"github.com/takahiroaoki/go-env/repository"
	"github.com/takahiroaoki/go-env/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewCmdServer() *cobra.Command {
	var profile string

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Boot command of web server with a profile.",
		Long:  "Boot command of web server with a profile. Usage example: 'go run main.go server -p prod'",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(profile)

			dsn := "dev-user:password@tcp(demo-mysql:3306)/demodb?charset=utf8"
			db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			sampleRepository := repository.NewSampleRepository(db)
			sampleService := service.NewSampleService(sampleRepository)
			ctx := context.Background()

			server := http.NewServeMux()
			server.Handle("/", handler.NewSampleHandler(ctx, sampleService))

			fmt.Println("Starting web server...")
			graceful.Run(":8080", 1*time.Second, server)
			fmt.Println("Stopping web server...")

			return nil
		},
	}

	serverCmd.Flags().StringVarP(&profile, "profile", "p", "local", "Running profile: 'local', 'prod'")
	return serverCmd
}
