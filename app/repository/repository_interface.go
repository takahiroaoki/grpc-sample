package repository

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

type DemoRepository interface {
	// transaction
	Transaction(func(dr DemoRepository) error) error

	// users table
	SelectOneUserByUserId(userId string) (*entity.User, util.AppError)
	CreateOneUser(u entity.User) (*entity.User, util.AppError)
}
