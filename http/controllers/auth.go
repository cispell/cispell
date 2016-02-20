package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetLogin(c *echo.Context) error {
	return c.String(http.StatusOK, "Login page")
}

func GetLoginGithub(c *echo.Context) error {
	return c.String(http.StatusOK, "Callback for github login")
}

func GetLogout(c *echo.Context) error {
	return c.String(http.StatusOK, "Logout user from the website")
}
