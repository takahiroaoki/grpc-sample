package repository

import (
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
)

type DemoRepository interface {
	// transaction
	Transaction(func(dr DemoRepository) error) error

	// users table
	SelectOneUserByUserId(userId string) (*entity.User, error)
	CreateOneUser(u entity.User) (*entity.User, error)
}
