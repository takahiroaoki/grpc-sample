package service

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type createUserService struct{}

func (s *createUserService) CreateUser(ctx context.Context, dr repository.DemoRepository, u entity.User) (*entity.User, domerr.DomErr) {
	return dr.CreateOneUser(ctx, u)
}

func NewCreateUserService() *createUserService {
	return &createUserService{}
}
