package service

import (
	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/repository"
	"gorm.io/gorm"
)

type GetUserInfoService interface {
	GetUserByUserId(db *gorm.DB, userId string) (*entity.User, error)
}

type GetUserInfoServiceImpl struct {
	userRepository repository.UserRepository
}

func (s *GetUserInfoServiceImpl) GetUserByUserId(db *gorm.DB, userId string) (*entity.User, error) {
	user, err := s.userRepository.SelectOneUserByUserId(db, userId)
	return user, err
}

func NewGetUserInfoService(userRepository repository.UserRepository) GetUserInfoService {
	return &GetUserInfoServiceImpl{userRepository: userRepository}
}
