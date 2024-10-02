package backend

import (
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClient interface {
	repository.DemoRepository
	CloseDB() error
}

type dbClientImpl struct {
	db *gorm.DB
}

func (dbc *dbClientImpl) Transaction(fn func(dr repository.DemoRepository) error) error {
	return dbc.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewDBClient(tx))
	})
}

func (dbc *dbClientImpl) SelectOneUserByUserId(userId string) (*entity.User, error) {
	var user entity.User
	if err := dbc.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (dbc *dbClientImpl) CreateOneUser(u entity.User) (*entity.User, error) {
	if err := dbc.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (dbc *dbClientImpl) CloseDB() error {
	sqlDB, err := dbc.db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}
	return nil
}

func NewDBClientFromDSN(dataSourceName string) (DBClient, error) {
	db, err := gorm.Open(
		mysql.Open(dataSourceName),
		&gorm.Config{
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		return nil, err
	}
	return NewDBClient(db), nil
}

func NewDBClient(db *gorm.DB) DBClient {
	return &dbClientImpl{
		db: db,
	}
}
