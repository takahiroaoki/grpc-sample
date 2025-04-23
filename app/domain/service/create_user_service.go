package service

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
)

type createUserService struct {
	dr repository.DemoRepository
}

func (s *createUserService) CreateUser(ctx context.Context, u entity.User) (*entity.User, domerr.DomErr) {
	var (
		createdUser *entity.User
		err         error
	)
	err = s.dr.Transaction(func(dr repository.DemoRepository) error {
		createdUser, err = s.dr.CreateOneUser(ctx, u)
		return err
	})
	if err != nil {
		appErr, ok := err.(domerr.DomErr)
		if !ok {
			return nil, domerr.NewDomErr(err, domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
		}
		return nil, appErr
	}
	return createdUser, nil
}

func NewCreateUserService(dr repository.DemoRepository) CreateUserService {
	return &createUserService{dr}
}
