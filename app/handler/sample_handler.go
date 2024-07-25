package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/takahiroaoki/go-env/pb"
	"github.com/takahiroaoki/go-env/service"
	"github.com/takahiroaoki/go-env/util"
)

type SampleHandler struct {
	pb.UnimplementedSampleServiceServer
	sampleService service.SampleService
}

func (h *SampleHandler) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	u, err := h.sampleService.GetUserByUserId(req.GetId())
	if err != nil {
		util.ErrorLog(fmt.Sprintf("Failed to get user info. ID=%v", req.GetId()))
		return nil, err
	}

	return &pb.GetUserInfoResponse{
		Id:    strconv.FormatUint(uint64(u.ID), 10),
		Email: u.Email,
	}, nil
}

func NewSampleHandler(sampleService service.SampleService) *SampleHandler {
	return &SampleHandler{
		sampleService: sampleService,
	}
}
