package service

import (
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type getUserInfoServiceImpl struct{}

func (s *getUserInfoServiceImpl) GetUserByUserId(dr repository.DemoRepository, userId string) (*entity.User, domerr.DomErr) {
	if s == nil {
		return nil, domerr.NewDomErrFromMsg("*getUserInfoServiceImpl is nil", domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
	}
	return dr.SelectOneUserByUserId(userId)
}

func NewGetUserInfoService() GetUserInfoService {
	return &getUserInfoServiceImpl{}
}
