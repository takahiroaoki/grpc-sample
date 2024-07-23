package service

import (
	"github.com/takahiroaoki/go-env/entity"
	"github.com/takahiroaoki/go-env/repository"
)

type SampleService interface {
	GetUserByUserId(userId string) (*entity.User, error)
}

type SampleServiceImpl struct {
	sampleRepository repository.SampleRepository
}

func (s *SampleServiceImpl) GetUserByUserId(userId string) (*entity.User, error) {
	user, err := s.sampleRepository.SelectOneUserByUserId(userId)
	return user, err
}

func NewSampleService(sampleRepository repository.SampleRepository) SampleService {
	return &SampleServiceImpl{sampleRepository: sampleRepository}
}
