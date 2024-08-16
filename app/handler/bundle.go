package handler

import (
	"context"

	"github.com/takahiroaoki/go-env/app/pb"
)

type Bundle struct {
	pb.UnimplementedSampleServiceServer
	createUserHandler  CreateUserHandler
	getUserInfoHandler GetUserInfoHandler
}

func (s *Bundle) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.createUserHandler.createUser(ctx, req)
}

func (s *Bundle) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return s.getUserInfoHandler.getUserInfo(ctx, req)
}

func NewBundle(createUserHandler CreateUserHandler, getUserInfoHandler GetUserInfoHandler) *Bundle {
	return &Bundle{
		createUserHandler:  createUserHandler,
		getUserInfoHandler: getUserInfoHandler,
	}
}
