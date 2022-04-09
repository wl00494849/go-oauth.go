package main

import (
	"github.com/gin-gonic/gin"
	oauthserver "go-oauth.go/oauth-server"
)

func main() {
	app := gin.Default()

	oauth := app.Group("oauth")
	{
		oauth.GET("/google", oauthserver.GoogleOauthLogin)
	}

	callback := app.Group("callback")
	{
		callback.GET("/google", oauthserver.GoogleCallBack)
	}
}
