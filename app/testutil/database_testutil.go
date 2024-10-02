package testutil

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/takahiroaoki/grpc-sample/app/backend"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMockDBClient() (backend.DBClient, sqlmock.Sqlmock, error) {
	sqlDB, sqlMock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      sqlDB,
			SkipInitializeWithVersion: true,
		},
	), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return backend.NewDBClient(db), sqlMock, nil
}
