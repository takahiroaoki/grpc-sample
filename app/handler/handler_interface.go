package handler

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/service"
)

type handler[Req, Res any] interface {
	Execute(ctx context.Context, req *Req) (*Res, error)
	validate(ctx context.Context, req *Req) error
}

/*
 * CreateUser
 */

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
	return cur.id
}

type CreateUserHandler interface {
	handler[CreateUserRequest, CreateUserResponse]
}

func NewCreateUserHandler(dr repository.DemoRepository, cus service.CreateUserService) CreateUserHandler {
	return &createUserHandlerImpl{
		dr:  dr,
		cus: cus,
	}
}

/*
 * GetUserInfo
 */

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
	return guihr.id
}

func (guihr *GetUserInfoResponse) Email() string {
	return guihr.email
}

type GetUserInfoHandler interface {
	handler[GetUserInfoRequest, GetUserInfoResponse]
}

func NewGetUserInfoHandler(dr repository.DemoRepository, guis service.GetUserInfoService) GetUserInfoHandler {
	return &getUserInfoHandlerImpl{
		dr:   dr,
		guis: guis,
	}
}
