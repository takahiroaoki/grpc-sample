package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
)

type getUserInfoHandlerImpl struct {
	dr   repository.DemoRepository
	guis service.GetUserInfoService
}

func (h *getUserInfoHandlerImpl) Invoke(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, domerr.DomErr) {
	u, err := h.guis.GetUserByUserId(ctx, h.dr, req.id)
	if err != nil {
		return nil, err
	}

	return &GetUserInfoResponse{
		id:    strconv.FormatUint(uint64(u.ID), 10),
		email: u.Email,
	}, nil
}

func NewGetUserInfoHandler(dr repository.DemoRepository, guis service.GetUserInfoService) GetUserInfoHandler {
	return &getUserInfoHandlerImpl{
		dr:   dr,
		guis: guis,
	}
}
