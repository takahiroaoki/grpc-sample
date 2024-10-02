package handler

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/backend"
	"github.com/takahiroaoki/grpc-sample/app/pb"
	"github.com/takahiroaoki/grpc-sample/app/service"
)

type Handler[Req, Res any] interface {
	execute(ctx context.Context, req *Req) (*Res, error)
	validate(ctx context.Context, req *Req) error
}

type CreateUserHandler interface {
	Handler[pb.CreateUserRequest, pb.CreateUserResponse]
}

func NewCreateUserHandler(dbw backend.DBWrapper, createUserService service.CreateUserService) CreateUserHandler {
	return &createUserHandlerImpl{
		dbw:               dbw,
		createUserService: createUserService,
	}
}

type GetUserInfoHandler interface {
	Handler[pb.GetUserInfoRequest, pb.GetUserInfoResponse]
}

func NewGetUserInfoHandler(dbw backend.DBWrapper, getUserInfoService service.GetUserInfoService) GetUserInfoHandler {
	return &getUserInfoHandlerImpl{
		dbw:                dbw,
		getUserInfoService: getUserInfoService,
	}
}
