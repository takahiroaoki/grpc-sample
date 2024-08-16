package handler

import (
	"context"
	"strconv"

	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/pb"
	"github.com/takahiroaoki/go-env/app/service"
	"gorm.io/gorm"
)

type CreateUserHandler interface {
	createUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
}

type createUserHandlerImpl struct {
	db                *gorm.DB
	createUserService service.CreateUserService
}

func (h *createUserHandlerImpl) createUser(_ context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var (
		u   *entity.User
		err error
	)
	err = h.db.Transaction(func(tx *gorm.DB) error {
		u, err = h.createUserService.CreateUser(tx, entity.User{
			Email: req.GetEmail(),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id: strconv.FormatUint(uint64(u.ID), 10),
	}, nil
}

func NewCreateUserHandler(db *gorm.DB, createUserService service.CreateUserService) CreateUserHandler {
	return &createUserHandlerImpl{
		db:                db,
		createUserService: createUserService,
	}
}
