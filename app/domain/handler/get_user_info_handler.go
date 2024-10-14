package handler

import (
	"context"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
)

type getUserInfoHandlerImpl struct {
	dr   repository.DemoRepository
	guis service.GetUserInfoService
}

func (h *getUserInfoHandlerImpl) process(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, domerr.DomErr) {
	u, err := h.guis.GetUserByUserId(h.dr, req.id)
	if err != nil {
		return nil, err
	}

	return &GetUserInfoResponse{
		id:    strconv.FormatUint(uint64(u.ID), 10),
		email: u.Email,
	}, nil
}

func (h *getUserInfoHandlerImpl) validate(ctx context.Context, req *GetUserInfoRequest) domerr.DomErr {
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.id, validation.Required, is.Digit))

	return domerr.NewDomErr(
		validation.ValidateStructWithContext(ctx, req, rules...),
		domerr.CAUSE_INVALID_ARGUMENT,
		domerr.LOG_LEVEL_INFO,
	)
}

func NewGetUserInfoHandler(dr repository.DemoRepository, guis service.GetUserInfoService) GetUserInfoHandler {
	return &getUserInfoHandlerImpl{
		dr:   dr,
		guis: guis,
	}
}
