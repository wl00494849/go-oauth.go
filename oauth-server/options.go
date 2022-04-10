package oauthserver

import (
	"os"

	"github.com/rs/xid"
)

type ClientOption struct {
	clientID     string
	clientSecret string
	redirectUrl  string
}

func CreateClientOption(company, redirectUrl string) *ClientOption {
	var ID, Secret string
	switch company {
	case "google":
		ID = os.Getenv("Google_ID")
		Secret = os.Getenv("Google_Secret")
	default:
		ID = ""
		Secret = ""
	}

	return &ClientOption{
		clientID:     ID,
		clientSecret: Secret,
		redirectUrl:  redirectUrl,
	}
}

func (option *ClientOption) SetRedirectUrl(url string) {
	option.redirectUrl = url
}

func (option *ClientOption) GetRedirectUrl() string {
	return option.redirectUrl
}

func GenerateID() string {
	return xid.New().String()
}
