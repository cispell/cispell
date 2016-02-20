package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetHome(c *echo.Context) error {
	return c.Render(http.StatusOK, "index.html", (""))
}

func GetTerms(c *echo.Context) error {
	return c.String(http.StatusOK, "Terms page")
}

func GetPrivacy(c *echo.Context) error {
	return c.String(http.StatusOK, "Privacy page")
}

func GetAbout(c *echo.Context) error {
	return c.String(http.StatusOK, "About page")
}

func GetDocs(c *echo.Context) error {
	return c.String(http.StatusOK, "Docs page")
}
