package service

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/infra"
	"github.com/takahiroaoki/grpc-sample/app/repository"
)

type CreateUserService interface {
	CreateUser(dbw infra.DBWrapper, u entity.User) (*entity.User, error)
}

func NewCreateUserService(demoRepository repository.DemoRepository) CreateUserService {
	return &createUserServiceImpl{demoRepository: demoRepository}
}

type GetUserInfoService interface {
	GetUserByUserId(dbw infra.DBWrapper, userId string) (*entity.User, error)
}

func NewGetUserInfoService(demoRepository repository.DemoRepository) GetUserInfoService {
	return &getUserInfoServiceImpl{demoRepository: demoRepository}
}
