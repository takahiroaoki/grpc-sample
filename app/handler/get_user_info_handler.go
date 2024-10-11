package handler

import (
	"context"
	"errors"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/service"
)

type getUserInfoHandlerImpl struct {
	dr   repository.DemoRepository
	guis service.GetUserInfoService
}

func (h *getUserInfoHandlerImpl) process(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	if h == nil {
		return nil, errors.New("*getUserInfoHandlerImpl is nil")
	}

	u, err := h.guis.GetUserByUserId(h.dr, req.id)
	if err != nil {
		return nil, err
	}

	return &GetUserInfoResponse{
		id:    strconv.FormatUint(uint64(u.ID), 10),
		email: u.Email,
	}, nil
}

func (h *getUserInfoHandlerImpl) validate(ctx context.Context, req *GetUserInfoRequest) error {
	if h == nil {
		return errors.New("*getUserInfoHandlerImpl is nil")
	}
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.id, validation.Required, is.Digit))

	return validation.ValidateStructWithContext(ctx, req, rules...)
}

func NewGetUserInfoHandler(dr repository.DemoRepository, guis service.GetUserInfoService) GetUserInfoHandler {
	return &getUserInfoHandlerImpl{
		dr:   dr,
		guis: guis,
	}
}
