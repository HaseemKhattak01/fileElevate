package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/HaseemKhattak01/mydriveuploader/config"
	"golang.org/x/oauth2"
)

// getOAuthConfig returns the OAuth2 configuration for Dropbox
func getOAuthConfig(appKey, appSecret, redirectURI string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     appKey,
		ClientSecret: appSecret,
		RedirectURL:  redirectURI,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.dropbox.com/oauth2/authorize",
			TokenURL: "https://api.dropboxapi.com/oauth2/token",
		},
		Scopes: []string{"files.metadata.read", "files.content.write"},
	}
}

func GetOAuthStartURL() string {
	cfg := config.GetConfig()
	return fmt.Sprintf("https://www.dropbox.com/oauth2/authorize?client_id=%s&token_access_type=offline&response_type=code&redirect_uri=%s", cfg.AppKey, cfg.RedirectURI)
}

// StartOAuthFlow initiates the OAuth 2.0 flow by redirecting to the Dropbox authorization URL
func StartOAuthFlow(w http.ResponseWriter, r *http.Request) {
	authURL := GetOAuthStartURL()
	http.Redirect(w, r, authURL, http.StatusFound)
}

// HandleOAuthCallback processes the OAuth callback and exchanges the code for an access token
func HandleOAuthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	cfg := config.GetConfig()
	token, err := ExchangeCodeForToken(cfg.AppKey, cfg.AppSecret, code, cfg.RedirectURI)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get token: %v", err), http.StatusInternalServerError)
		return
	}

	tokenFile := "DropBoxToken.json"
	if err := SaveToken(tokenFile, token); err != nil {
		http.Error(w, fmt.Sprintf("Failed to save token: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Access Token saved successfully")
}

// ExchangeCodeForToken exchanges the authorization code for an access token
func ExchangeCodeForToken(appKey, appSecret, code, redirectURI string) (*oauth2.Token, error) {
	conf := getOAuthConfig(appKey, appSecret, redirectURI)
	token, err := conf.Exchange(context.TODO(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %w", err)
	}
	return token, nil
}
