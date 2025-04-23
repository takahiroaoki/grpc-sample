package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
)

type getUserInfoHandler struct {
	guis service.GetUserInfoService
}

func (h *getUserInfoHandler) Invoke(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, domerr.DomErr) {
	u, err := h.guis.GetUserByUserId(ctx, req.id)
	if err != nil {
		return nil, err
	}

	return &GetUserInfoResponse{
		id:    strconv.FormatUint(uint64(u.ID), 10),
		email: u.Email,
	}, nil
}

func NewGetUserInfoHandler(guis service.GetUserInfoService) GetUserInfoHandler {
	return &getUserInfoHandler{
		guis: guis,
	}
}
