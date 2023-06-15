package helper

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerPort string
	SecretKey  string
	Database   Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		SecretKey:  os.Getenv("SECRET_KEY"),
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}

	return config
}
