package config

import (
	"fmt"
)

func GetDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", env.dbUser, env.dbPassword, env.dbHost, env.dbPort, env.dbDatabase)
}
