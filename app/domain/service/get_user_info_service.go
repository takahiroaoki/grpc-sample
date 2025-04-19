package service

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
)

type getUserInfoService struct {
	guir getUserInfoRepository
}

type getUserInfoRepository interface {
	SelectOneUserByUserId(ctx context.Context, userId string) (*entity.User, domerr.DomErr)
}

func (s *getUserInfoService) GetUserByUserId(ctx context.Context, userId string) (*entity.User, domerr.DomErr) {
	return s.guir.SelectOneUserByUserId(ctx, userId)
}

func NewGetUserInfoService(guir getUserInfoRepository) *getUserInfoService {
	return &getUserInfoService{
		guir: guir,
	}
}
