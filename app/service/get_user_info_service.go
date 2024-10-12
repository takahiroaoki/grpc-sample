package service

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

type getUserInfoServiceImpl struct{}

func (s *getUserInfoServiceImpl) GetUserByUserId(dr repository.DemoRepository, userId string) (*entity.User, util.AppError) {
	if s == nil {
		return nil, util.NewAppErrorFromMsg("*getUserInfoServiceImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR)
	}
	return dr.SelectOneUserByUserId(userId)
}

func NewGetUserInfoService() GetUserInfoService {
	return &getUserInfoServiceImpl{}
}
