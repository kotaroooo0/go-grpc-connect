// config/config.go
package config

import (
	"os"
)

const DRIVER_NAME = "postgres"

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}

func (c DBConfig) ConnectionString() string {
	return "host=" + c.DBHost + " port=" + c.DBPort + " user=" + c.DBUser + " dbname=" + c.DBName + " password=" + c.DBPassword + " sslmode=disable"
}
