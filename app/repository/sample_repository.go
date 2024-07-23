package repository

import (
	"github.com/takahiroaoki/go-env/entity"
	"gorm.io/gorm"
)

type SampleRepository interface {
	SelectOneUserByUserId(userId string) (*entity.User, error)
}

type SampleRepositoryImpl struct {
	db *gorm.DB
}

func (r *SampleRepositoryImpl) SelectOneUserByUserId(userId string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func NewSampleRepository(db *gorm.DB) SampleRepository {
	return &SampleRepositoryImpl{db: db}
}
