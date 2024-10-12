package repository

import (
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
)

type DemoRepository interface {
	// transaction
	Transaction(func(dr DemoRepository) error) error

	// users table
	SelectOneUserByUserId(userId string) (*entity.User, domerr.DomErr)
	CreateOneUser(u entity.User) (*entity.User, domerr.DomErr)
}
