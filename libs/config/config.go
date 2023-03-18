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
		Port				: os.Getenv(Port),
		Username		: os.Getenv(Username),
		Password		: os.Getenv(Password),
		DatabaseName: os.Getenv(DatabaseName),
		Host				: os.Getenv(Host),
	}
}