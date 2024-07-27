package config

import (
	"fmt"
)

func GetDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", envVars.dbUser, envVars.dbPassword, envVars.dbHost, envVars.dbPort, envVars.dbDatabase)
}
