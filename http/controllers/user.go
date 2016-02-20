package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetProfile(c *echo.Context) error {
	return c.String(http.StatusOK, "Profile page")
}

func GetRepos(c *echo.Context) error {
	return c.String(http.StatusOK, "Repositories page")
}
