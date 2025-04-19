package server

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/handler"
	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
)

type getUserInfoHandler interface {
	Invoke(ctx context.Context, req *handler.GetUserInfoRequest) (*handler.GetUserInfoResponse, domerr.DomErr)
}

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
