package repository

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
)

type DemoRepository interface {
	SelectOneUserByUserId(userId string) (*entity.User, error)
	CreateOneUser(u entity.User) (*entity.User, error)
}
