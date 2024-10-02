package backend

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"gorm.io/gorm"
)

type DBWrapper interface {
	Transaction(func(dbw DBWrapper) error) error
	repository.DemoRepository
}

type dbWrapperImpl struct {
	db *gorm.DB
}

func (dbw *dbWrapperImpl) Transaction(fn func(txw DBWrapper) error) error {
	return dbw.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewDBWrapper(tx))
	})
}

func (dbw *dbWrapperImpl) SelectOneUserByUserId(userId string) (*entity.User, error) {
	var user entity.User
	if err := dbw.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (dbw *dbWrapperImpl) CreateOneUser(u entity.User) (*entity.User, error) {
	if err := dbw.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func NewDBWrapper(db *gorm.DB) DBWrapper {
	return &dbWrapperImpl{
		db: db,
	}
}
