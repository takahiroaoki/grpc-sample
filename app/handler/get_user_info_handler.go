package handler

import (
	"context"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/service"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

type getUserInfoHandlerImpl struct {
	dr   repository.DemoRepository
	guis service.GetUserInfoService
}

func (h *getUserInfoHandlerImpl) process(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, util.AppError) {
	if h == nil {
		return nil, util.NewAppErrorFromMsg("*getUserInfoHandlerImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR)
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

func (h *getUserInfoHandlerImpl) validate(ctx context.Context, req *GetUserInfoRequest) util.AppError {
	if h == nil {
		return util.NewAppErrorFromMsg("*getUserInfoHandlerImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR)
	}
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.id, validation.Required, is.Digit))

	return util.NewAppError(
		validation.ValidateStructWithContext(ctx, req, rules...),
		util.CAUSE_INVALID_ARGUMENT,
		util.LOG_LEVEL_INFO,
	)
}

func NewGetUserInfoHandler(dr repository.DemoRepository, guis service.GetUserInfoService) GetUserInfoHandler {
	return &getUserInfoHandlerImpl{
		dr:   dr,
		guis: guis,
	}
}
