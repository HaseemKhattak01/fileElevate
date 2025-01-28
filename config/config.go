package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoogleClientID     string
	GoogleClientSecret string
	AppKey             string
	AppSecret          string
	RedirectURI        string
}

var config Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config = Config{
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		AppKey:             os.Getenv("APP_KEY"),
		AppSecret:          os.Getenv("APP_SECRET"),
		RedirectURI:        os.Getenv("DROPBOX_REDIRECT_URI"),
	}
}

func GetConfig() Config {
	return config
}
