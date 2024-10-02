package cmd

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/cobra"
	"github.com/takahiroaoki/grpc-sample/app/backend"
	"github.com/takahiroaoki/grpc-sample/app/config"
	"github.com/takahiroaoki/grpc-sample/app/handler"
	"github.com/takahiroaoki/grpc-sample/app/interceptor"
	"github.com/takahiroaoki/grpc-sample/app/pb"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/service"
	"github.com/takahiroaoki/grpc-sample/app/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func newServerCmd() *cobra.Command {
	var refFlg bool

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Boot command of web server.",
		Long:  "Boot command of web server.",
		RunE: func(cmd *cobra.Command, args []string) error {

			// Prepare db client
			dbc, err := backend.NewDBClientFromDSN(config.GetDataSourceName())
			if err != nil {
				util.FatalLog(fmt.Sprintf("Failed to get DB connection. Error: %v", err))
			}
			defer func() {
				err := dbc.CloseDB()
				if err != nil {
					util.FatalLog("Failed to close db connection")
					return
				}
				util.InfoLog("DB connection closed successfully")
			}()

			// Prepare grpc server settings
			lis, err := net.Listen("tcp", ":8080")
			if err != nil {
				util.FatalLog(fmt.Sprintf("Failed to listen: %v", err))
			}
			server := grpc.NewServer(grpc.UnaryInterceptor(
				middleware.ChainUnaryServer(
					interceptor.SetContext(),
					interceptor.Log(),
				),
			))
			if refFlg {
				reflection.Register(server)
				util.InfoLog("Server reflection is ON")
			}

			// Register gRPC handler
			pb.RegisterSampleServiceServer(server, getHandler(dbc))

			// Run
			go func() {
				if err := server.Serve(lis); err != nil {
					util.FatalLog("Failed to start gRPC server")
				}
			}()
			util.InfoLog("gRPC server started successfully")

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

	serverCmd.Flags().BoolVarP(&refFlg, "reflection", "r", false, "Reflection flag")
	return serverCmd
}

func getHandler(dr repository.DemoRepository) pb.SampleServiceServer {
	getUserInfoService := service.NewGetUserInfoService()
	createUserService := service.NewCreateUserService()

	return handler.NewBundle(
		handler.NewCreateUserHandler(dr, createUserService),
		handler.NewGetUserInfoHandler(dr, getUserInfoService),
	)
}
