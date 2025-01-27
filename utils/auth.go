package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HaseemKhattak01/mydriveuploader/config"
	"github.com/HaseemKhattak01/mydriveuploader/models"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func GetDriveClient() (*http.Client, *models.ErrorResponse) {
	config := &oauth2.Config{
		ClientID:     config.GetGoogleClientID(),
		ClientSecret: config.GetGoogleClientSecret(),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
		Scopes:       []string{drive.DriveFileScope},
	}

	client, errResp := GetClient(config)
	if errResp != nil {
		return nil, errResp
	}
	return client, nil
}

func GetClient(config *oauth2.Config) (*http.Client, *models.ErrorResponse) {
	tokenFile := "token.json"
	token, err := tokenFromFile(tokenFile)
	if err != nil {
		token = getTokenFromWeb(config)
		if err := saveToken(tokenFile, token); err != nil {
			return nil, &models.ErrorResponse{Error: fmt.Sprintf("unable to save oauth token: %v", err)}
		}
	}

	tokenSource := config.TokenSource(context.Background(), token)
	client := oauth2.NewClient(context.Background(), tokenSource)

	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, &models.ErrorResponse{Error: fmt.Sprintf("unable to refresh token: %v", err)}
	}

	if newToken.AccessToken != token.AccessToken {
		if err := saveToken(tokenFile, newToken); err != nil {
			return nil, &models.ErrorResponse{Error: fmt.Sprintf("unable to save refreshed token: %v", err)}
		}
	}

	return client, nil
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return token
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token: %v", err)
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}
