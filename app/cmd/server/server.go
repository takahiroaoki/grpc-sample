package server

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/graceful"
	"github.com/takahiroaoki/go-env/cmd/config"
	"github.com/takahiroaoki/go-env/handler"
	"github.com/takahiroaoki/go-env/pb"
	"github.com/takahiroaoki/go-env/repository"
	"github.com/takahiroaoki/go-env/service"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewCmdServer() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Boot command of web server.",
		Long:  "Boot command of web server.",
		RunE: func(cmd *cobra.Command, args []string) error {

			// Prepare db client
			db, err := gorm.Open(
				mysql.Open(config.NewDataBaseConfig().GetDataSourceName()),
				&gorm.Config{},
			)
			if err != nil {
				return err
			}
			defer closeDB(db)

			// Prepare repositories and servicies
			sampleRepository := repository.NewSampleRepository(db)
			sampleService := service.NewSampleService(sampleRepository)

			// Prepare grpc server settings
			server := grpc.NewServer()
			pb.RegisterSampleServiceServer(server, handler.NewSampleHandler(sampleService))

			fmt.Println("Starting grpc server...")
			graceful.Run(":8080", 1*time.Second, server)
			fmt.Println("Stopping grpc server...")

			return nil
		},
	}
	return serverCmd
}

func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to close db")
		return
	}
	if err := sqlDB.Close(); err != nil {
		fmt.Println("Failed to close db")
	}
}
