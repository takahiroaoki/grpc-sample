package server

import (
	"github.com/takahiroaoki/grpc-sample/app/domain/handler"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
	"github.com/takahiroaoki/grpc-sample/app/infra/server/interceptor"
	"github.com/takahiroaoki/grpc-sample/app/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

type sampleServiceServerImpl struct {
	pb.UnimplementedSampleServiceServer
	createUserHandler  handler.CreateUserHandler
	getUserInfoHandler handler.GetUserInfoHandler
}

func newSampleServiceServer(dr repository.DemoRepository) pb.SampleServiceServer {
	getUserInfoService := service.NewGetUserInfoService(dr)
	createUserService := service.NewCreateUserService(dr)

	return &sampleServiceServerImpl{
		createUserHandler:  handler.NewCreateUserHandler(createUserService),
		getUserInfoHandler: handler.NewGetUserInfoHandler(getUserInfoService),
	}
}

func NewGRPCServer(dr repository.DemoRepository, refFlg bool) *grpc.Server {
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
	pb.RegisterSampleServiceServer(server, newSampleServiceServer(dr))
	return server
}
