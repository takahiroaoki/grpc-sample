package server

import (
	"github.com/takahiroaoki/grpc-sample/app/domain/handler"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
	"github.com/takahiroaoki/grpc-sample/app/infra/database"
	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
	"github.com/takahiroaoki/grpc-sample/app/infra/server/interceptor"
	"github.com/takahiroaoki/grpc-sample/app/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

type sampleServiceServer struct {
	pb.UnimplementedSampleServiceServer
	createUserHandler  createUserHandler
	getUserInfoHandler getUserInfoHandler
}

func newSampleServiceServer(dbc *database.DBClient) pb.SampleServiceServer {
	getUserInfoService := service.NewGetUserInfoService(dbc)
	createUserService := service.NewCreateUserService(dbc)

	return &sampleServiceServer{
		createUserHandler:  handler.NewCreateUserHandler(createUserService),
		getUserInfoHandler: handler.NewGetUserInfoHandler(getUserInfoService),
	}
}

func NewGRPCServer(dbc *database.DBClient, refFlg bool) *grpc.Server {
	server := grpc.NewServer(grpc.UnaryInterceptor(
		middleware.ChainUnaryServer(
			interceptor.SetContext(),
			interceptor.PerformanceLog(),
			interceptor.Validate(),
		),
	))
	if refFlg {
		reflection.Register(server)
		util.InfoLog("Server reflection is ON")
	}

	// Register gRPC handler
	pb.RegisterSampleServiceServer(server, newSampleServiceServer(dbc))
	return server
}
