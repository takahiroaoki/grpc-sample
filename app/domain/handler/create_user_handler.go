package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
)

type CreateUserRequest struct {
	email string
}

func NewCreateUserRequest(email string) *CreateUserRequest {
	return &CreateUserRequest{
		email: email,
	}
}

type CreateUserResponse struct {
	id string
}

func (cur *CreateUserResponse) Id() string {
	if cur == nil {
		return ""
	}
	return cur.id
}

func NewCreateUserResponse(id string) *CreateUserResponse {
	return &CreateUserResponse{
		id: id,
	}
}

type createUserHandler struct {
	cus createUserService
}

type createUserService interface {
	CreateUser(ctx context.Context, u entity.User) (*entity.User, domerr.DomErr)
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

func NewCreateUserHandler(cus createUserService) *createUserHandler {
	return &createUserHandler{
		cus: cus,
	}
}
