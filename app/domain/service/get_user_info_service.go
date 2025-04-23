package service

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type getUserInfoServiceImpl struct {
	dr repository.DemoRepository
}

func (s *getUserInfoServiceImpl) GetUserByUserId(ctx context.Context, userId string) (*entity.User, domerr.DomErr) {
	return s.dr.SelectOneUserByUserId(ctx, userId)
}

func NewGetUserInfoService(dr repository.DemoRepository) GetUserInfoService {
	return &getUserInfoServiceImpl{dr}
}
