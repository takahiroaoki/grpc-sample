package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
)

type GetUserInfoRequest struct {
	id string
}

func NewGetUserInfoRequest(id string) *GetUserInfoRequest {
	return &GetUserInfoRequest{
		id: id,
	}
}

type GetUserInfoResponse struct {
	id    string
	email string
}

func (guihr *GetUserInfoResponse) Id() string {
	if guihr == nil {
		return ""
	}
	return guihr.id
}

func (guihr *GetUserInfoResponse) Email() string {
	if guihr == nil {
		return ""
	}
	return guihr.email
}

func NewGetUserInfoResponse(id, email string) *GetUserInfoResponse {
	return &GetUserInfoResponse{
		id:    id,
		email: email,
	}
}

type getUserInfoHandler struct {
	guis getUserInfoService
}

type getUserInfoService interface {
	GetUserByUserId(ctx context.Context, userId string) (*entity.User, domerr.DomErr)
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

func NewGetUserInfoHandler(guis getUserInfoService) *getUserInfoHandler {
	return &getUserInfoHandler{
		guis: guis,
	}
}
