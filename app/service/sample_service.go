package service

import (
	"context"

	"github.com/takahiroaoki/go-env/entity"
	"github.com/takahiroaoki/go-env/repository"
)

type SampleServiceInterface interface {
	GetUserByUserId(ctx context.Context, userId string) (*entity.User, error)
}

func NewSampleService(sampleRepository repository.SampleRepositoryInterface) SampleServiceInterface {
	return &SampleServiceImple{sampleRepository: sampleRepository}
}

type SampleServiceImple struct {
	sampleRepository repository.SampleRepositoryInterface
}

func (s *SampleServiceImple) GetUserByUserId(ctx context.Context, userId string) (*entity.User, error) {
	user, err := s.sampleRepository.SelectOneUserByUserId(ctx, userId)
	return user, err
}
