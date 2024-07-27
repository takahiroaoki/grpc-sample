package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/pb"
	"github.com/takahiroaoki/go-env/app/service"
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

func (h *SampleHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var (
		u   *entity.User
		err error
	)
	err = h.db.Transaction(func(tx *gorm.DB) error {
		u, err = h.sampleService.CreateUser(tx, entity.User{
			Email: req.GetEmail(),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id: strconv.FormatUint(uint64(u.ID), 10),
	}, nil
}

func NewSampleHandler(db *gorm.DB, sampleService service.SampleService) *SampleHandler {
	return &SampleHandler{
		db:            db,
		sampleService: sampleService,
	}
}
