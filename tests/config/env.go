package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var env *envVars

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(fmt.Errorf("failed to load .env file: %v", err))
	}
	env = &envVars{
		targetAddress: os.Getenv("TARGET_ADDRESS"),
	}
}

type envVars struct {
	targetAddress string
}
