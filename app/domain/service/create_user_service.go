package service

import (
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type createUserServiceImpl struct{}

func (s *createUserServiceImpl) CreateUser(dr repository.DemoRepository, u entity.User) (*entity.User, domerr.DomErr) {
	if s == nil {
		return nil, domerr.NewDomErrFromMsg("*createUserServiceImpl is nil", domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
	}
	return dr.CreateOneUser(u)
}

func NewCreateUserService() CreateUserService {
	return &createUserServiceImpl{}
}
