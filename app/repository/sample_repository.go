package repository

import (
	"context"

	"github.com/takahiroaoki/go-env/entity"
	"gorm.io/gorm"
)

type SampleRepositoryInterface interface {
	SelectOneUserByUserId(ctx context.Context, userId string) (*entity.User, error)
}

func NewSampleRepository(db *gorm.DB) SampleRepositoryInterface {
	return &SampleRepositoryImpl{db: db}
}

type SampleRepositoryImpl struct {
	db *gorm.DB
}

func (r *SampleRepositoryImpl) SelectOneUserByUserId(ctx context.Context, userId string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
