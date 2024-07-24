package handler

import (
	"context"
	"fmt"

	"github.com/takahiroaoki/go-env/pb"
	"github.com/takahiroaoki/go-env/service"
)

type SampleHandler struct {
	pb.UnimplementedSampleServiceServer
	sampleService service.SampleService
}

func (h *SampleHandler) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	u, err := h.sampleService.GetUserByUserId(req.GetId())
	if err != nil {
		fmt.Println("Failed to get user info")
		return nil, err
	}

	return &pb.GetUserInfoResponse{
		Id:    req.GetId(),
		Email: u.Email,
	}, nil
}

func NewSampleHandler(sampleService service.SampleService) *SampleHandler {
	return &SampleHandler{
		sampleService: sampleService,
	}
}
