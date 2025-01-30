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
	"google.golang.org/api/drive/v2"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetDriveClient() (*http.Client, *models.ErrorResponse) {
	cfg := config.GetConfig()
	oauthConfig := &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  cfg.RedirectURIs,
		Scopes:       []string{drive.DriveFileScope},
	}

	return GetClient(oauthConfig)
}

func GetClient(oauthConfig *oauth2.Config) (*http.Client, *models.ErrorResponse) {
	tokenFile := "token.json"
	token, err := TokenFromFile(tokenFile)
	if err != nil {
		token = GetTokenFromWeb(oauthConfig)
		if err := SaveToken(tokenFile, token); err != nil {
			return nil, &models.ErrorResponse{Error: fmt.Sprintf("unable to save oauth token: %v", err)}
		}
	}

	tokenSource := oauthConfig.TokenSource(context.Background(), token)
	client := oauth2.NewClient(context.Background(), tokenSource)

	newToken, err := tokenSource.Token()
	if err != nil {
		fmt.Println("Token is invalid or expired. Please re-authenticate.")
		token = GetTokenFromWeb(oauthConfig)
		if err := SaveToken(tokenFile, token); err != nil {
			return nil, &models.ErrorResponse{Error: fmt.Sprintf("unable to save refreshed token: %v", err)}
		}
		client = oauth2.NewClient(context.Background(), oauthConfig.TokenSource(context.Background(), token))
	}

	if newToken.AccessToken != token.AccessToken {
		if err := SaveToken(tokenFile, newToken); err != nil {
			return nil, &models.ErrorResponse{Error: fmt.Sprintf("unable to save refreshed token: %v", err)}
		}
	}

	return client, nil
}

func GetTokenFromWeb(oauthConfig *oauth2.Config) *oauth2.Token {
	authURL := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser: \n%v\n", authURL)

	// Start a local server to handle the callback
	codeChan := make(chan string)
	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/oauth2callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Code not found", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Authorization successful, you can close this window.")
		codeChan <- code
	})

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	authCode := <-codeChan
	srv.Shutdown(context.Background())

	token, err := oauthConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return token
}

func TokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func SaveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token: %v", err)
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}
