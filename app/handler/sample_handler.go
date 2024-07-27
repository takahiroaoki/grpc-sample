package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/go-env/pb"
	"github.com/takahiroaoki/go-env/service"
	"gorm.io/gorm"
)

type SampleHandler struct {
	pb.UnimplementedSampleServiceServer
	db            *gorm.DB
	sampleService service.SampleService
}

func (h *SampleHandler) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	u, err := h.sampleService.GetUserByUserId(h.db, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserInfoResponse{
		Id:    strconv.FormatUint(uint64(u.ID), 10),
		Email: u.Email,
	}, nil
}

func NewSampleHandler(db *gorm.DB, sampleService service.SampleService) *SampleHandler {
	return &SampleHandler{
		db:            db,
		sampleService: sampleService,
	}
}
