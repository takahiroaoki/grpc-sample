package service

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"gorm.io/gorm"
)

type GetUserInfoService interface {
	GetUserByUserId(db *gorm.DB, userId string) (*entity.User, error)
}

type getUserInfoServiceImpl struct {
	userRepository repository.UserRepository
}

func (s *getUserInfoServiceImpl) GetUserByUserId(db *gorm.DB, userId string) (*entity.User, error) {
	user, err := s.userRepository.SelectOneUserByUserId(db, userId)
	return user, err
}

func NewGetUserInfoService(userRepository repository.UserRepository) GetUserInfoService {
	return &getUserInfoServiceImpl{userRepository: userRepository}
}
