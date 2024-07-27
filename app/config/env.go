package config

import (
	"os"

	"github.com/takahiroaoki/go-env/util"
)

var envVars *EnvVars

func init() {
	envVars = &EnvVars{
		dbHost:     os.Getenv("DB_HOST"),
		dbPort:     os.Getenv("DB_PORT"),
		dbDatabase: os.Getenv("DB_DATABASE"),
		dbUser:     os.Getenv("DB_USER"),
		dbPassword: os.Getenv("DB_PASSWORD"),
	}
	util.InfoLog("Environment variables loaded successfully")
}

type EnvVars struct {
	/* About Database */
	dbHost     string
	dbPort     string
	dbDatabase string
	dbUser     string
	dbPassword string
}
