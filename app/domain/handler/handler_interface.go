package handler

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
)

type handler[Req, Res any] interface {
	validate(ctx context.Context, req *Req) domerr.DomErr
	process(ctx context.Context, req *Req) (*Res, domerr.DomErr)
}

func Execute[Req, Res any](ctx context.Context, req *Req, handler handler[Req, Res]) (*Res, domerr.DomErr) {
	if err := handler.validate(ctx, req); err != nil {
		return nil, err
	}
	return handler.process(ctx, req)
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
	if cur == nil {
		return ""
	}
	return cur.id
}

type CreateUserHandler interface {
	handler[CreateUserRequest, CreateUserResponse]
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

type GetUserInfoHandler interface {
	handler[GetUserInfoRequest, GetUserInfoResponse]
}
