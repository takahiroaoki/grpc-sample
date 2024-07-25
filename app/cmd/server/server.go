package server

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/takahiroaoki/go-env/cmd/config"
	"github.com/takahiroaoki/go-env/handler"
	"github.com/takahiroaoki/go-env/pb"
	"github.com/takahiroaoki/go-env/repository"
	"github.com/takahiroaoki/go-env/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewCmdServer() *cobra.Command {
	var profile string

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
			lis, err := net.Listen("tcp", ":8080")
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			server := grpc.NewServer()
			pb.RegisterSampleServiceServer(server, handler.NewSampleHandler(sampleService))

			if profile != "prod" {
				reflection.Register(server)
			}

			fmt.Println("Starting grpc server...")
			if err := server.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
			fmt.Println("Stopping grpc server...")

			return nil
		},
	}

	serverCmd.Flags().StringVarP(&profile, "profile", "p", "local", "Running profile: 'local', 'prod'")
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
