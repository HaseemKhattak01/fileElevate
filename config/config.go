package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoogleClientID     string
	GoogleClientSecret string
	RedirectURIs       string
	AppKey             string
	AppSecret          string
	RedirectURL        string
}

var config Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config = Config{
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURIs:       os.Getenv("REDIRECT_URIS"),
		AppKey:             os.Getenv("APP_KEY"),
		AppSecret:          os.Getenv("APP_SECRET"),
		RedirectURL:        os.Getenv("DROPBOX_REDIRECT_URL"),
	}
}

func GetConfig() Config {
	return config
}
