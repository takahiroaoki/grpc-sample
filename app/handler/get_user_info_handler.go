package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/go-env/app/pb"
	"github.com/takahiroaoki/go-env/app/service"
	"gorm.io/gorm"
)

type GetUserInfoHandler interface {
	getUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error)
}

type GetUserInfoHandlerImpl struct {
	db                 *gorm.DB
	getUserInfoService service.GetUserInfoService
}

func (h *GetUserInfoHandlerImpl) getUserInfo(_ context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	u, err := h.getUserInfoService.GetUserByUserId(h.db, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserInfoResponse{
		Id:    strconv.FormatUint(uint64(u.ID), 10),
		Email: u.Email,
	}, nil
}

func NewGetUserInfoHandler(db *gorm.DB, getUserInfoService service.GetUserInfoService) GetUserInfoHandler {
	return &GetUserInfoHandlerImpl{
		db:                 db,
		getUserInfoService: getUserInfoService,
	}
}
