package config

import (
	"catetduit/internal/helper"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuth2Config struct {
	GoogleConfig      *oauth2.Config
	StateString       string
	RedirectUrl       string
	FailedRedirectUrl string
}

func NewOAuth2Config() *OAuth2Config {

	stateString, err := helper.GenerateRandomString(32)
	if err != nil {
		panic("failed to generate state string for OAuth2: " + err.Error())
	}

	return &OAuth2Config{
		GoogleConfig: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
		StateString:       stateString,
		RedirectUrl:       os.Getenv("OAUTH2_SUCCESS_REDIRECT_URL"),
		FailedRedirectUrl: os.Getenv("OAUTH2_FAILED_REDIRECT_URL"),
	}
}
