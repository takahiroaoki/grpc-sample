package service

import (
	"github.com/takahiroaoki/grpc-sample/app/backend"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
)

type CreateUserService interface {
	CreateUser(dbw backend.DBWrapper, u entity.User) (*entity.User, error)
}

func NewCreateUserService(demoRepository repository.DemoRepository) CreateUserService {
	return &createUserServiceImpl{demoRepository: demoRepository}
}

type GetUserInfoService interface {
	GetUserByUserId(dbw backend.DBWrapper, userId string) (*entity.User, error)
}

func NewGetUserInfoService(demoRepository repository.DemoRepository) GetUserInfoService {
	return &getUserInfoServiceImpl{demoRepository: demoRepository}
}
