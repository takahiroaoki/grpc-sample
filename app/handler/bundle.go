package handler

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/pb"
)

type bundle struct {
	pb.UnimplementedSampleServiceServer
	createUserHandler  Handler[*pb.CreateUserRequest, *pb.CreateUserResponse]
	getUserInfoHandler Handler[*pb.GetUserInfoRequest, *pb.GetUserInfoResponse]
}

func (s *bundle) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.createUserHandler.execute(ctx, req)
}

func (s *bundle) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return s.getUserInfoHandler.execute(ctx, req)
}

func NewBundle(
	createUserHandler Handler[*pb.CreateUserRequest, *pb.CreateUserResponse],
	getUserInfoHandler Handler[*pb.GetUserInfoRequest, *pb.GetUserInfoResponse],
) pb.SampleServiceServer {
	return &bundle{
		createUserHandler:  createUserHandler,
		getUserInfoHandler: getUserInfoHandler,
	}
}
