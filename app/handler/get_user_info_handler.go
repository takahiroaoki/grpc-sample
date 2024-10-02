package handler

import (
	"context"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/takahiroaoki/grpc-sample/app/backend"
	"github.com/takahiroaoki/grpc-sample/app/pb"
	"github.com/takahiroaoki/grpc-sample/app/service"
)

type getUserInfoHandlerImpl struct {
	dbw                backend.DBWrapper
	getUserInfoService service.GetUserInfoService
}

func (h *getUserInfoHandlerImpl) execute(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	if err := h.validate(ctx, req); err != nil {
		return nil, err
	}

	u, err := h.getUserInfoService.GetUserByUserId(h.dbw, req.GetId())
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
