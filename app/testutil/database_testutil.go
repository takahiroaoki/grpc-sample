package testutil

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/takahiroaoki/grpc-sample/app/infra/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMockDBClient() (*database.DBClient, sqlmock.Sqlmock, error) {
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
	return database.NewDBClient(db), sqlMock, nil
}
