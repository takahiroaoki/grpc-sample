package repository

import (
	"github.com/takahiroaoki/grpc-sample/app/backend"
	"github.com/takahiroaoki/grpc-sample/app/entity"
)

type DemoRepository interface {
	// user_repository.go
	SelectOneUserByUserId(dbw backend.DBWrapper, userId string) (*entity.User, error)
	CreateOneUser(dbw backend.DBWrapper, u entity.User) (*entity.User, error)
}

type demoRepositoryImpl struct{}

func NewDemoRepository() DemoRepository {
	return &demoRepositoryImpl{}
}
