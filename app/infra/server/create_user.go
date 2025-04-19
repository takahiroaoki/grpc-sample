package server

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/handler"
	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
)

type createUserHandler interface {
	Invoke(ctx context.Context, req *handler.CreateUserRequest) (*handler.CreateUserResponse, domerr.DomErr)
}

func (s *sampleServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := s.createUserHandler.Invoke(ctx, handler.NewCreateUserRequest(req.GetEmail()))
	if err != nil {
		return nil, handleError(ctx, err)
	}
	return &pb.CreateUserResponse{
		Id: res.Id(),
	}, nil
}
