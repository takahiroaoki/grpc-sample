package config

import (
	"os"

	"github.com/takahiroaoki/go-env/util"
)

type EnvVars struct {
	/* About Database */
	dbHost     string
	dbPort     string
	dbDatabase string
	dbUser     string
	dbPassword string
}

func (e *EnvVars) GetDBHost() string {
	return e.dbHost
}

func (e *EnvVars) GetDBPort() string {
	return e.dbPort
}

func (e *EnvVars) GetDBDatabase() string {
	return e.dbDatabase
}

func (e *EnvVars) GetDBUser() string {
	return e.dbUser
}

func (e *EnvVars) GetDBPassword() string {
	return e.dbPassword
}

var envVars *EnvVars

func GetEnvVars() *EnvVars {
	if envVars == nil {
		util.InfoLog("Reading environment variables...")
		envVars = &EnvVars{
			dbHost:     os.Getenv("DB_HOST"),
			dbPort:     os.Getenv("DB_PORT"),
			dbDatabase: os.Getenv("DB_DATABASE"),
			dbUser:     os.Getenv("DB_USER"),
			dbPassword: os.Getenv("DB_PASSWORD"),
		}
	}
	return envVars
}
