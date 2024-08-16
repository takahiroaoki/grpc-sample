package handler

import (
	"context"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/takahiroaoki/go-env/app/pb"
	"github.com/takahiroaoki/go-env/app/service"
	"gorm.io/gorm"
)

type GetUserInfoHandler interface {
	getUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error)
}

type getUserInfoHandlerImpl struct {
	db                 *gorm.DB
	getUserInfoService service.GetUserInfoService
}

func (h *getUserInfoHandlerImpl) getUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	if err := h.validate(ctx, req); err != nil {
		return nil, err
	}

	u, err := h.getUserInfoService.GetUserByUserId(h.db, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserInfoResponse{
		Id:    strconv.FormatUint(uint64(u.ID), 10),
		Email: u.Email,
	}, nil
}

func (h *getUserInfoHandlerImpl) validate(ctx context.Context, req *pb.GetUserInfoRequest) error {
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.Id, validation.Required, is.Digit))

	return validation.ValidateStructWithContext(ctx, req, rules...)
}

func NewGetUserInfoHandler(db *gorm.DB, getUserInfoService service.GetUserInfoService) GetUserInfoHandler {
	return &getUserInfoHandlerImpl{
		db:                 db,
		getUserInfoService: getUserInfoService,
	}
}
