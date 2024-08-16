package handler

import (
	"context"

	"github.com/takahiroaoki/go-env/app/pb"
)

type BundleServer struct {
	pb.UnimplementedSampleServiceServer
	createUserHandler  CreateUserHandler
	getUserInfoHandler GetUserInfoHandler
}

func (s *BundleServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.createUserHandler.createUser(ctx, req)
}

func (s *BundleServer) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return s.getUserInfoHandler.getUserInfo(ctx, req)
}

func NewBundleServer(createUserHandler CreateUserHandler, getUserInfoHandler GetUserInfoHandler) *BundleServer {
	return &BundleServer{
		createUserHandler:  createUserHandler,
		getUserInfoHandler: getUserInfoHandler,
	}
}
