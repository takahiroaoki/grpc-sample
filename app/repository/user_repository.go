package repository

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/infra"
)

func (r *demoRepositoryImpl) SelectOneUserByUserId(dbw infra.DBWrapper, userId string) (*entity.User, error) {
	var user entity.User
	if err := dbw.DB().Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *demoRepositoryImpl) CreateOneUser(dbw infra.DBWrapper, u entity.User) (*entity.User, error) {
	if err := dbw.DB().Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
