package service

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"gorm.io/gorm"
)

type CreateUserService interface {
	CreateUser(db *gorm.DB, u entity.User) (*entity.User, error)
}

type createUserServiceImpl struct {
	userRepository repository.UserRepository
}

func (s *createUserServiceImpl) CreateUser(db *gorm.DB, u entity.User) (*entity.User, error) {
	return s.userRepository.CreateOneUser(db, u)
}

func NewCreateUserService(userRepository repository.UserRepository) CreateUserService {
	return &createUserServiceImpl{userRepository: userRepository}
}
