package handlers

import (
	"fmt"
	"net/http"

	"github.com/HaseemKhattak01/mydriveuploader/config"
	"github.com/HaseemKhattak01/mydriveuploader/oauth" // Assuming this is the correct package for oauth
)

// StartOAuthFlow initiates the OAuth 2.0 flow
func StartOAuthFlow(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()
	authURL, err := oauth.GenerateAuthURL(cfg.AppKey, cfg.AppSecret, cfg.RedirectURI)
	if err != nil {
		http.Error(w, "Failed to generate auth URL", http.StatusInternalServerError)
		return
	}
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

	token, err := oauth.TokenFromCode(cfg.AppKey, cfg.AppSecret, code, cfg.RedirectURI) // Assuming oauth package has TokenFromCode
	if err != nil {
		http.Error(w, "Failed to get token", http.StatusInternalServerError)
		return
	}

	// Save the access token securely
	fmt.Printf("Access Token: %s\n", token.Token)
	// Implement secure storage
}
