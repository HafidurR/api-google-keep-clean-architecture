package config

import "os"

const (
	Port         = "DB_PORT"
	Username     = "DB_USERNAME"
	Password     = "DB_PASSWORD"
	DatabaseName = "DATABASE_NAME"
	Host         = "DB_HOST"
)

type Config struct {
	Port         string
	Username     string
	Password     string
	DatabaseName string
	Host         string
}

func Get() Config {
	return Config{
		Port				: os.Getenv("DBHOST"),
		Username		: os.Getenv("DB_USERNAME"),
		Password		: os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		Host				: os.Getenv("DB_HOST"),
	}
}