package service

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type createUserServiceImpl struct{}

func (s *createUserServiceImpl) CreateUser(ctx context.Context, dr repository.DemoRepository, u entity.User) (*entity.User, domerr.DomErr) {
	return dr.CreateOneUser(ctx, u)
}

func NewCreateUserService() CreateUserService {
	return &createUserServiceImpl{}
}
