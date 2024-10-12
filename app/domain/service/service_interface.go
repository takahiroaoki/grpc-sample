package service

import (
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type CreateUserService interface {
	CreateUser(dr repository.DemoRepository, u entity.User) (*entity.User, domerr.DomErr)
}

type GetUserInfoService interface {
	GetUserByUserId(dr repository.DemoRepository, userId string) (*entity.User, domerr.DomErr)
}
