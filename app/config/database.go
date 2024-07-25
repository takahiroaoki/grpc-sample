package config

import (
	"fmt"
)

type DataBaseConfig struct {
	user     string
	password string
	host     string
	port     string
	database string
}

func (dbc *DataBaseConfig) GetDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbc.user, dbc.password, dbc.host, dbc.port, dbc.database)
}

func NewDataBaseConfig() *DataBaseConfig {
	envVars := GetEnvVars()
	return &DataBaseConfig{
		user:     envVars.GetDBUser(),
		password: envVars.GetDBPassword(),
		host:     envVars.GetDBHost(),
		port:     envVars.GetDBPort(),
		database: envVars.GetDBDatabase(),
	}
}
