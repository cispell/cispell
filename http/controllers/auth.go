package controllers

import (
	// "github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/cispell/cispell/config"
	// "github.com/google/go-github/github"
	"net/http"
)

func GetLogin(c *echo.Context) error {
	oauthConf := config.Oauth()
	redirectTo := oauthConf.AuthCodeURL(config.OauthStateString())
	return c.Redirect(http.StatusTemporaryRedirect, redirectTo)
}

func GetLoginGithub(c *echo.Context) error {
	code := c.Query("code")
	state := c.Query("state")
	return c.String(http.StatusOK, code+state)
}

func GetLogout(c *echo.Context) error {
	return c.String(http.StatusOK, "Logout user from the website")
}
