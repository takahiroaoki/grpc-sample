package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
)

type createUserHandlerImpl struct {
	dr  repository.DemoRepository
	cus service.CreateUserService
}

func (h *createUserHandlerImpl) Invoke(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, domerr.DomErr) {
	var (
		u   *entity.User
		err error
	)
	err = h.dr.Transaction(func(dr repository.DemoRepository) error {
		u, err = h.cus.CreateUser(dr, entity.User{
			Email: req.email,
		})
		return err
	})
	if err != nil {
		appErr, ok := err.(domerr.DomErr)
		if !ok {
			return nil, domerr.NewDomErr(err, domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
		}
		return nil, appErr
	}
	return &CreateUserResponse{
		id: strconv.FormatUint(uint64(u.ID), 10),
	}, nil
}

func NewCreateUserHandler(dr repository.DemoRepository, cus service.CreateUserService) CreateUserHandler {
	return &createUserHandlerImpl{
		dr:  dr,
		cus: cus,
	}
}
