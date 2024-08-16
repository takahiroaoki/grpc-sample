package service

import (
	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/repository"
	"gorm.io/gorm"
)

type CreateUserService interface {
	CreateUser(db *gorm.DB, u entity.User) (*entity.User, error)
}

type CreateUserServiceImpl struct {
	userRepository repository.UserRepository
}

func (s *CreateUserServiceImpl) CreateUser(db *gorm.DB, u entity.User) (*entity.User, error) {
	return s.userRepository.CreateOneUser(db, u)
}

func NewCreateUserService(userRepository repository.UserRepository) CreateUserService {
	return &CreateUserServiceImpl{userRepository: userRepository}
}
