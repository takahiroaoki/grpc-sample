package server

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/handler"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
	"github.com/takahiroaoki/grpc-sample/app/domain/util"
	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
	"github.com/takahiroaoki/grpc-sample/app/infra/server/interceptor"
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
	res, err := s.createUserHandler.Invoke(ctx, handler.NewCreateUserRequest(req.GetEmail()))
	if err != nil {
		return nil, handleError(ctx, err)
	}
	return &pb.CreateUserResponse{
		Id: res.Id(),
	}, nil
}

func (s *sampleServiceServerImpl) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	res, err := s.getUserInfoHandler.Invoke(ctx, handler.NewGetUserInfoRequest(req.GetId()))
	if err != nil {
		return nil, handleError(ctx, err)
	}
	return &pb.GetUserInfoResponse{
		Id:    res.Id(),
		Email: res.Email(),
	}, nil
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
