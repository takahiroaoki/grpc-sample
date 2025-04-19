package database

import (
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// check whether DBClient implements repository.DemoRepository interface
var _ repository.DemoRepository = (*DBClient)(nil)

type DBClient struct {
	db *gorm.DB
}

func (dbc *DBClient) Transaction(fn func(dr repository.DemoRepository) error) error {
	return dbc.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewDBClient(tx))
	})
}

func (dbc *DBClient) SelectOneUserByUserId(userId string) (*entity.User, domerr.DomErr) {
	var user user
	if err := dbc.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, domerr.NewDomErr(err, domerr.CAUSE_NOT_FOUND, domerr.LOG_LEVEL_INFO)
	}

	return convertUserSchema(user), nil
}

func (dbc *DBClient) CreateOneUser(u entity.User) (*entity.User, domerr.DomErr) {
	s := convertUserEntity(u)
	if err := dbc.db.Create(s).Error; err != nil {
		return nil, domerr.NewDomErr(err, domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
	}
	return convertUserSchema(*s), nil
}

func (dbc *DBClient) CloseDB() error {
	sqlDB, err := dbc.db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}
	return nil
}

func NewDBClientFromDSN(dataSourceName string) (*DBClient, error) {
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

func NewDBClient(db *gorm.DB) *DBClient {
	return &DBClient{
		db: db,
	}
}
