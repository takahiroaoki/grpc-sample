package client

import (
	"errors"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
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
	if dbc == nil {
		return errors.New("*dbClientImpl is nil")
	}
	return dbc.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewDBClient(tx))
	})
}

func (dbc *dbClientImpl) SelectOneUserByUserId(userId string) (*entity.User, domerr.DomErr) {
	if dbc == nil {
		return nil, domerr.NewDomErrFromMsg("*dbClientImpl is nil", domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
	}
	var user entity.User
	if err := dbc.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, domerr.NewDomErr(err, domerr.CAUSE_NOT_FOUND, domerr.LOG_LEVEL_INFO)
	}

	return &user, nil
}

func (dbc *dbClientImpl) CreateOneUser(u entity.User) (*entity.User, domerr.DomErr) {
	if dbc == nil {
		return nil, domerr.NewDomErrFromMsg("*dbClientImpl is nil", domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
	}
	if err := dbc.db.Create(&u).Error; err != nil {
		return nil, domerr.NewDomErr(err, domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
	}
	return &u, nil
}

func (dbc *dbClientImpl) CloseDB() error {
	if dbc == nil {
		return errors.New("*dbClientImpl is nil")
	}
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
