package handler

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/pb"
)

type bundle struct {
	pb.UnimplementedSampleServiceServer
	createUserHandler  CreateUserHandler
	getUserInfoHandler GetUserInfoHandler
}

func (s *bundle) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.createUserHandler.execute(ctx, req)
}

func (s *bundle) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return s.getUserInfoHandler.execute(ctx, req)
}

func NewBundle(
	createUserHandler CreateUserHandler,
	getUserInfoHandler GetUserInfoHandler,
) pb.SampleServiceServer {
	return &bundle{
		createUserHandler:  createUserHandler,
		getUserInfoHandler: getUserInfoHandler,
	}
}
