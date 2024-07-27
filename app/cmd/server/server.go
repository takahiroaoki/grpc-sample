package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/cobra"
	"github.com/takahiroaoki/go-env/config"
	"github.com/takahiroaoki/go-env/handler"
	"github.com/takahiroaoki/go-env/interceptor"
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

			// Prepare grpc server settings
			lis, err := net.Listen("tcp", ":8080")
			if err != nil {
				util.FatalLog(fmt.Sprintf("Failed to listen: %v", err))
			}
			server := grpc.NewServer(grpc.UnaryInterceptor(
				middleware.ChainUnaryServer(
					interceptor.Log(),
					interceptor.ValidateReq(),
				),
			))
			if profile != "prod" {
				reflection.Register(server)
			}

			// Prepare repositories and servicies for dependency injection
			sampleRepository := repository.NewSampleRepository(db)
			sampleService := service.NewSampleService(sampleRepository)

			// Register gRPC handler
			pb.RegisterSampleServiceServer(server, handler.NewSampleHandler(sampleService))

			// Run
			go func() {
				if err := server.Serve(lis); err != nil {
					util.FatalLog("Failed to start gRPC server")
				}
			}()
			util.InfoLog(fmt.Sprintf("gRPC server started successfully on profile: %v", profile))

			// Shutdown settings
			quitCh := make(chan os.Signal, 1)
			signal.Notify(quitCh, syscall.SIGTERM, syscall.SIGINT)
			<-quitCh
			// Stop accepting new request. This must be called before closeDB() method.
			server.GracefulStop()
			util.InfoLog("gRPC server stopped successfully")
			return nil
		},
	}

	serverCmd.Flags().StringVarP(&profile, "profile", "p", "local", "Running profile: 'local', 'prod'")
	return serverCmd
}

// closeDB Close DB connection. This method must be called after gracefully stop of server.
func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		util.FatalLog("Failed to close db connection")
		return
	}
	if err := sqlDB.Close(); err != nil {
		util.FatalLog("Failed to close db connection")
		return
	}
	util.InfoLog("DB connection closed successfully")
}
