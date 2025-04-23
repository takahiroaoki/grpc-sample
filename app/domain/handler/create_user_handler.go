package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
)

type createUserHandler struct {
	cus service.CreateUserService
}

func (h *createUserHandler) Invoke(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, domerr.DomErr) {
	created, err := h.cus.CreateUser(ctx, entity.User{
		Email: req.email,
	})
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{
		id: strconv.FormatUint(uint64(created.ID), 10),
	}, nil
}

func NewCreateUserHandler(cus service.CreateUserService) CreateUserHandler {
	return &createUserHandler{
		cus: cus,
	}
}
