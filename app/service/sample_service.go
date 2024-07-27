package service

import (
	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/repository"
	"gorm.io/gorm"
)

type SampleService interface {
	GetUserByUserId(db *gorm.DB, userId string) (*entity.User, error)
	CreateUser(db *gorm.DB, u entity.User) (*entity.User, error)
}

type SampleServiceImpl struct {
	sampleRepository repository.SampleRepository
}

func (s *SampleServiceImpl) GetUserByUserId(db *gorm.DB, userId string) (*entity.User, error) {
	user, err := s.sampleRepository.SelectOneUserByUserId(db, userId)
	return user, err
}

func (s *SampleServiceImpl) CreateUser(db *gorm.DB, u entity.User) (*entity.User, error) {
	return s.sampleRepository.CreateOneUser(db, u)
}

func NewSampleService(sampleRepository repository.SampleRepository) SampleService {
	return &SampleServiceImpl{sampleRepository: sampleRepository}
}
