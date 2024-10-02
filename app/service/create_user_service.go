package service

import (
	"errors"

	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
)

type createUserServiceImpl struct{}

func (s *createUserServiceImpl) CreateUser(dr repository.DemoRepository, u entity.User) (*entity.User, error) {
	if s == nil {
		return nil, errors.New("*createUserServiceImpl is nil")
	}
	return dr.CreateOneUser(u)
}
