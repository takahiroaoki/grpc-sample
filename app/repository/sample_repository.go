package repository

import (
	"github.com/takahiroaoki/go-env/app/entity"
	"gorm.io/gorm"
)

type SampleRepository interface {
	SelectOneUserByUserId(db *gorm.DB, userId string) (*entity.User, error)
}

type SampleRepositoryImpl struct {
}

func (r *SampleRepositoryImpl) SelectOneUserByUserId(db *gorm.DB, userId string) (*entity.User, error) {
	var user entity.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func NewSampleRepository() SampleRepository {
	return &SampleRepositoryImpl{}
}
