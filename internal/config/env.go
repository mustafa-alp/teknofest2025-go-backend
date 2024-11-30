package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LocalDB  DBConfig
	RemoteDB DBConfig
	DevName  string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		LocalDB: DBConfig{
			Host:     os.Getenv("DB_LOCAL_HOST"),
			Port:     os.Getenv("DB_LOCAL_PORT"),
			User:     os.Getenv("DB_LOCAL_USER"),
			Password: os.Getenv("DB_LOCAL_PASSWORD"),
			Name:     os.Getenv("DB_LOCAL_NAME"),
		},
		RemoteDB: DBConfig{
			Host:     os.Getenv("DB_REMOTE_HOST"),
			Port:     os.Getenv("DB_REMOTE_PORT"),
			User:     os.Getenv("DB_REMOTE_USER"),
			Password: os.Getenv("DB_REMOTE_PASSWORD"),
			Name:     os.Getenv("DB_REMOTE_NAME"),
		},
		DevName: os.Getenv("DEVELOPER_NAME"),
	}

	return config, nil
}