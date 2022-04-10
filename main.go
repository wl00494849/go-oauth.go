package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	oauthserver "go-oauth.go/oauth-server"
)

func main() {
	var port string
	flag.StringVar(&port, "Port", ":7788", "server port")

	godotenv.Load()
	app := gin.Default()

	oauth := app.Group("oauth")
	{
		oauth.GET("/google", oauthserver.GoogleOauthLogin)
	}

	callback := app.Group("callback")
	{
		callback.GET("/google", oauthserver.GoogleCallBack)
	}

	app.Run(port)
}
