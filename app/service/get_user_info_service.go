package service

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
)

type getUserInfoServiceImpl struct{}

func (s *getUserInfoServiceImpl) GetUserByUserId(dr repository.DemoRepository, userId string) (*entity.User, error) {
	user, err := dr.SelectOneUserByUserId(userId)
	return user, err
}
