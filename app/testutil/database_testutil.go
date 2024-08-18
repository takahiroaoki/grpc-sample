package testutil

import (
	"github.com/takahiroaoki/grpc-sample/app/config"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	return gorm.Open(
		mysql.Open(config.GetDataSourceName()),
		&gorm.Config{
			SkipDefaultTransaction: false,
		},
	)
}

func CleanDB(db *gorm.DB) error {
	err := db.Where("true").Delete(&entity.User{}).Error
	return err
}
