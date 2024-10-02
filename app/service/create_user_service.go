package service

import (
	"github.com/takahiroaoki/grpc-sample/app/backend"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
)

type createUserServiceImpl struct {
	demoRepository repository.DemoRepository
}

func (s *createUserServiceImpl) CreateUser(dbw backend.DBWrapper, u entity.User) (*entity.User, error) {
	return s.demoRepository.CreateOneUser(dbw, u)
}
