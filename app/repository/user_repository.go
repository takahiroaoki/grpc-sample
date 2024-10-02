package repository

import (
	"github.com/takahiroaoki/grpc-sample/app/backend"
	"github.com/takahiroaoki/grpc-sample/app/entity"
)

func (r *demoRepositoryImpl) SelectOneUserByUserId(dbw backend.DBWrapper, userId string) (*entity.User, error) {
	var user entity.User
	if err := dbw.DB().Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *demoRepositoryImpl) CreateOneUser(dbw backend.DBWrapper, u entity.User) (*entity.User, error) {
	if err := dbw.DB().Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
