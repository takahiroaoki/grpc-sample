package config

import "fmt"

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

func NewDataBaseConfig() (*DataBaseConfig, error) {
	// TODO: error check because of failuer to get environmental variables
	return &DataBaseConfig{
		user:     "dev-user",
		password: "password",
		host:     "demo-mysql",
		port:     "3306",
		database: "demodb",
	}, nil
}
