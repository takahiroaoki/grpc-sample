package service

import (

	"github.com/takahiroaoki/go-env/entity"
	"github.com/takahiroaoki/go-env/repository"
)

type SampleServiceInterface interface {
	GetUserByUserId(userId string) (*entity.User, error)
}

func NewSampleService(sampleRepository repository.SampleRepositoryInterface) SampleServiceInterface {
	return &SampleServiceImple{sampleRepository: sampleRepository}
}

type SampleServiceImple struct {
	sampleRepository repository.SampleRepositoryInterface
}

func (s *SampleServiceImple) GetUserByUserId(userId string) (*entity.User, error) {
	user, err := s.sampleRepository.SelectOneUserByUserId(userId)
	return user, err
}
