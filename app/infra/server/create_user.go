package server

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/handler"
	pb "github.com/takahiroaoki/protobuf/gen_go/proto/grpc_sample/v1"
)

func (s *sampleServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := s.createUserHandler.Invoke(ctx, handler.NewCreateUserRequest(req.GetEmail()))
	if err != nil {
		return nil, handleError(ctx, err)
	}
	return &pb.CreateUserResponse{
		Id: res.Id(),
	}, nil
}
