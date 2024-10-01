package testutil

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/takahiroaoki/grpc-sample/app/infra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetTestDBWrapper() (infra.DBWrapper, sqlmock.Sqlmock, error) {
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
	return infra.NewDBWrapper(db), sqlMock, nil
}
