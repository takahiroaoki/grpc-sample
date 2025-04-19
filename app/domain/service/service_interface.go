package service

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type CreateUserService interface {
	CreateUser(ctx context.Context, dr repository.DemoRepository, u entity.User) (*entity.User, domerr.DomErr)
}

type GetUserInfoService interface {
	GetUserByUserId(ctx context.Context, dr repository.DemoRepository, userId string) (*entity.User, domerr.DomErr)
}
