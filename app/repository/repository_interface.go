package repository

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/infra"
)

type DemoRepository interface {
	// user_repository.go
	SelectOneUserByUserId(dbw infra.DBWrapper, userId string) (*entity.User, error)
	CreateOneUser(dbw infra.DBWrapper, u entity.User) (*entity.User, error)
}

type demoRepositoryImpl struct{}

func NewDemoRepository() DemoRepository {
	return &demoRepositoryImpl{}
}
