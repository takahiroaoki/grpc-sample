package server

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/takahiroaoki/go-env/cmd/config"
	"github.com/takahiroaoki/go-env/handler"
	"github.com/takahiroaoki/go-env/pb"
	"github.com/takahiroaoki/go-env/repository"
	"github.com/takahiroaoki/go-env/service"
	"github.com/takahiroaoki/go-env/util"
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
				util.FatalLog(fmt.Sprintf("Failed to get DB connection. Error: %v", err))
			}
			defer closeDB(db)

			// Prepare repositories and servicies
			sampleRepository := repository.NewSampleRepository(db)
			sampleService := service.NewSampleService(sampleRepository)

			// Prepare grpc server settings
			lis, err := net.Listen("tcp", ":8080")
			if err != nil {
				util.FatalLog(fmt.Sprintf("Failed to listen: %v", err))
			}
			server := grpc.NewServer()
			pb.RegisterSampleServiceServer(server, handler.NewSampleHandler(sampleService))

			if profile != "prod" {
				reflection.Register(server)
			}

			util.InfoLog(fmt.Sprintf("Started gRPC server on profile: %v", profile))
			if err := server.Serve(lis); err != nil {
				util.FatalLog(fmt.Sprintf("Failed to serve: %v", err))
			}

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
