package routes

import (
	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {
	initPage(e)
	initAuth(e)
	initUser(e)
}
