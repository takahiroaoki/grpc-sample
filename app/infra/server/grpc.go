package server

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
	"github.com/takahiroaoki/grpc-sample/app/handler"
	"github.com/takahiroaoki/grpc-sample/app/infra/interceptor"
	"github.com/takahiroaoki/grpc-sample/app/pb"
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

func (s *sampleServiceServerImpl) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.createUserHandler.Execute(ctx, req)
}

func (s *sampleServiceServerImpl) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return s.getUserInfoHandler.Execute(ctx, req)
}

func newSampleServiceServer(dr repository.DemoRepository) pb.SampleServiceServer {
	getUserInfoService := service.NewGetUserInfoService()
	createUserService := service.NewCreateUserService()

	return &sampleServiceServerImpl{
		createUserHandler:  handler.NewCreateUserHandler(dr, createUserService),
		getUserInfoHandler: handler.NewGetUserInfoHandler(dr, getUserInfoService),
	}
}

func NewGRPCServer(dr repository.DemoRepository, refFlg bool) *grpc.Server {
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
	pb.RegisterSampleServiceServer(server, newSampleServiceServer(dr))
	return server
}
