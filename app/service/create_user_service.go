package service

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
)

type createUserServiceImpl struct{}

func (s *createUserServiceImpl) CreateUser(dr repository.DemoRepository, u entity.User) (*entity.User, error) {
	return dr.CreateOneUser(u)
}
