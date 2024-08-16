package repository

import (
	"github.com/takahiroaoki/go-env/app/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	SelectOneUserByUserId(db *gorm.DB, userId string) (*entity.User, error)
	CreateOneUser(db *gorm.DB, u entity.User) (*entity.User, error)
}

type userRepositoryImpl struct {
}

func (r *userRepositoryImpl) SelectOneUserByUserId(db *gorm.DB, userId string) (*entity.User, error) {
	var user entity.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) CreateOneUser(db *gorm.DB, u entity.User) (*entity.User, error) {
	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}
