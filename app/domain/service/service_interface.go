package service

import (
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type CreateUserService interface {
	CreateUser(dr repository.DemoRepository, u entity.User) (*entity.User, error)
}

func NewCreateUserService() CreateUserService {
	return &createUserServiceImpl{}
}

type GetUserInfoService interface {
	GetUserByUserId(dr repository.DemoRepository, userId string) (*entity.User, error)
}

func NewGetUserInfoService() GetUserInfoService {
	return &getUserInfoServiceImpl{}
}
