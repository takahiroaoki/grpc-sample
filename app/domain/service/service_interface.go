package service

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
)

type CreateUserService interface {
	CreateUser(ctx context.Context, u entity.User) (entity.User, domerr.DomErr)
}

type GetUserInfoService interface {
	GetUserByUserId(ctx context.Context, userId string) (entity.User, domerr.DomErr)
}
