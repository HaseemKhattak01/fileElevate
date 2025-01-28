package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoogleClientID     string
	GoogleClientSecret string
	AccessToken        string
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
		AccessToken:        os.Getenv("DROPBOX_ACCESS_TOKEN"),
		AppKey:             os.Getenv("DROPBOX_APP_KEY"),
		AppSecret:          os.Getenv("DROPBOX_APP_SECRET"),
		RedirectURI:        os.Getenv("DROPBOX_REDIRECT_URL"),
	}
}

func GetConfig() Config {
	return config
}
