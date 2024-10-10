package service

import (
	"errors"

	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
)

type getUserInfoServiceImpl struct{}

func (s *getUserInfoServiceImpl) GetUserByUserId(dr repository.DemoRepository, userId string) (*entity.User, error) {
	if s == nil {
		return nil, errors.New("*getUserInfoServiceImpl is nil")
	}
	return dr.SelectOneUserByUserId(userId)
}

func NewGetUserInfoService() GetUserInfoService {
	return &getUserInfoServiceImpl{}
}
