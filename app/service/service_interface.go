package service

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

type CreateUserService interface {
	CreateUser(dr repository.DemoRepository, u entity.User) (*entity.User, util.AppError)
}

type GetUserInfoService interface {
	GetUserByUserId(dr repository.DemoRepository, userId string) (*entity.User, util.AppError)
}
