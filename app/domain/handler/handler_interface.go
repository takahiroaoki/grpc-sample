package handler

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
)

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
	Invoke(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, domerr.DomErr)
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
	Invoke(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, domerr.DomErr)
}
