package service

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

type createUserServiceImpl struct{}

func (s *createUserServiceImpl) CreateUser(dr repository.DemoRepository, u entity.User) (*entity.User, util.AppError) {
	if s == nil {
		return nil, util.NewAppErrorFromMsg("*createUserServiceImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR)
	}
	return dr.CreateOneUser(u)
}

func NewCreateUserService() CreateUserService {
	return &createUserServiceImpl{}
}
