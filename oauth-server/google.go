package oauthserver

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var google_config *oauth2.Config

type googleUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
	Hd            string `json:"hd"`
}

func getGoogleOauthUrl() *oauth2.Config {
	option := CreateClientOption("google", "")

	googleUrl := &oauth2.Config{
		ClientID:     option.clientID,
		ClientSecret: option.clientSecret,
		RedirectURL:  option.redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return googleUrl
}

func GoogleOauthLogin(ctx *gin.Context) {
	config := getGoogleOauthUrl()
	redirect_uri := config.AuthCodeURL("")

	ctx.Redirect(http.StatusSeeOther, redirect_uri)
}

func GoogleCallBack(ctx *gin.Context) {
	code := ctx.Query("code")
	token, _ := google_config.Exchange(ctx, code)

	client := google_config.Client(context.TODO(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		ctx.AbortWithError(500, err)
	}

	defer userInfo.Body.Close()

	info, _ := ioutil.ReadAll(userInfo.Body)
	var user googleUser
	json.Unmarshal(info, &user)

	ctx.JSON(200, user)
}
