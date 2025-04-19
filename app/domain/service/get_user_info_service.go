package service

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type getUserInfoServiceImpl struct{}

func (s *getUserInfoServiceImpl) GetUserByUserId(ctx context.Context, dr repository.DemoRepository, userId string) (*entity.User, domerr.DomErr) {
	return dr.SelectOneUserByUserId(ctx, userId)
}

func NewGetUserInfoService() GetUserInfoService {
	return &getUserInfoServiceImpl{}
}
