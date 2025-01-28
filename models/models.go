package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type Token struct {
	AccessToken string
}