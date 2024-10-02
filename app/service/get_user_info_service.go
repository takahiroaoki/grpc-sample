package service

import (
	"github.com/takahiroaoki/grpc-sample/app/backend"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
)

type getUserInfoServiceImpl struct {
	demoRepository repository.DemoRepository
}

func (s *getUserInfoServiceImpl) GetUserByUserId(dbw backend.DBWrapper, userId string) (*entity.User, error) {
	user, err := s.demoRepository.SelectOneUserByUserId(dbw, userId)
	return user, err
}
