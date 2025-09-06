package server

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/handler"
	pb "github.com/takahiroaoki/protobuf/gen_go/proto/grpc_sample/v1"
)

func (s *sampleServiceServer) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	res, err := s.getUserInfoHandler.Invoke(ctx, handler.NewGetUserInfoRequest(req.GetId()))
	if err != nil {
		return nil, handleError(ctx, err)
	}
	return &pb.GetUserInfoResponse{
		Id:    res.Id(),
		Email: res.Email(),
	}, nil
}
