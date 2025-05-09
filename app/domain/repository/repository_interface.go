package repository

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
)

type DemoRepository interface {
	// transaction
	Transaction(func(dr DemoRepository) error) error

	// users table
	SelectOneUserByUserId(ctx context.Context, userId string) (entity.User, domerr.DomErr)
	CreateOneUser(ctx context.Context, u entity.User) (entity.User, domerr.DomErr)
}
