package config

import (
	"fmt"
	"os"
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
	return &DataBaseConfig{
		user:     os.Getenv(ENV_DB_USER),
		password: os.Getenv(ENV_DB_PASSWORD),
		host:     os.Getenv(ENV_DB_HOST),
		port:     os.Getenv(ENV_DB_PORT),
		database: os.Getenv(ENV_DB_DATABASE),
	}
}
