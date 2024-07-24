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
	fmt.Println("Receiving request")

	return &pb.GetUserInfoResponse{
		Id:    "1",
		Email: "user@test.example",
	}, nil
}

func NewSampleHandler(sampleService service.SampleService) *SampleHandler {
	return &SampleHandler{
		sampleService: sampleService,
	}
}
