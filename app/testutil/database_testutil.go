package testutil

import (
	"github.com/takahiroaoki/go-env/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	return gorm.Open(
		mysql.Open(config.GetDataSourceName()),
		&gorm.Config{},
	)
}
