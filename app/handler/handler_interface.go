package handler

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/pb"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/service"
)

type Handler[Req, Res any] interface {
	execute(ctx context.Context, req *Req) (*Res, error)
	validate(ctx context.Context, req *Req) error
}

type CreateUserHandler interface {
	Handler[pb.CreateUserRequest, pb.CreateUserResponse]
}

func NewCreateUserHandler(dr repository.DemoRepository, cus service.CreateUserService) CreateUserHandler {
	return &createUserHandlerImpl{
		dr:  dr,
		cus: cus,
	}
}

type GetUserInfoHandler interface {
	Handler[pb.GetUserInfoRequest, pb.GetUserInfoResponse]
}

func NewGetUserInfoHandler(dr repository.DemoRepository, guis service.GetUserInfoService) GetUserInfoHandler {
	return &getUserInfoHandlerImpl{
		dr:   dr,
		guis: guis,
	}
}
